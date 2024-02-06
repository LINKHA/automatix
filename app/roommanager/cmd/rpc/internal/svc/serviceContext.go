package svc

import (
	"github.com/LINKHA/automatix/app/roommanager/cmd/rpc/internal/config"
	"github.com/LINKHA/automatix/common/utils"
	"github.com/bwmarrin/snowflake"
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

type ServiceContext struct {
	Config        config.Config
	StreamManager *utils.ShardLockMaps

	Snowflake         *snowflake.Node
	Redis             *redis.Redis
	KqueueServerQueue *kq.Pusher
}

func NewServiceContext(c config.Config) *ServiceContext {
	streamManager := utils.NewShardLockMaps()

	snowflake, _ := snowflake.NewNode(c.Id)
	rds := redis.MustNewRedis(c.Redis.RedisConf)

	return &ServiceContext{
		Config:            c,
		StreamManager:     &streamManager,
		Snowflake:         snowflake,
		Redis:             rds,
		KqueueServerQueue: kq.NewPusher(c.KqConfServerQueue.Brokers, c.KqConfServerQueue.Topic),
	}
}
