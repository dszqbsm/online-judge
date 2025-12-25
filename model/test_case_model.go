package model

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ TestCaseModel = (*customTestCaseModel)(nil)

type (
	// TestCaseModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTestCaseModel.
	TestCaseModel interface {
		testCaseModel
		testCaseModelCustom
		withSession(session sqlx.Session) TestCaseModel
	}

	customTestCaseModel struct {
		*defaultTestCaseModel
	}

	testCaseModelCustom interface {
		FindAllByProblemKey(ctx context.Context, problem_key string) ([]TestCase, error)
	}
)

// NewTestCaseModel returns a model for the database table.
func NewTestCaseModel(conn sqlx.SqlConn) TestCaseModel {
	return &customTestCaseModel{
		defaultTestCaseModel: newTestCaseModel(conn),
	}
}

func (m *customTestCaseModel) withSession(session sqlx.Session) TestCaseModel {
	return NewTestCaseModel(sqlx.NewSqlConnFromSession(session))
}

func (m *customTestCaseModel) FindAllByProblemKey(ctx context.Context, problem_key string) ([]TestCase, error) {
	var resp []TestCase
	var args []interface{}
	var conditions string

	if problem_key != "" {
		conditions = "and `problem_key` = ?"
		args = append(args, conditions)
	}

	query := fmt.Sprintf("select %s from %s where `delete_time` IS NULL", testCaseRows, m.table)
	if len(conditions) > 0 {
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
