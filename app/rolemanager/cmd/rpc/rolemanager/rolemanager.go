// Code generated by goctl. DO NOT EDIT.
// Source: rolemanager.proto

package rolemanager

import (
	"context"

	"github.com/LINKHA/automatix/app/rolemanager/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	CreateRoleReq  = pb.CreateRoleReq
	CreateRoleResp = pb.CreateRoleResp

	Rolemanager interface {
		CreateRole(ctx context.Context, in *CreateRoleReq, opts ...grpc.CallOption) (*CreateRoleResp, error)
		CreateRoleStream(ctx context.Context, opts ...grpc.CallOption) (pb.Rolemanager_CreateRoleStreamClient, error)
	}

	defaultRolemanager struct {
		cli zrpc.Client
	}
)

func NewRolemanager(cli zrpc.Client) Rolemanager {
	return &defaultRolemanager{
		cli: cli,
	}
}

func (m *defaultRolemanager) CreateRole(ctx context.Context, in *CreateRoleReq, opts ...grpc.CallOption) (*CreateRoleResp, error) {
	client := pb.NewRolemanagerClient(m.cli.Conn())
	return client.CreateRole(ctx, in, opts...)
}

func (m *defaultRolemanager) CreateRoleStream(ctx context.Context, opts ...grpc.CallOption) (pb.Rolemanager_CreateRoleStreamClient, error) {
	client := pb.NewRolemanagerClient(m.cli.Conn())
	return client.CreateRoleStream(ctx, opts...)
}
