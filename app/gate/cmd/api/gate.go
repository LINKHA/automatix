package main

import (
	"flag"
	"fmt"

	"automatix/app/gate/cmd/api/internal/config"
	"automatix/app/gate/cmd/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
)

var configFile = flag.String("f", "etc/gate-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	// server := rest.MustNewServer(c.RestConf)
	// defer server.Stop()

	ctx := svc.NewServiceContext(c)
	// handler.RegisterHandlers(server, ctx)

	svc.NewGateServer(ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.TcpHost, c.TcpPort)
	if ctx == nil {
		fmt.Printf("")
	}
	// server.Start()
}
