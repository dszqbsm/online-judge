package submissions

import (
	"context"

	"github.com/dszqbsm/online-judge/api/internal/svc"
	"github.com/dszqbsm/online-judge/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RetrieveSubmissionDetailLogic struct {
	logger logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRetrieveSubmissionDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RetrieveSubmissionDetailLogic {
	return &RetrieveSubmissionDetailLogic{
		logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RetrieveSubmissionDetailLogic) RetrieveSubmissionDetail(req *types.RetrieveSubmissionDetailRequest) (resp *types.RetrieveSubmissionDetailResponse, err error) {
	// TODO: add your logic here and delete this line
	// TODO: new response first, set status code to 200

	return
}
