package user

import (
	"context"
	"net/http"

	"github.com/dszqbsm/errorx"
	"github.com/dszqbsm/online-judge/api/internal/svc"
	"github.com/dszqbsm/online-judge/api/internal/types"
	"golang.org/x/crypto/bcrypt"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateUserLogic struct {
	logger logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateUserLogic {
	return &CreateUserLogic{
		logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateUserLogic) CreateUser(req *types.CreateUserRequest) (resp *types.GeneralResponse, err error) {
	resp = &types.GeneralResponse{}

	user, err := l.svcCtx.UserModel.FindOneByUserName(l.ctx, req.UserName)
	if err != nil {
		return nil, err
	}
	if user.Id != 0 {
		return nil, errorx.New(http.StatusBadRequest, "BAD_REQUEST", "not in family")
	}

	hashBytes, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	userId, err := l.svcCtx.UserModel.


	return
}
