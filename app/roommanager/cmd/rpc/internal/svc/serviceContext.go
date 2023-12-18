package svc

import "automatix/app/roommanager/cmd/rpc/internal/config"

type ServiceContext struct {
	Config    config.Config
	StreamManager newStreamManager()
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
	}
}
