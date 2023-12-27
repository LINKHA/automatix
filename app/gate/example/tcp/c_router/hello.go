package c_router

import (
	"automatix/app/gate/cmd/api/internal/net/iface"
	"automatix/app/gate/cmd/api/internal/net/net"
	"automatix/common/log"
)

type HelloRouter struct {
	net.BaseRouter
}

// HelloZinxRouter Handle
func (this *HelloRouter) Handle(request iface.IRequest) {
	log.Debug("Call HelloZinxRouter Handle")

	log.Debug("recv from server : msgId=", request.GetMsgID(), ", data=", string(request.GetData()))
}
