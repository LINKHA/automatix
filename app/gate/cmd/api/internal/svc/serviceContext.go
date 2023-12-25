package svc

import (
	"automatix/app/gate/cmd/api/internal/config"
	"automatix/app/gate/cmd/api/internal/handler"
	"automatix/app/rolemanager/cmd/rpc/rolemanager"
	"automatix/app/roommanager/cmd/rpc/roommanager"

	"github.com/bwmarrin/snowflake"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config         config.Config
	RoleManagerRpc rolemanager.Rolemanager
	RoomManagerRpc roommanager.Roommanager
	Snowflake      *snowflake.Node
	MsgHandler     *handler.MsgHandler
}

var ServiceContextObj *ServiceContext

func NewServiceContext(c config.Config) *ServiceContext {
	snowflake, _ := snowflake.NewNode(c.Id)
	msgHandler := handler.NewMsgHandler()

	ServiceContextObj = &ServiceContext{
		Config:         c,
		RoleManagerRpc: rolemanager.NewRolemanager(zrpc.MustNewClient(c.RoleManagerRpcConf)),
		RoomManagerRpc: roommanager.NewRoommanager(zrpc.MustNewClient(c.RoomManagerRpcConf)),
		Snowflake:      snowflake,
		MsgHandler:     msgHandler,
	}
	return ServiceContextObj
}
