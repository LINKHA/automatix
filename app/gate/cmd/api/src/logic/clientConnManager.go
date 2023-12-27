package logic

import "sync"

type ClientConnManager struct {
	Conns map[int32]*ClientConn
	pLock sync.RWMutex
}

var ClientConnMgrObj *ClientConnManager

func NewClientConnManager() *ClientConnManager {
	return &ClientConnManager{
		Conns: make(map[int32]*ClientConn),
	}
}

func (s *ClientConnManager) AddConn(clientConn *ClientConn) {
	s.pLock.Lock()
	s.Conns[clientConn.PID] = clientConn
	s.pLock.Unlock()
}

func (s *ClientConnManager) RemoveConnByPID(pID int32) {
	s.pLock.Lock()
	delete(s.Conns, pID)
	s.pLock.Unlock()
}

func (s *ClientConnManager) GetByPID(pID int32) *ClientConn {
	s.pLock.RLock()
	defer s.pLock.RUnlock()

	return s.Conns[pID]
}
