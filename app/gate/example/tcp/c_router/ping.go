package c_router

import (
	"github.com/LINKHA/automatix/app/gate/cmd/api/src/net/iface"
	"github.com/LINKHA/automatix/app/gate/cmd/api/src/net/net"
	"github.com/LINKHA/automatix/common/log"
)

// Ping test custom routing.
type PingRouter struct {
	net.BaseRouter
}

// Ping Handle
func (this *PingRouter) Handle(request iface.IRequest) {
	log.Debug("Call PingRouter Handle")
	log.Debug("recv from server : msgId=", request.GetMsgID(), ", data=", string(request.GetData()))

	if err := request.GetConnection().SendBuffMsg(1, []byte("Hello[from client]")); err != nil {
		log.Error(err)
	}
}
