package problems

import (
	"context"
	"net/http"

	"github.com/dszqbsm/errorx"
	"github.com/dszqbsm/online-judge/api/internal/svc"
	"github.com/dszqbsm/online-judge/api/internal/types"
	"github.com/dszqbsm/online-judge/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteProblemDetailLogic struct {
	logger logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteProblemDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteProblemDetailLogic {
	return &DeleteProblemDetailLogic{
		logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteProblemDetailLogic) DeleteProblemDetail(req *types.DeleteProblemDetailRequest) (resp *types.GeneralResponse, err error) {
	resp = &types.GeneralResponse{}

	problem, err := l.svcCtx.PromblemModel.FindOneByKey(l.ctx, req.ProblemKey)
	if err == model.ErrNotFound {
		return nil, errorx.New(http.StatusBadRequest, "BAD_REQUEST", "problem not exists!")
	}
	if err != nil {
		return nil, err
	}

	err = l.svcCtx.PromblemModel.Delete(l.ctx, problem.Id)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
