package problems

import (
	"context"

	"github.com/dszqbsm/online-judge/api/internal/svc"
	"github.com/dszqbsm/online-judge/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RetrieveProblemsLogic struct {
	logger logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRetrieveProblemsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RetrieveProblemsLogic {
	return &RetrieveProblemsLogic{
		logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RetrieveProblemsLogic) RetrieveProblems(req *types.RetrieveProblemsRequest) (resp *types.RetrieveProblemsResponse, err error) {
	resp = &types.RetrieveProblemsResponse{}

	problems, err := l.svcCtx.PromblemModel.FindAllByTitleDifficultyTags(l.ctx, req.Title, req.Difficulty, req.Tags)
	if err != nil {
		return nil, err
	}

	problemsList := make([]types.ProblemListDetail, 0, len(problems))
	for _, problem := range problems {
		problemsList = append(problemsList, types.ProblemListDetail{
			ProblemKey: problem.Key,
			Title:      problem.Title,
			Difficulty: problem.Difficulty,
			Tags:       problem.Tags,
		})
	}
	resp.Data = problemsList

	return resp, nil
}
