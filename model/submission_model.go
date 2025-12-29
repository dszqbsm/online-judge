package model

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ SubmissionModel = (*customSubmissionModel)(nil)

type (
	// SubmissionModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSubmissionModel.
	SubmissionModel interface {
		submissionModel
		submissionModelCustom
		withSession(session sqlx.Session) SubmissionModel
	}

	customSubmissionModel struct {
		*defaultSubmissionModel
	}

	submissionModelCustom interface {
		FindAllByUserIdProblemKey(ctx context.Context, user_id int64, problem_key string) ([]Submission, error)
	}
)

// NewSubmissionModel returns a model for the database table.
func NewSubmissionModel(conn sqlx.SqlConn) SubmissionModel {
	return &customSubmissionModel{
		defaultSubmissionModel: newSubmissionModel(conn),
	}
}

func (m *customSubmissionModel) withSession(session sqlx.Session) SubmissionModel {
	return NewSubmissionModel(sqlx.NewSqlConnFromSession(session))
}

func (m *customSubmissionModel) FindAllByUserIdProblemKey(ctx context.Context, user_id int64, problem_key string) ([]Submission, error) {
	var resp []Submission
	query := fmt.Sprintf("select %s from %s where `user_id` = ? and `problem_key` = ?", submissionRows, m.table)
	err := m.conn.QueryRowsCtx(ctx, &resp, query, user_id, problem_key)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
