package main

import (
	"flag"
	"fmt"

	"github.com/LINKHA/automatix/app/gate/cmd/api/src/config"
	"github.com/LINKHA/automatix/app/gate/cmd/api/src/handler"

	"github.com/LINKHA/automatix/app/gate/cmd/api/src/net/iface"
	"github.com/LINKHA/automatix/app/gate/cmd/api/src/net/net"

	"github.com/zeromicro/go-zero/core/conf"
)

var configFile = flag.String("f", "etc/gate.yaml", "the config file")

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
}

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := handler.NewServiceContext(c)

	s := net.NewServer()

	handler.RegisterHandlers(ctx, s)

	s.SetOnConnStart(OnConnectionAdd)
	s.SetOnConnStop(OnConnectionLost)

	s.Serve()
}
