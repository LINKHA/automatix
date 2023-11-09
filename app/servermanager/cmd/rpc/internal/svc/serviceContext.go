package svc

import (
	"automatix/app/servermanager/cmd/rpc/internal/config"
	"automatix/app/servermanager/model"

	"github.com/bwmarrin/snowflake"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config config.Config

	ServerModel model.ServerModel
	Snowflake   *snowflake.Node
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewMysql(c.DB.DataSource)
	snowflake, _ := snowflake.NewNode(c.Id)
	return &ServiceContext{
		Config:      c,
		ServerModel: model.NewServerModel(sqlConn, c.Cache),
		Snowflake:   snowflake,
	}
}
