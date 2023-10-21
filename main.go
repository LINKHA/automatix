package main

import (
	"context"
	"fmt"
	"time"
)

type Connection struct {
	ctx    context.Context
	cancel context.CancelFunc
}

func (c *Connection) Start() {
	c.ctx, c.cancel = context.WithCancel(context.Background())
}

func (c *Connection) Func(id int) {
	fmt.Println(fmt.Sprintf("start func: %d", id))
	go func() {
		select {
		case <-c.ctx.Done():
			fmt.Println(fmt.Sprintf("finish: %d", id))
		}
	}()
}

func (c *Connection) Cancel() {
	c.cancel()
}

func main() {
	var conn = &Connection{}
	conn.Start()
	conn.Func(1)
	conn.Func(2)
	time.Sleep(time.Duration(2) * time.Second)

	conn.Cancel()
	time.Sleep(time.Duration(100) * time.Second)
}
