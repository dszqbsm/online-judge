package cases

import (
	"context"
	"net/http"

	"github.com/dszqbsm/errorx"
	"github.com/dszqbsm/online-judge/api/internal/svc"
	"github.com/dszqbsm/online-judge/api/internal/types"
	"github.com/dszqbsm/online-judge/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateCaseDetailLogic struct {
	logger logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateCaseDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateCaseDetailLogic {
	return &UpdateCaseDetailLogic{
		logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateCaseDetailLogic) UpdateCaseDetail(req *types.UpdateCaseDetailRequest) (resp *types.GeneralResponse, err error) {
	resp = &types.GeneralResponse{}

	caseDetail, err := l.svcCtx.CaseModel.FindOneByKey(l.ctx, req.CaseKey)
	if err == model.ErrNotFound {
		return nil, errorx.New(http.StatusBadRequest, "BAD_REQUEST", "case not exists!")
	}
	if err != nil {
		return nil, err
	}

	if req.ProblemKey != "" {
		caseDetail.ProblemKey = req.ProblemKey
	}

	if req.Input != "" {
		caseDetail.Input = req.Input
	}

	if req.Output != "" {
		caseDetail.Output = req.Output
	}

	if req.Score != nil {
		caseDetail.Score = int64(*req.Score)
	}

	err = l.svcCtx.CaseModel.Update(l.ctx, caseDetail)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
