package svc

import (
	"github.com/LINKHA/automatix/app/login/cmd/api/internal/config"
	"github.com/LINKHA/automatix/app/servermanager/cmd/rpc/servermanager"
	"github.com/LINKHA/automatix/app/usercenter/cmd/rpc/usercenter"

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
