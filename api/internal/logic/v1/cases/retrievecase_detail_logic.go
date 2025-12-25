package cases

import (
	"context"

	"github.com/dszqbsm/online-judge/api/internal/svc"
	"github.com/dszqbsm/online-judge/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RetrievecaseDetailLogic struct {
	logger logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRetrievecaseDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RetrievecaseDetailLogic {
	return &RetrievecaseDetailLogic{
		logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RetrievecaseDetailLogic) RetrievecaseDetail(req *types.RetrievecaseDetailRequest) (resp *types.RetrievecaseDetailResponse, err error) {
	// TODO: add your logic here and delete this line
	// TODO: new response first, set status code to 200

	return
}
