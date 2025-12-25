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
	// TODO: add your logic here and delete this line
	// TODO: new response first, set status code to 200

	return
}
