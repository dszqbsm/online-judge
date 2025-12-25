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

type UpdateProblemDetailLogic struct {
	logger logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateProblemDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateProblemDetailLogic {
	return &UpdateProblemDetailLogic{
		logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateProblemDetailLogic) UpdateProblemDetail(req *types.UpdateProblemDetailRequest) (resp *types.GeneralResponse, err error) {
	resp = &types.GeneralResponse{}

	problem, err := l.svcCtx.PromblemModel.FindOneByKey(l.ctx, req.ProblemKey)
	if err == model.ErrNotFound {
		return nil, errorx.New(http.StatusBadRequest, "BAD_REQUEST", "problem not exists!")
	}
	if err != nil {
		return nil, err
	}

	if req.Title != "" {
		problem.Title = req.Title
	}
	if req.Description != "" {
		problem.Description = req.Description
	}
	if req.Inputdesc != "" {
		problem.InputDesc = req.Inputdesc
	}
	if req.Outputdesc != "" {
		problem.OutputDesc = req.Outputdesc
	}
	if req.SampleCases != "" {
		problem.SampleCases = req.SampleCases
	}
	if req.Difficulty != "" {
		problem.Difficulty = req.Difficulty
	}
	if req.TimeLimit != nil {
		problem.TimeLimit = int64(*req.TimeLimit)
	}
	if req.MemoryLimit != nil {
		problem.MemoryLimit = int64(*req.MemoryLimit)
	}
	if req.Tags != "" {
		problem.Tags = req.Tags
	}
	if req.IsPublished != nil {
		problem.IsPublished = bool(*req.IsPublished)
	}

	err = l.svcCtx.PromblemModel.Update(l.ctx, problem)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
