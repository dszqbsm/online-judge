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
	resp = &types.RetrieveSubmissionsResponse{}

	submissions, err := l.svcCtx.SubmissionModel.FindAllByUserIdProblemKey(l.ctx, int64(req.UserId), req.ProblemKey)
	if err != nil {
		return nil, err
	}

	submissionsFlag := make([]types.SubmissionsList, 0, len(submissions))
	for _, submission := range submissions {
		submissionsFlag = append(submissionsFlag, types.SubmissionsList{
			SubmissionKey: submission.Key,
			Status:        submission.Status,
			Score:         int(submission.Score),
			SubmitTime:    int(submission.TimeUsed),
		})
	}
	resp.Data = submissionsFlag

	return resp, nil
}
