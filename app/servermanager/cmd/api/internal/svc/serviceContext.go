package svc

import (
	"looklook/app/servermanager/cmd/api/internal/config"
	"looklook/app/servermanager/cmd/rpc/servermanager"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config           config.Config
	ServerManagerRpc servermanager.Servermanager
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:           c,
		ServerManagerRpc: servermanager.NewServermanager(zrpc.MustNewClient(c.ServerManagerRpcConf)),
	}
}
