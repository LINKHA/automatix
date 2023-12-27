package c_router

import (
	"automatix/app/gate/cmd/api/internal/net/iface"
	"automatix/app/gate/cmd/api/internal/net/net"
	"automatix/common/log"
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
