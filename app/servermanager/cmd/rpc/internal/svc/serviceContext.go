package svc

import (
	"looklook/app/servermanager/cmd/rpc/internal/config"
	"looklook/app/servermanager/model"
)

type ServiceContext struct {
	Config config.Config

	ServerModel model.ServerModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
	}
}
