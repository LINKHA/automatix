package listen

import (
	"context"

	"github.com/LINKHA/automatix/app/servermanager/cmd/rpc/internal/config"
	kqMq "github.com/LINKHA/automatix/app/servermanager/cmd/rpc/internal/mqs/kq"
	"github.com/LINKHA/automatix/app/servermanager/cmd/rpc/internal/svc"

	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/service"
)

// pub sub use kq (kafka)
func KqMqs(c config.Config, ctx context.Context, svcContext *svc.ServiceContext) []service.Service {

	return []service.Service{
		//Listening for changes in consumption flow status
		kq.MustNewQueue(c.KqConfServerQueue, kqMq.NewServerQueueMq(ctx, svcContext)),
		//.....
	}

}
