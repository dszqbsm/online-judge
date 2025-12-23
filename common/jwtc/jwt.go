package jwtc

import (
	"github.com/zeromicro/go-zero/core/jwt"
)

type JwtPayload struct {
	UserId int64 `json:"user_id"`
}

func GenerateToken(svcCtx *svc.ServiceContext, userId int64) (string, error) {
	claims := jwt.MapClaims{
		"userId": userId,
		"exp":    jwt.GetExpireAt(svcCtx.Config.JwtAuth.AccessExpire),
	}
	token, err := jwt.NewToken(svcCtx.Config.JwtAuth.AccessSecret, claims)
	if err != nil {
		return "", err
	}
	return token, nil
}

func ValidateToken(svcCtx *svc.ServiceContext, token string) (int64, error) {
	claims, err := jwt.ParseToken(svcCtx.Config.JwtAuth.AccessSecret, token)
}
