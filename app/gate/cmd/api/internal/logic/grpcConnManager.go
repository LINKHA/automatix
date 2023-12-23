package logic

import (
	"sync"
)

type GrpcConnManager struct {
	Conns map[int32]*GrpcConn
	pLock sync.RWMutex
}

var GrpcConnMgrObj *GrpcConnManager

func NewGrpcConnManager() *GrpcConnManager {
	return &GrpcConnManager{
		Conns: make(map[int32]*GrpcConn),
	}
}

func (s *GrpcConnManager) AddConn(clientConn *GrpcConn) {
	s.pLock.Lock()
	s.Conns[clientConn.PID] = clientConn
	s.pLock.Unlock()
}

func (s *GrpcConnManager) RemoveConnByPID(pID int32) {
	s.pLock.Lock()
	delete(s.Conns, pID)
	s.pLock.Unlock()
}

func (s *GrpcConnManager) GetConnByPID(pID int32) *GrpcConn {
	s.pLock.RLock()
	defer s.pLock.RUnlock()

	return s.Conns[pID]
}
