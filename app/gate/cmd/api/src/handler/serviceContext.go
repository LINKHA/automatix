package handler

import (
	"github.com/LINKHA/automatix/app/gate/cmd/api/src/config"
	"github.com/LINKHA/automatix/app/rolemanager/cmd/rpc/rolemanager"
	"github.com/LINKHA/automatix/app/roommanager/cmd/rpc/roommanager"

	"github.com/bwmarrin/snowflake"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config         config.Config
	RoleManagerRpc rolemanager.Rolemanager
	RoomManagerRpc roommanager.Roommanager
	Snowflake      *snowflake.Node
}

var ServiceContextObj *ServiceContext

func NewServiceContext(c config.Config) *ServiceContext {
	snowflake, _ := snowflake.NewNode(c.Id)

	ServiceContextObj = &ServiceContext{
		Config:         c,
		RoleManagerRpc: rolemanager.NewRolemanager(zrpc.MustNewClient(c.RoleManagerRpcConf)),
		RoomManagerRpc: roommanager.NewRoommanager(zrpc.MustNewClient(c.RoomManagerRpcConf)),
		Snowflake:      snowflake,
	}
	return ServiceContextObj
}
