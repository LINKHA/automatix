package net

import (
	"net"
	"strconv"
)

type Connection struct {
	conn      net.Conn
	connID    uint64
	connIdStr string
}

func newServerConn(server Server, conn net.Conn, connID uint64) Connection {

	// Initialize Conn properties
	c := &Connection{
		conn:      conn,
		connID:    connID,
		connIdStr: strconv.FormatUint(connID, 10),
	}
	return *c
}

func (c *Connection) GetConnID() uint64 {
	return c.connID
}

func (c *Connection) GetConnIdStr() string {
	return c.connIdStr
}
