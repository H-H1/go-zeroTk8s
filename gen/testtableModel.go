package gen

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ TesttableModel = (*customTesttableModel)(nil)

type (
	// TesttableModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTesttableModel.
	TesttableModel interface {
		testtableModel
	}

	customTesttableModel struct {
		*defaultTesttableModel
	}
)

// NewTesttableModel returns a model for the database table.
func NewTesttableModel(conn sqlx.SqlConn, r *redis.Redis) TesttableModel {
	return &customTesttableModel{
		defaultTesttableModel: newTesttableModel(conn, r),
	}
}
