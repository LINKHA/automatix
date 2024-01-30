package svc

import (
	"github.com/LINKHA/automatix/app/rolemanager/cmd/model"
	"github.com/LINKHA/automatix/app/rolemanager/cmd/rpc/internal/config"

	"github.com/bwmarrin/snowflake"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config config.Config

	RoleModel model.RoleModel
	Snowflake *snowflake.Node
	Redis     *redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewMysql(c.DB.DataSource)
	snowflake, _ := snowflake.NewNode(c.Id)
	rds := redis.MustNewRedis(c.Redis.RedisConf)

	return &ServiceContext{
		Config:    c,
		RoleModel: model.NewRoleModel(sqlConn, c.Cache),
		Snowflake: snowflake,
		Redis:     rds,
	}
}
