// Code generated by goctl. DO NOT EDIT.
// Source: rolemanager.proto

package server

import (
	"context"

	"github.com/LINKHA/automatix/app/rolemanager/cmd/rpc/internal/logic"
	"github.com/LINKHA/automatix/app/rolemanager/cmd/rpc/internal/svc"
	"github.com/LINKHA/automatix/app/rolemanager/cmd/rpc/pb"
)

type RolemanagerServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedRolemanagerServer
}

func NewRolemanagerServer(svcCtx *svc.ServiceContext) *RolemanagerServer {
	return &RolemanagerServer{
		svcCtx: svcCtx,
	}
}

func (s *RolemanagerServer) CreateRole(ctx context.Context, in *pb.CreateRoleReq) (*pb.CreateRoleResp, error) {
	l := logic.NewCreateRoleLogic(ctx, s.svcCtx)
	return l.CreateRole(in)
}

func (s *RolemanagerServer) CreateRoleStream(stream pb.Rolemanager_CreateRoleStreamServer) error {
	l := logic.NewCreateRoleStreamLogic(stream.Context(), s.svcCtx)
	return l.CreateRoleStream(stream)
}
