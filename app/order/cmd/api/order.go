package main

import (
	"flag"

	"github.com/LINKHA/automatix/app/order/cmd/api/internal/config"
	"github.com/LINKHA/automatix/app/order/cmd/api/internal/handler"
	"github.com/LINKHA/automatix/app/order/cmd/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/order.yaml", "the config file")

func main() {

	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	ctx := svc.NewServiceContext(c)
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	handler.RegisterHandlers(server, ctx)

	server.Start()
}
