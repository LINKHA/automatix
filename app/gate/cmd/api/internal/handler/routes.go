package handler

import (
	"automatix/app/gate/cmd/api/internal/svc"
)

func handle(serverCtx *svc.ServiceContext) {
	// client, _ := serverCtx.RolemanagerRpc.CreateRoleStream(context.Background())

	// go func() {
	// 	for {
	// 		select {
	// 		case data, ok := <-c.msgBuffChan:
	// 			client.Send(&pb.CreateRoleReq{
	// 				Id:        "1",
	// 				Name:      "a",
	// 				AccountId: 1,
	// 			})
	// 		case <-c.ctx.Done():
	// 			return
	// 		}

	// 	}
	// }()

	// go func() {
	// 	for {

	// 		rec, _ := client.Recv()
	// 		fmt.Printf("xxx 1------------------   :   %v\n", rec)
	// 	}
	// }()
}

func RegisterHandlers(serverCtx *svc.ServiceContext) {
	handle(serverCtx)
}
