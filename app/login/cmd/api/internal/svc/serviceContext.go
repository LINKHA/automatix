package svc

import (
	"automatix/app/login/cmd/api/internal/config"
	"automatix/app/servermanager/cmd/rpc/servermanager"
	"automatix/app/usercenter/cmd/rpc/usercenter"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config           config.Config
	UsercenterRpc    usercenter.Usercenter
	ServerManagerRpc servermanager.Servermanager
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:           c,
		UsercenterRpc:    usercenter.NewUsercenter(zrpc.MustNewClient(c.UsercenterRpcConf)),
		ServerManagerRpc: servermanager.NewServermanager(zrpc.MustNewClient(c.ServerManagerRpcConf)),
	}
}
