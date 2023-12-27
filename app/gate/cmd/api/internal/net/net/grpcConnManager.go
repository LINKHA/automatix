package net

import (
	"automatix/app/gate/cmd/api/internal/net/iface"
	"automatix/common/utils"
	"errors"
	"strconv"
)

type GrpcConnManager struct {
	connections utils.ShardLockMaps
}

var GrpcConnMgrObj *GrpcConnManager

func newGrpcConnManager() *GrpcConnManager {
	return &GrpcConnManager{
		connections: utils.NewShardLockMaps(),
	}
}

func (s *GrpcConnManager) Add(conn iface.IGrpcConnection) {
	s.connections.Set(conn.GetConnIdStr(), conn)
}

func (s *GrpcConnManager) Remove(conn iface.IGrpcConnection) {
	s.connections.Remove(conn.GetConnIdStr())
}

func (s *GrpcConnManager) Get(connID uint64) (iface.IGrpcConnection, error) {
	strConnId := strconv.FormatUint(connID, 10)
	if conn, ok := s.connections.Get(strConnId); ok {
		return conn.(iface.IGrpcConnection), nil
	}

	return nil, errors.New("connection not found")
}

func (s *GrpcConnManager) Get2(strConnId string) (iface.IGrpcConnection, error) {
	if conn, ok := s.connections.Get(strConnId); ok {
		return conn.(iface.IGrpcConnection), nil
	}

	return nil, errors.New("connection not found")
}
