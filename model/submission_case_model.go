package model

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ SubmissionCaseModel = (*customSubmissionCaseModel)(nil)

type (
	// SubmissionCaseModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSubmissionCaseModel.
	SubmissionCaseModel interface {
		submissionCaseModel
		submissionCaseModelCustom
		withSession(session sqlx.Session) SubmissionCaseModel
	}

	customSubmissionCaseModel struct {
		*defaultSubmissionCaseModel
	}

	submissionCaseModelCustom interface {
		FindAllBySubmissionKey(ctx context.Context, submission_key string) ([]SubmissionCase, error)
	}
)

// NewSubmissionCaseModel returns a model for the database table.
func NewSubmissionCaseModel(conn sqlx.SqlConn) SubmissionCaseModel {
	return &customSubmissionCaseModel{
		defaultSubmissionCaseModel: newSubmissionCaseModel(conn),
	}
}

func (m *customSubmissionCaseModel) withSession(session sqlx.Session) SubmissionCaseModel {
	return NewSubmissionCaseModel(sqlx.NewSqlConnFromSession(session))
}

func (m *customSubmissionCaseModel) FindAllBySubmissionKey(ctx context.Context, submission_key string) ([]SubmissionCase, error) {
	var resp []SubmissionCase
	query := fmt.Sprintf("select %s from %s where `submission_key` = ?", submissionRows, m.table)
	err := m.conn.QueryRowsCtx(ctx, &resp, query, submission_key)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
