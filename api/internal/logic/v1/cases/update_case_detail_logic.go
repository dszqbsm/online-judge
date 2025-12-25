package cases

import (
	"context"

	"github.com/dszqbsm/online-judge/api/internal/svc"
	"github.com/dszqbsm/online-judge/api/internal/types"

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
	// TODO: add your logic here and delete this line
	// TODO: new response first, set status code to 200

	return
}
