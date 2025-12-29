package submissions

import (
	"bytes"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"syscall"
	"time"

	"github.com/dszqbsm/online-judge/api/internal/svc"
	"github.com/dszqbsm/online-judge/api/internal/types"
	"github.com/dszqbsm/online-judge/model"
	"github.com/google/uuid"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateSubmissionDetailLogic struct {
	logger logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateSubmissionDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateSubmissionDetailLogic {
	return &CreateSubmissionDetailLogic{
		logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateSubmissionDetailLogic) CreateSubmissionDetail(req *types.CreateSubmissionDetailRequest) (resp *types.CreateSubmissionDetailResponse, err error) {
	resp = &types.CreateSubmissionDetailResponse{}

	// 获取题目信息
	problem, err := l.svcCtx.PromblemModel.FindOneByKey(l.ctx, req.ProblemKey)
	if err != nil {
		return nil, err
	}

	// 获取测试用例信息
	testCase, err := l.svcCtx.TestCaseModel.FindAllByProblemKey(l.ctx, req.ProblemKey)
	if err != nil {
		return nil, err
	}
	totalCaseCount := len(testCase)

	// 创建提交记录
	submission := &model.Submission{
		Key:            strings.ReplaceAll(uuid.New().String(), "-", ""),
		UserId:         uint64(req.UserId),
		ProblemKey:     req.ProblemKey,
		Code:           req.Code,
		Language:       req.Language,
		Status:         "pending",
		ErrorMag:       sql.NullString{String: "", Valid: false},
		Score:          0,
		PassCaseCount:  0,
		TotalCaseCount: uint64(totalCaseCount),
		TimeUsed:       0,
		MemoryUsed:     0,
	}
	_, err = l.svcCtx.SubmissionModel.Insert(l.ctx, submission)
	if err != nil {
		return nil, err
	}

	// 创建隔离的临时运行环境
	tempDir := filepath.Join(os.TempDir(), "oj_submission_"+submission.Key)
	if err = os.MkdirAll(tempDir, 0700); err != nil {
		submission.Status = "runtime_error"
		submission.ErrorMag = sql.NullString{String: err.Error(), Valid: true}
		if updateErr := l.svcCtx.SubmissionModel.Update(l.ctx, submission); updateErr != nil {
			return nil, updateErr
		}
		return nil, err
	}
	defer func() {
		os.RemoveAll(tempDir)
	}()

	// 编译用户代码
	var runCmd *exec.Cmd
	if compileErr := l.compileUserCode(req, tempDir, &runCmd); compileErr != nil {
		submission.Status = "compile_error"
		submission.ErrorMag = sql.NullString{String: compileErr.Error(), Valid: true}
		if updateErr := l.svcCtx.SubmissionModel.Update(l.ctx, submission); updateErr != nil {
			return nil, updateErr
		}
		return nil, err
	}

	// 逐用例判题，创建用例判题记录，通过标志位记录判题结果状态并记录最大用例用时和内存占用
	passCount, maxTimeUsed, maxMenUsed := 0, int64(0), int64(0)
	hasRE, hasTLE, hasMLE, hasWA := false, false, false, false
	for _, tc := range testCase {
		caseStatus, caseTime, caseMen, caseActualOutput, caseErrMsg := l.runSingleTestCase(l.ctx, runCmd, tempDir, int64(problem.TimeLimit), int64(problem.MemoryLimit), &tc)

		if caseTime > maxTimeUsed {
			maxTimeUsed = caseTime
		}
		if caseMen > maxMenUsed {
			maxMenUsed = caseMen
		}

		switch caseStatus {
		case "re":
			hasRE = true
		case "tle":
			hasTLE = true
		case "mle":
			hasMLE = true
		case "wa":
			hasWA = true
		case "ac":
			passCount++
		}

		caseKey := strings.ReplaceAll(uuid.New().String(), "-", "")[:32]
		isPassed := 0
		if caseStatus == "ac" {
			isPassed = 1
		}
		submissionCase := &model.SubmissionCase{
			Key:           caseKey,
			SubmissionKey: submission.Key,
			TestCaseKey:   tc.Key,
			UserId:        submission.UserId,
			ProblemKey:    submission.ProblemKey,
			Passed:        int64(isPassed),
			ActualOutput:  sql.NullString{String: caseActualOutput, Valid: true},
			ErrorMsg:      sql.NullString{String: caseErrMsg, Valid: true},
			TimeUsed:      sql.NullInt64{Int64: caseTime, Valid: true},
			MemoryUsed:    sql.NullInt64{Int64: caseMen, Valid: true},
		}
		_, err = l.svcCtx.SubmissionCaseModel.Insert(l.ctx, submissionCase)
		if err != nil {
			return nil, err
		}
	}

	// 根据标志位记录判题状态和报错信息，更新提交记录
	var finalStatus, finalMsg string
	switch {
	case hasRE:
		finalStatus = "runtime_error"
		finalMsg = "运行时错误"
	case hasTLE:
		finalStatus = "time_limit_exceeded"
		finalMsg = "运行超时"
	case hasMLE:
		finalStatus = "memory_limit_exceeded"
		finalMsg = "内存超限"
	case hasWA:
		finalStatus = "wrong_answer"
		finalMsg = "答案错误"
	default:
		finalStatus = "accepted"
		finalMsg = "Accepted"
	}
	score := float64(passCount) / float64(len(testCase)) * 100
	submission.Score = uint64(score)
	submission.PassCaseCount = uint64(passCount)
	submission.Status = finalStatus
	submission.ErrorMag = sql.NullString{String: finalMsg, Valid: true}
	submission.TimeUsed = int64(maxTimeUsed)
	submission.MemoryUsed = int64(maxMenUsed)
	_, err = l.svcCtx.SubmissionModel.Insert(l.ctx, submission)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (l *CreateSubmissionDetailLogic) compileUserCode(req *types.CreateSubmissionDetailRequest, tempDir string, runCmd **exec.Cmd) error {
	var codeFilePath string
	var compileCmd *exec.Cmd

	switch strings.ToLower(req.Language) {
	case "go":
		codeFilePath = filepath.Join(tempDir, "main.go")
		if err := os.WriteFile(codeFilePath, []byte(req.Code), 0600); err != nil {
			return errors.New(err.Error() + "写入Go代码文件失败")
		}
		compileCmd = exec.Command("go", "build", "-o", filepath.Join(tempDir, "main"), codeFilePath)
	default:
		return errors.New("不支持的编程语言：" + req.Language)
	}

	var compileErrBuf bytes.Buffer
	compileCmd.Stderr = &compileErrBuf
	compileCmd.Dir = tempDir
	if err := compileCmd.Run(); err != nil {
		return errors.New(err.Error() + "编译错误：" + compileErrBuf.String())
	}

	*runCmd = exec.Command(filepath.Join(tempDir, "main"))
	return nil
}

// 执行单个测试用例
func (l *CreateSubmissionDetailLogic) runSingleTestCase(ctx context.Context, baseCmd *exec.Cmd, tempDir string, timeLimitMs int64, memoryLimitKB int64, testCase *model.TestCase) (status string, timeUsed int64, memoryUsed int64, actualOutput string, errMsg string) {
	status = "ac"
	actualOutput = ""
	errMsg = ""
	timeUsed = 0
	memoryUsed = 0

	// 配置子进程可执行文件路径，配置子进程工作目录
	runCmd := &exec.Cmd{
		Path: baseCmd.Path,
		Dir:  tempDir,
	}

	// 配置输入输出，设置进程组方便超时单独杀死
	var outputBuf, errBuf bytes.Buffer
	runCmd.Stdin = strings.NewReader(testCase.Input)
	runCmd.Stdout = &outputBuf
	runCmd.Stderr = &errBuf
	runCmd.SysProcAttr = &syscall.SysProcAttr{
		Setpgid: true,
	}

	// 运行时间限制
	var rLimit syscall.Rlimit
	rLimit.Cur = uint64(timeLimitMs / 1000)
	rLimit.Max = uint64(timeLimitMs/1000 + 1)
	_ = syscall.Setrlimit(syscall.RLIMIT_CPU, &rLimit)

	// 内存占用限制
	rLimit.Cur = uint64(memoryLimitKB * 1024)
	rLimit.Max = uint64(memoryLimitKB * 1024)
	_ = syscall.Setrlimit(syscall.RLIMIT_AS, &rLimit)

	ctxWithTimeout, cancel := context.WithTimeout(l.ctx, time.Duration(timeLimitMs)*time.Millisecond)
	defer cancel()

	// 启动进程
	startTime := time.Now()
	if err := runCmd.Start(); err != nil {
		status = "re"
		errMsg = "程序启动失败: " + err.Error()
		return status, 0, 0, "", errMsg
	}

	// 异步等待进程结束
	done := make(chan error, 1)
	go func() { done <- runCmd.Wait() }()

	// 等待进程运行结果
	select {
	case <-ctxWithTimeout.Done():
		// 超时
		syscall.Kill(-runCmd.Process.Pid, syscall.SIGKILL)
		status = "tle"
		errMsg = fmt.Sprintf("运行超时(限制: %dms)", timeLimitMs)
		// 用+1标记为超时，需要调整
		timeUsed = timeLimitMs + 1
		return status, timeUsed, 0, "", errMsg

	case err := <-done:
		// 计算运行时间和内存占用
		timeUsed = int64(time.Since(startTime).Milliseconds())
		if runCmd.Process != nil {
			memoryUsed, _ = l.GetProcessMemoryUsage(runCmd.Process.Pid)
		}
		if memoryUsed > memoryLimitKB {
			status = "mle"
			errMsg = fmt.Sprintf("内存超限(使用: %dKB, 限制: %dKB)", memoryUsed, memoryLimitKB)
			actualOutput = strings.TrimSpace(outputBuf.String())
			return status, timeUsed, memoryUsed, actualOutput, errMsg
		}

		// 根据运行结果设置输出
		if err != nil {
			exitErr, ok := err.(*exec.ExitError)
			if !ok {
				status = "re"
				errMsg = "运行错误: " + err.Error()
			} else {
				status = "re"
				errMsg = fmt.Sprintf("程序异常退出(退出码: %d), 错误信息: %s", exitErr.ExitCode(), errBuf.String())
			}
			actualOutput = strings.TrimSpace(outputBuf.String())
			return status, timeUsed, memoryUsed, actualOutput, errMsg
		}

		actualOutput = strings.TrimSpace(outputBuf.String())
		expectedOutput := strings.TrimSpace(testCase.ExpectedOutput)
		if actualOutput != expectedOutput {
			status = "wa"
			errMsg = fmt.Sprintf("答案错误(期望: %s, 实际: %s)", expectedOutput, actualOutput)
			return status, timeUsed, memoryUsed, actualOutput, errMsg
		}

		actualOutput = outputBuf.String()
		return status, timeUsed, memoryUsed, actualOutput, errMsg
	}
}

// 获取进程内存使用量
func (l *CreateSubmissionDetailLogic) GetProcessMemoryUsage(pid int) (int64, error) {
	statmPath := filepath.Join("/proc", strconv.Itoa(pid), "statm")
	content, err := os.ReadFile(statmPath)
	if err != nil {
		return 0, err
	}

	fields := strings.Fields(string(content))
	if len(fields) < 2 {
		return 0, os.ErrInvalid
	}

	residentPages, _ := strconv.ParseInt(fields[1], 10, 64)
	pageSize := int64(syscall.Getpagesize() / 1024)
	return residentPages * pageSize, nil
}
