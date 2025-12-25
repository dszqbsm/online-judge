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

type DeleteCaseDetailLogic struct {
	logger logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteCaseDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteCaseDetailLogic {
	return &DeleteCaseDetailLogic{
		logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteCaseDetailLogic) DeleteCaseDetail(req *types.DeleteCaseDetailRequest) (resp *types.GeneralResponse, err error) {
	resp = &types.GeneralResponse{}

	caseDetail, err := l.svcCtx.TestCaseModel.FindOneByKey(l.ctx, req.CaseKey)
	if err == model.ErrNotFound {
		return nil, errorx.New(http.StatusBadRequest, "BAD_REQUEST", "problem not exists!")
	}
	if err != nil {
		return nil, err
	}

	err = l.svcCtx.TestCaseModel.Delete(l.ctx, caseDetail.Id)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
