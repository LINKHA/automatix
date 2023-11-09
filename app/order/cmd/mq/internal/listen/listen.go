package listen

import (
	"automatix/app/order/cmd/mq/internal/config"
	"automatix/app/order/cmd/mq/internal/svc"
	"context"

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
