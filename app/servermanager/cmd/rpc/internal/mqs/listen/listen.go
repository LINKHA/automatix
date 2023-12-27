package listen

import (
	"context"

	"github.com/LINKHA/automatix/app/servermanager/cmd/rpc/internal/config"
	"github.com/LINKHA/automatix/app/servermanager/cmd/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/service"
)

// back to all consumers
func Mqs(c config.Config) []service.Service {

	svcContext := svc.NewServiceContext(c)
	ctx := context.Background()

	var services []service.Service

	//kq ：pub sub
	services = append(services, KqMqs(c, ctx, svcContext)...)

	return services
}
