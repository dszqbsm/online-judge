package cases

import (
	"context"

	"github.com/dszqbsm/online-judge/api/internal/svc"
	"github.com/dszqbsm/online-judge/api/internal/types"

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
	// TODO: add your logic here and delete this line
	// TODO: new response first, set status code to 200

	return
}
