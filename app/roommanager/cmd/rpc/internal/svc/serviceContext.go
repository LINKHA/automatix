package svc

import "automatix/app/roommanager/cmd/rpc/internal/config"

type ServiceContext struct {
	Config    config.Config
	StreamMgr StreamManager
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
	}
}
