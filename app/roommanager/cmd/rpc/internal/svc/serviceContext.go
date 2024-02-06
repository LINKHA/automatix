package svc

import (
	"github.com/LINKHA/automatix/app/roommanager/cmd/rpc/internal/config"
	"github.com/LINKHA/automatix/common/utils"
	"github.com/bwmarrin/snowflake"
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

type Group struct {
	GroupID   string   `json:"group_id"`
	GroupName string   `json:"group_name"`
	MaxPlayer int32    `json:"max_player"`
	RoomID    string   `json:"room_id"`
	Roles     []string `json:"roles"`
}

type Room struct {
	RoomId    string   `json:"room_id"`
	RoomName  string   `json:"room_name"`
	MaxPlayer int32    `json:"max_player"`
	Roles     []string `json:"roles"`
	Groups    []string `json:"groups"`
}

type Role struct {
	RoleId  string `json:"role_id"`
	GroupId string `json:"group_id"`
	RoomId  int32  `json:"room_id"`
}

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
