package main

import (
	"flag"

	"github.com/LINKHA/automatix/app/gate/cmd/api/src/config"
	"github.com/LINKHA/automatix/app/gate/cmd/api/src/handler"
	"github.com/LINKHA/automatix/app/gate/cmd/api/src/server"

	"github.com/LINKHA/automatix/app/gate/cmd/api/src/net/net"

	"github.com/zeromicro/go-zero/core/conf"
)

var configFile = flag.String("f", "etc/gate.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := handler.NewServiceContext(c)

	s := net.NewServer()

	server.RunServer(ctx, s)

}
