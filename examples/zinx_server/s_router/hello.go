package s_router

import (
	"github.com/linkha/automatix/pkg/net/ziface"
	"github.com/linkha/automatix/pkg/net/zlog"
	"github.com/linkha/automatix/pkg/net/znet"
)

type HelloZinxRouter struct {
	znet.BaseRouter
}

// HelloZinxRouter Handle
func (this *HelloZinxRouter) Handle(request ziface.IRequest) {
	zlog.Ins().DebugF("Call HelloZinxRouter Handle")
	// Read the data from the client first, then send back "ping...ping...ping"
	zlog.Ins().DebugF("recv from client : msgId=%d, data=%+v, len=%d", request.GetMsgID(), string(request.GetData()), len(request.GetData()))

	err := request.GetConnection().SendBuffMsg(3, []byte("Hello Zinx Router[FromServer]"))
	if err != nil {
		zlog.Error(err)
	}
}
