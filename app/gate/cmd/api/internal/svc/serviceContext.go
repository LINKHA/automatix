package svc

import (
	"automatix/app/gate/cmd/api/internal/config"
	"automatix/app/rolemanager/cmd/rpc/rolemanager"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config         config.Config
	RolemanagerRpc rolemanager.Rolemanager
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:         c,
		RolemanagerRpc: rolemanager.NewRolemanager(zrpc.MustNewClient(c.RoleManagerRpcConf)),
	}
}
