// Code generated by goctl. DO NOT EDIT.

package gen

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	testtableFieldNames          = builder.RawFieldNames(&Testtable{})
	testtableRows                = strings.Join(testtableFieldNames, ",")
	testtableRowsExpectAutoSet   = strings.Join(stringx.Remove(testtableFieldNames, "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	testtableRowsWithPlaceHolder = strings.Join(stringx.Remove(testtableFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	testtableModel interface {
		Insert(ctx context.Context, data *Testtable) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*Testtable, error)
		Update(ctx context.Context, data *Testtable) error
		Delete(ctx context.Context, id int64) error
	}

	defaultTesttableModel struct {
		conn  sqlx.SqlConn
		table string
	}

	Testtable struct {
		Id   int64          `db:"id"`
		Name sql.NullString `db:"name"`
	}
)

func newTesttableModel(conn sqlx.SqlConn,r *redis.Redis) *defaultTesttableModel {
	return &defaultTesttableModel{
		conn:  conn,
		table: "`testtable`",
	}
}

func (m *defaultTesttableModel) withSession(session sqlx.Session) *defaultTesttableModel {
	return &defaultTesttableModel{
		conn:  sqlx.NewSqlConnFromSession(session),
		table: "`testtable`",
	}
}

func (m *defaultTesttableModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultTesttableModel) FindOne(ctx context.Context, id int64) (*Testtable, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", testtableRows, m.table)
	var resp Testtable
	err := m.conn.QueryRowCtx(ctx, &resp, query, id)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultTesttableModel) Insert(ctx context.Context, data *Testtable) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?)", m.table, testtableRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.Id, data.Name)
	return ret, err
}

func (m *defaultTesttableModel) Update(ctx context.Context, data *Testtable) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, testtableRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.Name, data.Id)
	return err
}

func (m *defaultTesttableModel) tableName() string {
	return m.table
}
