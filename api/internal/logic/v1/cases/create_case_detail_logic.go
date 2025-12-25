package cases

import (
	"context"

	"github.com/dszqbsm/online-judge/api/internal/svc"
	"github.com/dszqbsm/online-judge/api/internal/types"
	"github.com/dszqbsm/online-judge/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateCaseDetailLogic struct {
	logger logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateCaseDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateCaseDetailLogic {
	return &CreateCaseDetailLogic{
		logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateCaseDetailLogic) CreateCaseDetail(req *types.CreateCaseDetailRequest) (resp *types.GeneralResponse, err error) {
	resp = &types.GeneralResponse{}

	caseDetail := &model.TestCase{
		Key:            req.CaseKey,
		ProblemKey:     req.ProblemKey,
		Input:          req.Input,
		ExpectedOutput: req.Output,
		Score:          uint64(req.Score),
	}

	_, err = l.svcCtx.TestCaseModel.Insert(l.ctx, caseDetail)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
