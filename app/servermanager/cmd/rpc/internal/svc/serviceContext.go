package svc

import (
	"looklook/app/servermanager/cmd/rpc/internal/config"
	"looklook/app/servermanager/model"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config config.Config

	ServerModel model.ServerModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewMysql(c.DB.DataSource)

	return &ServiceContext{
		Config:      c,
		ServerModel: model.NewServerModel(sqlConn, c.Cache),
	}
}
