package users

import (
	"context"
	"database/sql"
	"net/http"

	"github.com/dszqbsm/errorx"
	"github.com/dszqbsm/online-judge/api/internal/svc"
	"github.com/dszqbsm/online-judge/api/internal/types"
	"github.com/dszqbsm/online-judge/model"
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

	if _, err := l.svcCtx.UserModel.FindOneByUserName(l.ctx, req.UserName); err == nil {
		return nil, errorx.New(http.StatusBadRequest, "BAD_REQUEST", "user name already exists!")
	}
	if err != nil && err != model.ErrNotFound {
		return nil, err
	}

	hashBytes, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := model.User{
		UserName:     req.UserName,
		PasswordHash: string(hashBytes),
		Status:       true,
		DeleteTime:   sql.NullTime{},
	}

	_, err = l.svcCtx.UserModel.Insert(l.ctx, &user)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
