package model

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ CaseModel = (*customCaseModel)(nil)

type (
	// CaseModel is an interface to be customized, add more methods here,
	// and implement the added methods in customCaseModel.
	CaseModel interface {
		caseModel
		caseModelCustom
		withSession(session sqlx.Session) CaseModel
	}

	customCaseModel struct {
		*defaultCaseModel
	}

	caseModelCustom interface {
		FindAllByProblemKey(ctx context.Context, problem_key string) ([]Case, error)
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

func (m *customCaseModel) FindAllByProblemKey(ctx context.Context, problem_key string) ([]Case, error) {
	var resp []Case
	var args []interface{}
	var conditions string

	if problem_key != "" {
		conditions = "and `problem_key` = ?"
		args = append(args, problem_key)
	}

	query := fmt.Sprintf("select %s from %s where `delete_time` IS NULL", caseRows, m.table)
	if conditions != "" {
		query += conditions
	}
	query += "order by `id` desc"
	err := m.conn.QueryRowsCtx(ctx, &resp, query, args)
	switch err {
	case nil:
		return resp, nil
	case sqlx.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
