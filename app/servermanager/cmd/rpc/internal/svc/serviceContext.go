package svc

import (
	"github.com/LINKHA/automatix/app/servermanager/cmd/rpc/internal/config"
	"github.com/LINKHA/automatix/app/servermanager/model"

	"github.com/bwmarrin/snowflake"
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config config.Config

	ServerModel       model.ServerModel
	Snowflake         *snowflake.Node
	Redis             *redis.Redis
	KqueueServerQueue *kq.Pusher
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewMysql(c.DB.DataSource)
	snowflake, _ := snowflake.NewNode(c.Id)
	rds := redis.MustNewRedis(c.Redis.RedisConf)

	return &ServiceContext{
		Config:            c,
		ServerModel:       model.NewServerModel(sqlConn, c.Cache),
		Snowflake:         snowflake,
		Redis:             rds,
		KqueueServerQueue: kq.NewPusher(c.KqConfServerQueue.Brokers, c.KqConfServerQueue.Topic),
	}
}
