package main

import (
	"flag"
	"fmt"

	"automatix/app/gate/cmd/api/internal/config"
	"automatix/app/gate/cmd/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
)

var configFile = flag.String("f", "etc/gate.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	if ctx == nil {
		fmt.Printf("")
	}

	s := svc.NewServer(ctx)

	// Start Service
	s.Serve()
}
