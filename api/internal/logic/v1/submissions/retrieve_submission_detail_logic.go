package submissions

import (
	"context"
	"database/sql"

	"github.com/dszqbsm/online-judge/api/internal/svc"
	"github.com/dszqbsm/online-judge/api/internal/types"
	"github.com/dszqbsm/online-judge/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type RetrieveSubmissionDetailLogic struct {
	logger logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRetrieveSubmissionDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RetrieveSubmissionDetailLogic {
	return &RetrieveSubmissionDetailLogic{
		logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RetrieveSubmissionDetailLogic) RetrieveSubmissionDetail(req *types.RetrieveSubmissionDetailRequest) (resp *types.RetrieveSubmissionDetailResponse, err error) {
	resp = &types.RetrieveSubmissionDetailResponse{}

	// 获取提交记录
	submission, err := l.svcCtx.SubmissionModel.FindOneByKey(l.ctx, req.SubmissionKey)
	if err != nil {
		return nil, err
	}

	// 获取该问题的所有测试用例信息
	testCases, err := l.svcCtx.TestCaseModel.FindAllByProblemKey(l.ctx, submission.ProblemKey)
	if err != nil {
		return nil, err
	}
	testCasesFlag := make(map[string]model.TestCase, len(testCases))
	for _, testCase := range testCases {
		testCasesFlag[testCase.Key] = testCase
	}

	// 获取该问题的所有测试用例执行结果
	submissionCases, err := l.svcCtx.SubmissionCaseModel.FindAllBySubmissionKey(l.ctx, req.SubmissionKey)
	if err != nil {
		return nil, err
	}

	submissionCasesFlag := make([]types.CaseResult, 0, len(submissionCases))
	for _, submissionCase := range submissionCases {
		caseResult := types.CaseResult{
			CaseKey:      submissionCase.Key,
			Passed:       int(submissionCase.Passed),
			Input:        testCasesFlag[submissionCase.TestCaseKey].Input,
			ActualOutput: sql.NullString{String: submissionCase.ActualOutput.String(), Valid: true},
		}
	}

	resp.Data = submission

	return
}
