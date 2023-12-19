package svc

import (
	"automatix/app/roommanager/cmd/rpc/internal/config"
	"automatix/common/utils"
)

type ServiceContext struct {
	Config        config.Config
	StreamManager *utils.ShardLockMaps
}

func NewServiceContext(c config.Config) *ServiceContext {
	streamManager := utils.NewShardLockMaps()

	return &ServiceContext{
		Config:        c,
		StreamManager: &streamManager,
	}
}
