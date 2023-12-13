package svc

import (
	"automatix/app/gate/cmd/api/internal/config"
	"automatix/app/rolemanager/cmd/rpc/rolemanager"

	"github.com/bwmarrin/snowflake"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config         config.Config
	RolemanagerRpc rolemanager.Rolemanager
	Snowflake      *snowflake.Node
}

func NewServiceContext(c config.Config) *ServiceContext {
	snowflake, _ := snowflake.NewNode(c.Id)
	return &ServiceContext{
		Config:         c,
		RolemanagerRpc: rolemanager.NewRolemanager(zrpc.MustNewClient(c.RoleManagerRpcConf)),
		Snowflake:      snowflake,
	}
}
