package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ SubmissionCaseModel = (*customSubmissionCaseModel)(nil)

type (
	// SubmissionCaseModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSubmissionCaseModel.
	SubmissionCaseModel interface {
		submissionCaseModel
		withSession(session sqlx.Session) SubmissionCaseModel
	}

	customSubmissionCaseModel struct {
		*defaultSubmissionCaseModel
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
