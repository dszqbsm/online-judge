package problems

import (
	"context"

	"github.com/dszqbsm/online-judge/api/internal/svc"
	"github.com/dszqbsm/online-judge/api/internal/types"
	"github.com/dszqbsm/online-judge/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateProblemDetailLogic struct {
	logger logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateProblemDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateProblemDetailLogic {
	return &CreateProblemDetailLogic{
		logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateProblemDetailLogic) CreateProblemDetail(req *types.CreateProblemDetailRequest) (resp *types.GeneralResponse, err error) {
	resp = &types.GeneralResponse{}

	problem := &model.Problem{
		Key:         req.ProblemKey,
		Title:       req.Title,
		Description: req.Description,
		InputDesc:   req.Inputdesc,
		OutputDesc:  req.Outputdesc,
		SampleCases: req.SampleCases,
		Difficulty:  req.Difficulty,
		TimeLimit:   int64(req.TimeLimit),
		MemoryLimit: int64(req.MemoryLimit),
		Tags:        req.Tags,
		IsPublished: req.IsPublished,
	}

	_, err = l.svcCtx.PromblemModel.Insert(l.ctx, problem)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
