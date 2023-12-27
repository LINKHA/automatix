package handler

import (
	"fmt"

	"github.com/LINKHA/automatix/app/gate/cmd/api/src/logic"
)

type MsgHandler struct {
	GrpcConnManager   *logic.GrpcConnManager
	ClientConnManager *logic.ClientConnManager

	grpcTaskQueue   chan GrpcMessage
	clientTaskQueue chan ClientMessage
}

func NewMsgHandler() *MsgHandler {
	grpcConnManager := logic.NewGrpcConnManager()
	clientConnManager := logic.NewClientConnManager()

	return &MsgHandler{
		GrpcConnManager:   grpcConnManager,
		ClientConnManager: clientConnManager,
		grpcTaskQueue:     make(chan GrpcMessage, 10),
		clientTaskQueue:   make(chan ClientMessage, 10),
	}
}

func (s *MsgHandler) HandleGrpc(msg GrpcMessage) {
	s.grpcTaskQueue <- msg
}

func (s *MsgHandler) HandleClient(msg ClientMessage) {
	s.clientTaskQueue <- msg
}

func (s *MsgHandler) doHandleGrpc(msg GrpcMessage) {
	clientId := 1
	clientMsg := s.ClientConnManager.Conns[int32(clientId)]

	fmt.Println("client: ", clientId, " clientMsg: ", clientMsg)
}

func (s *MsgHandler) doHandleClient(msg ClientMessage) {
	grpcId := 1
	grpcMsg := s.GrpcConnManager.Conns[int32(grpcId)]

	fmt.Println("grpcId: ", grpcId, " grpcMsg: ", grpcMsg)
}

func (s *MsgHandler) StartOneWorker() {
	go func() {
		for {
			select {
			case grpcTaskMsg := <-s.grpcTaskQueue:
				s.doHandleGrpc(grpcTaskMsg)
			case clientTaskMsg := <-s.clientTaskQueue:
				s.doHandleClient(clientTaskMsg)
			}
		}
	}()
}
