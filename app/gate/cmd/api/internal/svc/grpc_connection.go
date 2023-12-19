package svc

import (
	"context"
	"fmt"
	"io"

	"google.golang.org/grpc"
)

type StreamClientInterface interface {
	grpc.ClientStream
}

type GrpcConnection[T1 StreamClientInterface, T2 any, T3 any] struct {
	conn   T1
	connID string

	ctx    context.Context
	cancel context.CancelFunc

	msgReqChan  chan T2
	msgRespChan chan T3
}

func NewGrpcConnection[T1 StreamClientInterface, T2 any, T3 any](ctx *ServiceContext, conn T1, connID string) GrpcConnection[T1, T2, T3] {

	// Initialize Conn properties
	c := GrpcConnection[T1, T2, T3]{
		conn:        conn,
		connID:      connID,
		msgReqChan:  make(chan T2, 1000),
		msgRespChan: make(chan T3, 1000),
	}

	ctx.StreamManager.Set(c.GetConnID(), c)
	return c
}

func (c *GrpcConnection[T1, T2, T3]) GetConnID() string {
	return c.connID
}

func (c *GrpcConnection[T1, T2, T3]) Start() {
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

func (c *GrpcConnection[T1, T2, T3]) Stop() {
	c.cancel()
}

func (c *GrpcConnection[T1, T2, T3]) Send(m T2) error {
	return c.conn.SendMsg(m)
}

func (c *GrpcConnection[T1, T2, T3]) Recv(m T3) (interface{}, error) {
	if err := c.conn.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// bug note
func (c *GrpcConnection[T1, T2, T3]) StartReader() {
	m := new(T3)
	for {
		select {
		case <-c.ctx.Done():
			return
		default:
			//NOTE: 这里如果rpc服务没起来会报错
			err := c.conn.RecvMsg(m)

			if err == io.EOF {
				fmt.Printf("Received GRPC EOF\n")
			}

			if err != nil {
				fmt.Printf("Received GRPC err: %s\n", err)
			}

			fmt.Printf("Received GRPC message: %v\n", m)
		}
	}
}

func (c *GrpcConnection[T1, T2, T3]) StartWriter() {
	for {
		select {
		case data, ok := <-c.msgReqChan:
			if ok {
				if err := c.Send(data); err != nil {
					fmt.Printf("Send Buff Data error:, %s Conn Writer exit\n", err)
					break
				}

			} else {
				fmt.Println("msgReqChan is Closed")
				break
			}
		case <-c.ctx.Done():
			return
		}
	}
}
