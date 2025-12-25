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

type RetrieveCaseDetailLogic struct {
	logger logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRetrieveCaseDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RetrieveCaseDetailLogic {
	return &RetrieveCaseDetailLogic{
		logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RetrieveCaseDetailLogic) RetrieveCaseDetail(req *types.RetrieveCaseDetailRequest) (resp *types.RetrieveCaseDetailResponse, err error) {
	resp = &types.RetrieveCaseDetailResponse{}

	caseDetail, err := l.svcCtx.TestCaseModel.FindOneByKey(l.ctx, req.CaseKey)
	if err == model.ErrNotFound {
		return nil, errorx.New(http.StatusBadRequest, "BAD_REQUEST", "case not exists!")
	}
	if err != nil {
		return nil, err
	}

	resp.Data = types.CaseDetail{
		Input:  caseDetail.Input,
		Output: caseDetail.ExpectedOutput,
		Score:  int(caseDetail.Score),
	}

	return resp, nil
}
