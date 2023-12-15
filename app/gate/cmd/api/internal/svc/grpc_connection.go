package svc

import (
	"context"
	"fmt"

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
	m := make([]byte, 1024)

	for {
		select {
		case <-c.ctx.Done():
			return
		default:
			c.conn.RecvMsg(m)
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
