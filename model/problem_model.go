package model

import (
	"context"
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ProblemModel = (*customProblemModel)(nil)

type (
	// ProblemModel is an interface to be customized, add more methods here,
	// and implement the added methods in customProblemModel.
	ProblemModel interface {
		problemModel
		problemModelCustom
		withSession(session sqlx.Session) ProblemModel
	}

	customProblemModel struct {
		*defaultProblemModel
	}

	problemModelCustom interface {
		FindAllByTitleDifficultyTags(ctx context.Context, title string, difficulty string, tags string) ([]Problem, error)
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

func (m *customProblemModel) FindAllByTitleDifficultyTags(ctx context.Context, title string, difficulty string, tags string) ([]Problem, error) {
	var resp []Problem
	var args []interface{}
	var conditions []string

	if title != "" {
		conditions = append(conditions, "and `title` like ?")
		args = append(args, "%"+title+"%")
	}
	if difficulty != "" {
		conditions = append(conditions, "and `difficulty` = ?")
		args = append(args, difficulty)
	}
	if tags != "" {
		conditions = append(conditions, "and `tags` = ?")
		args = append(args, tags)
	}

	query := fmt.Sprintf("select %s from %s where `delete_time` IS NULL", problemRows, m.table)
	if len(conditions) > 0 {
		query += strings.Join(conditions, " ")
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
