package svc

import (
	"automatix/app/rolemanager/cmd/rpc/pb"
	"context"
	"fmt"
	"time"

	"google.golang.org/grpc"
)

type StreamClientInterface interface {
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

func (c *GrpcConnection[T]) GetConnID() string {
	return c.connID
}

func (c *GrpcConnection[T]) Start() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("Connection Start() error: %v\n", err)
		}
	}()
	c.ctx, c.cancel = context.WithCancel(context.Background())

	go c.StartReader()
	go c.StartWriter()

	select {
	case <-c.ctx.Done():
		return
	}
}

func (c *GrpcConnection[T]) Stop() {
	c.cancel()
}

func (c *GrpcConnection[T]) Send(m interface{}) error {
	return c.conn.SendMsg(m)
}

func (c *GrpcConnection[T]) Recv(m interface{}) (interface{}, error) {
	if err := c.conn.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// bug note
func (c *GrpcConnection[T]) StartReader() {
	// m := make([]proto.Message, 1024)
	m := new(pb.CreateRoleResp)
	for {
		select {
		case <-c.ctx.Done():
			return
		default:
			time.Sleep(time.Duration(1) * time.Second)
			err := c.conn.RecvMsg(m)
			fmt.Printf("Received GRPC err: %s\n", err)
			fmt.Printf("Received GRPC message: %v\n", m)
		}
	}
}

func (c *GrpcConnection[T]) StartWriter() {
	for {
		select {
		case data, ok := <-c.msgBuffChan:
			if ok {
				if err := c.Send(data); err != nil {
					fmt.Printf("Send Buff Data error:, %s Conn Writer exit\n", err)
					break
				}

			} else {
				fmt.Println("msgBuffChan is Closed")
				break
			}
		case <-c.ctx.Done():
			return
		}
	}
}
