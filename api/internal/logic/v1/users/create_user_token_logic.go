package users

import (
	"context"
	"net/http"

	"github.com/dszqbsm/errorx"
	"github.com/dszqbsm/online-judge/api/internal/svc"
	"github.com/dszqbsm/online-judge/api/internal/types"
	"github.com/dszqbsm/online-judge/common/jwtc"
	"github.com/dszqbsm/online-judge/model"
	"golang.org/x/crypto/bcrypt"

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
	resp = &types.CreateUserTokenResponse{}

	user, err := l.svcCtx.UserModel.FindOneByUserName(l.ctx, req.UserName)
	if err == model.ErrNotFound {
		return nil, errorx.New(http.StatusBadRequest, "BAD_REQUEST", "user not found!")
	}
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password))
	if err != nil {
		return nil, errorx.New(http.StatusBadRequest, "BAD_REQUEST", "password is incorrect!")
	}

	token, err := jwtc.GenerateToken(jwtc.JwtConfig(l.svcCtx.Config.Auth), int64(user.Id))
	if err != nil {
		return nil, err
	}

	resp.Data.Token = token

	return resp, nil
}
