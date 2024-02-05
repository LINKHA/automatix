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
	client := mNet.NewClient("127.0.0.1", 8999)
	client.SetOnConnStart(func(conn iface.IConnection) {
		go func() {
			data := &pb.RegisterRoleReq{
				AccountId:     "1",
				ServerId:      "2",
				TemplateValue: "{}",
			}
			msg, _ := proto.Marshal(data)

			//for {
			err := conn.SendMsg(101, msg)
			if err != nil {
				fmt.Println(err)
				log.Error(err)
				//break
			}

			// 	time.Sleep(1 * time.Second)
			// }
		}()
	})

	client.Start()

	// // close
	// c := make(chan os.Signal, 1)
	// signal.Notify(c, os.Interrupt, os.Kill)
	// sig := <-c
	// fmt.Println("===exit===", sig)
	// client.Stop()
	time.Sleep(time.Second * 2)
}

// func TestDeleteRole(t *testing.T) {
// 	client := mNet.NewClient("127.0.0.1", 8999)
// 	client.SetOnConnStart(func(conn iface.IConnection) {
// 		go func() {
// 			data := &pb.DeleteRoleReq{
// 				RoleId: "1753983203577696256",
// 			}
// 			msg, _ := proto.Marshal(data)

// 			//for {
// 			err := conn.SendMsg(104, msg)
// 			if err != nil {
// 				fmt.Println(err)
// 				log.Error(err)
// 				//break
// 			}

// 			// 	time.Sleep(1 * time.Second)
// 			// }
// 		}()
// 	})

// 	client.Start()

// 	// // close
// 	// c := make(chan os.Signal, 1)
// 	// signal.Notify(c, os.Interrupt, os.Kill)
// 	// sig := <-c
// 	// fmt.Println("===exit===", sig)
// 	// client.Stop()
// 	time.Sleep(time.Second * 2)
// }

// func TestSetRole(t *testing.T) {
// 	client := mNet.NewClient("127.0.0.1", 8999)
// 	client.SetOnConnStart(func(conn iface.IConnection) {
// 		go func() {
// 			tmp := make(map[string]interface{})
// 			tmp["level"] = 1
// 			templateValue, _ := json.Marshal(tmp)

// 			data := &pb.SetRoleReq{
// 				RoleId:        "1754055433376501760",
// 				TemplateValue: string(templateValue),
// 			}
// 			msg, _ := proto.Marshal(data)

// 			err := conn.SendMsg(102, msg)
// 			if err != nil {
// 				fmt.Println(err)
// 				log.Error(err)
// 			}
// 		}()
// 	})

// 	client.Start()

// 	time.Sleep(time.Second * 2)
// }

// func TestGetRole(t *testing.T) {
// 	client := mNet.NewClient("127.0.0.1", 8999)
// 	client.SetOnConnStart(func(conn iface.IConnection) {
// 		go func() {
// 			data := &pb.GetRoleReq{
// 				RoleId: "1754055433376501760",
// 			}
// 			msg, _ := proto.Marshal(data)

// 			err := conn.SendMsg(103, msg)
// 			if err != nil {
// 				fmt.Println(err)
// 				log.Error(err)
// 			}
// 		}()
// 	})

// 	client.Start()

// 	time.Sleep(time.Second * 2)
// }
