package test

import (
	"fmt"
	"testing"
	"time"

	"github.com/LINKHA/automatix/app/gate/cmd/api/src/net/iface"
	mNet "github.com/LINKHA/automatix/app/gate/cmd/api/src/net/net"
	"github.com/LINKHA/automatix/app/rolemanager/cmd/rpc/pb"
	"github.com/LINKHA/automatix/common/log"
	"google.golang.org/protobuf/proto"
)

func TestRegisterRole(t *testing.T) {
	client := mNet.NewClient("127.0.0.1", 10242)
	client.SetOnConnStart(func(conn iface.IConnection) {
		go func() {
			data := &pb.RegisterRoleReq{
				AccountId:     "1",
				ServerId:      "2",
				TemplateValue: "{}",
			}
			msg, _ := proto.Marshal(data)

			for {
				err := conn.SendMsg(102, msg)
				if err != nil {
					fmt.Println(err)
					log.Error(err)
					break
				}

				time.Sleep(1 * time.Second)
			}
		}()
	})
}