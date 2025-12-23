package user

import (
	"net/http"

	"github.com/dszqbsm/errorx"
	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/dszqbsm/online-judge/api/internal/logic/v1/user"
	"github.com/dszqbsm/online-judge/api/internal/svc"
	"github.com/dszqbsm/online-judge/api/internal/types"
)

func CreateUserHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CreateUserRequest
		if err := httpx.Parse(r, &req); err != nil {
			stat := errorx.New(http.StatusBadRequest, errorx.ErrorCodeInvalidParameter, err.Error())
			httpx.WriteJson(w, stat.StatusCode, stat)
			return
		}

		l := user.NewCreateUserLogic(r.Context(), svcCtx)
		resp, err := l.CreateUser(&req)

		switch errResp := err.(type) {
		case nil:
			if resp.StatusCode == 0 {
				resp.StatusCode = http.StatusOK
			}
			httpx.WriteJson(w, resp.StatusCode, resp)
		case *errorx.ErrorResponse:
			if errResp.StatusCode == 0 {
				errResp.StatusCode = http.StatusBadRequest
			}
			httpx.WriteJson(w, errResp.StatusCode, errResp)
		default:
			stat := errorx.New(http.StatusInternalServerError, errorx.ErrorCodeInvalidParameter, err.Error())
			httpx.WriteJson(w, stat.StatusCode, stat)
		}
	}
}
