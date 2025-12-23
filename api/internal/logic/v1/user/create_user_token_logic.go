package user

import (
	"context"
	"net/http"

	"github.com/dszqbsm/errorx"
	"github.com/dszqbsm/online-judge/api/internal/svc"
	"github.com/dszqbsm/online-judge/api/internal/types"

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
	if _, err := l.svcCtx.UserModel.FindOneByUserName(l.ctx, req.UserName); err == ErrNotFound {
		return nil, errorx.New(http.StatusBadRequest, "BAD_REQUEST", "user not found!")
	}

	return
}
