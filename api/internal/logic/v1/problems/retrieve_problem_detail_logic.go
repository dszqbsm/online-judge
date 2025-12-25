package problems

import (
	"context"
	"net/http"

	"github.com/dszqbsm/errorx"
	"github.com/dszqbsm/online-judge/api/internal/svc"
	"github.com/dszqbsm/online-judge/api/internal/types"
	"github.com/dszqbsm/online-judge/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type RetrieveProblemDetailLogic struct {
	logger logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRetrieveProblemDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RetrieveProblemDetailLogic {
	return &RetrieveProblemDetailLogic{
		logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RetrieveProblemDetailLogic) RetrieveProblemDetail(req *types.RetrieveProblemDetailRequest) (resp *types.RetrieveProblemDetailResponse, err error) {
	resp = &types.RetrieveProblemDetailResponse{}

	problem, err := l.svcCtx.PromblemModel.FindOneByKey(l.ctx, req.ProblemKey)
	if err == model.ErrNotFound {
		return nil, errorx.New(http.StatusBadRequest, "BAD_REQUEST", "problem not exists!")
	}
	if err != nil {
		return nil, err
	}

	resp.Data = types.ProblemDetail{
		Title:       problem.Title,
		Description: problem.Description,
		Inputdesc:   problem.InputDesc,
		Outputdesc:  problem.OutputDesc,
		SampleCases: problem.SampleCases,
		Difficulty:  problem.Difficulty,
		TimeLimit:   int(problem.TimeLimit),
		MemoryLimit: int(problem.MemoryLimit),
		Tags:        problem.Tags,
	}

	return resp, nil
}
