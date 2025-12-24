package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ CaseModel = (*customCaseModel)(nil)

type (
	// CaseModel is an interface to be customized, add more methods here,
	// and implement the added methods in customCaseModel.
	CaseModel interface {
		caseModel
		withSession(session sqlx.Session) CaseModel
	}

	customCaseModel struct {
		*defaultCaseModel
	}
)

// NewCaseModel returns a model for the database table.
func NewCaseModel(conn sqlx.SqlConn) CaseModel {
	return &customCaseModel{
		defaultCaseModel: newCaseModel(conn),
	}
}

func (m *customCaseModel) withSession(session sqlx.Session) CaseModel {
	return NewCaseModel(sqlx.NewSqlConnFromSession(session))
}
