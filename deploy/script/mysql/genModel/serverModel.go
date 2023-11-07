package genModel

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ServerModel = (*customServerModel)(nil)

type (
	// ServerModel is an interface to be customized, add more methods here,
	// and implement the added methods in customServerModel.
	ServerModel interface {
		serverModel
	}

	customServerModel struct {
		*defaultServerModel
	}
)

// NewServerModel returns a model for the database table.
func NewServerModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) ServerModel {
	return &customServerModel{
		defaultServerModel: newServerModel(conn, c, opts...),
	}
}
