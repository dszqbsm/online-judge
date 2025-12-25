package submissions

import (
	"context"

	"github.com/dszqbsm/online-judge/api/internal/svc"
	"github.com/dszqbsm/online-judge/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RetrieveSubmissionsLogic struct {
	logger logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRetrieveSubmissionsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RetrieveSubmissionsLogic {
	return &RetrieveSubmissionsLogic{
		logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RetrieveSubmissionsLogic) RetrieveSubmissions(req *types.RetrieveSubmissionsRequest) (resp *types.RetrieveSubmissionsResponse, err error) {
	// TODO: add your logic here and delete this line
	// TODO: new response first, set status code to 200

	return
}
