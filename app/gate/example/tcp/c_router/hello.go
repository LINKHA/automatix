package c_router

import (
	"github.com/LINKHA/automatix/app/gate/cmd/api/src/net/iface"
	"github.com/LINKHA/automatix/app/gate/cmd/api/src/net/net"
	"github.com/LINKHA/automatix/common/log"
)

type HelloRouter struct {
	net.BaseRouter
}

// HelloZinxRouter Handle
func (this *HelloRouter) Handle(request iface.IRequest) {
	log.Debug("Call HelloZinxRouter Handle")

	log.Debug("recv from server : msgId=", request.GetMsgID(), ", data=", string(request.GetData()))
}
