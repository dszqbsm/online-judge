package user

import (
	"context"

	"gitlab.homerunsmartapi.com/go-zero-libs/templatex/api/internal/svc"
	"gitlab.homerunsmartapi.com/go-zero-libs/templatex/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateUserTokenLogic struct {
	logger logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateUserTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateUserTokenLogic {
	return &CreateUserTokenLogic{
		logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateUserTokenLogic) CreateUserToken(req *types.CreateUserTokenRequest) (resp *types.CreateUserTokenResponse, err error) {
	// TODO: add your logic here and delete this line
	// TODO: new response first, set status code to 200

	return
}
