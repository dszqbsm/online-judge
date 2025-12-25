package cases

import (
	"context"

	"github.com/dszqbsm/online-judge/api/internal/svc"
	"github.com/dszqbsm/online-judge/api/internal/types"

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
	// TODO: add your logic here and delete this line
	// TODO: new response first, set status code to 200

	return
}
