package submissions

import (
	"context"

	"github.com/dszqbsm/online-judge/api/internal/svc"
	"github.com/dszqbsm/online-judge/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateSubmissionDetailLogic struct {
	logger logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateSubmissionDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateSubmissionDetailLogic {
	return &CreateSubmissionDetailLogic{
		logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateSubmissionDetailLogic) CreateSubmissionDetail(req *types.CreateSubmissionDetailRequest) (resp *types.CreateSubmissionDetailResponse, err error) {
	// TODO: add your logic here and delete this line
	// TODO: new response first, set status code to 200

	return
}
