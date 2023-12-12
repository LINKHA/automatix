package net

import "net"

type GrpcConnection struct {
	conn      net.Conn
	connID    uint64
	connIdStr string
}
