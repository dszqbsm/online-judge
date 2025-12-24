package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ ProblemModel = (*customProblemModel)(nil)

type (
	// ProblemModel is an interface to be customized, add more methods here,
	// and implement the added methods in customProblemModel.
	ProblemModel interface {
		problemModel
		withSession(session sqlx.Session) ProblemModel
	}

	customProblemModel struct {
		*defaultProblemModel
	}
)

// NewProblemModel returns a model for the database table.
func NewProblemModel(conn sqlx.SqlConn) ProblemModel {
	return &customProblemModel{
		defaultProblemModel: newProblemModel(conn),
	}
}

func (m *customProblemModel) withSession(session sqlx.Session) ProblemModel {
	return NewProblemModel(sqlx.NewSqlConnFromSession(session))
}
