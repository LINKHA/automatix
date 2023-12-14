package svc

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"
)

type StreamClientInterface interface {
	Send(*proto.Message) error
	Recv() (*proto.Message, error)
	grpc.ClientStream
}

type GrpcConnection[T StreamClientInterface] struct {
	conn   T
	connID string

	ctx    context.Context
	cancel context.CancelFunc

	msgBuffChan chan interface{}
}

func NewGrpcConnection[T StreamClientInterface](conn T, connID string) GrpcConnection[T] {

	// Initialize Conn properties
	c := GrpcConnection[T]{
		conn:   conn,
		connID: connID,
	}

	// server.GetConnMgr().Add(c)
	return c
}

func (c *GrpcConnection[T]) Send(msg *proto.Message) error {
	return c.conn.Send(msg)
}

// func (c *GrpcConnection[T]) StartReader() {
// 	// buffer := make([]byte, 1024)

// 	// for {
// 	// 	select {
// 	// 	case <-c.ctx.Done():
// 	// 		return
// 	// 	default:
// 	// 		c.conn.Recv()
// 	// 	}
// 	// }
// }

// func (c *GrpcConnection[T]) StartWriter() {
// 	// for {
// 	// 	select {
// 	// 	case data, ok := <-c.msgBuffChan:
// 	// 		if ok {
// 	// 			if err := c.Send(data); err != nil {
// 	// 				fmt.Printf("Send Buff Data error:, %s Conn Writer exit\n", err)
// 	// 				break
// 	// 			}

// 	// 		} else {
// 	// 			fmt.Println("msgBuffChan is Closed")
// 	// 			break
// 	// 		}
// 	// 	case <-c.ctx.Done():
// 	// 		return
// 	// 	}
// 	// }
// }
