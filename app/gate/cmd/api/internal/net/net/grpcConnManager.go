package net

import (
	"automatix/app/gate/cmd/api/internal/net/iface"
	"sync"
)

type GrpcConnManager struct {
	Conns map[string]iface.IGrpcConnection
	pLock sync.RWMutex
}

var GrpcConnMgrObj *GrpcConnManager

func newGrpcConnManager() *GrpcConnManager {
	return &GrpcConnManager{
		Conns: make(map[string]iface.IGrpcConnection),
	}
}

func (s *GrpcConnManager) Add(clientConn iface.IGrpcConnection) {
	s.pLock.Lock()
	s.Conns[clientConn.GetConnIdStr()] = clientConn
	s.pLock.Unlock()
}

func (s *GrpcConnManager) Remove(pID string) {
	s.pLock.Lock()
	delete(s.Conns, pID)
	s.pLock.Unlock()
}

func (s *GrpcConnManager) Get(pID string) iface.IGrpcConnection {
	s.pLock.RLock()
	defer s.pLock.RUnlock()

	return s.Conns[pID]
}
