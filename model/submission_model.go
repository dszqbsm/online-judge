package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ SubmissionModel = (*customSubmissionModel)(nil)

type (
	// SubmissionModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSubmissionModel.
	SubmissionModel interface {
		submissionModel
		withSession(session sqlx.Session) SubmissionModel
	}

	customSubmissionModel struct {
		*defaultSubmissionModel
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
