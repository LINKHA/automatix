package server

import (
	"fmt"

	gPb "github.com/LINKHA/automatix/app/gate/cmd/api/pb"
	"github.com/LINKHA/automatix/app/gate/cmd/api/src/handler"
	"github.com/LINKHA/automatix/app/gate/cmd/api/src/net/iface"
	"github.com/LINKHA/automatix/app/gate/cmd/api/src/net/net"
	"google.golang.org/protobuf/proto"
)

var (
	RoleMgr = NewRoleManager()
)

func OnConnectionAdd(conn iface.IConnection) {
	fmt.Println("Connection is start")
}

func OnConnectionLost(conn iface.IConnection) {
	pID, _ := conn.GetProperty("pId")
	var connId int32

	if pID != nil {
		connId = pID.(int32)
	}
	fmt.Println("Connection is Lost. connId: ", connId)
	RoleMgr.RemoveByConnId(uint32(connId))
}

type GateLoginRouter struct {
	net.BaseRouter
}

func (this *GateLoginRouter) Handle(request iface.IRequest) {
	pbMsg := &gPb.GateLoginReq{}
	proto.Unmarshal(request.GetData(), pbMsg)
	connId := request.GetConnection().GetConnID()

	RoleMgr.Add(pbMsg.RoldId, uint32(connId))
}

func RunServer(ctx *handler.ServiceContext, s iface.IServer) {
	RegisterHandlers(ctx, s)

	s.SetOnConnStart(OnConnectionAdd)
	s.SetOnConnStop(OnConnectionLost)

	s.AddRouter(10001, &GateLoginRouter{})

	s.Serve()
}
