package cases

import (
	"context"

	"github.com/dszqbsm/online-judge/api/internal/svc"
	"github.com/dszqbsm/online-judge/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RetrieveCasesLogic struct {
	logger logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRetrieveCasesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RetrieveCasesLogic {
	return &RetrieveCasesLogic{
		logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RetrieveCasesLogic) RetrieveCases(req *types.RetrieveCasesRequest) (resp *types.RetrieveCasesResponse, err error) {
	resp = &types.RetrieveCasesResponse{}

	casesDetail, err := l.svcCtx.TestCaseModel.FindAllByProblemKey(l.ctx, req.ProblemKey)
	if err != nil {
		return nil, err
	}

	casesList := make([]types.CaseDetailList, 0, len(casesDetail))
	for _, caseDetail := range casesDetail {
		casesList = append(casesList, types.CaseDetailList{
			CaseKey: caseDetail.Key,
		})
	}
	resp.Data = casesList

	return resp, nil
}
