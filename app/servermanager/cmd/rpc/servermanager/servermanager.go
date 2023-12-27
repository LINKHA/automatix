// Code generated by goctl. DO NOT EDIT.
// Source: servermanager.proto

package servermanager

import (
	"context"

	"github.com/LINKHA/automatix/app/servermanager/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	CreateServerReq   = pb.CreateServerReq
	CreateServerResp  = pb.CreateServerResp
	EnterServerReq    = pb.EnterServerReq
	EnterServerResp   = pb.EnterServerResp
	GetServerCodeReq  = pb.GetServerCodeReq
	GetServerCodeResp = pb.GetServerCodeResp
	GetServerListReq  = pb.GetServerListReq
	GetServerListResp = pb.GetServerListResp
	GetServerReq      = pb.GetServerReq
	GetServerResp     = pb.GetServerResp
	LoginServerReq    = pb.LoginServerReq
	LoginServerResp   = pb.LoginServerResp
	Server            = pb.Server
	ServerInfo        = pb.ServerInfo
	SetServerReq      = pb.SetServerReq
	SetServerResp     = pb.SetServerResp

	Servermanager interface {
		CreateServer(ctx context.Context, in *CreateServerReq, opts ...grpc.CallOption) (*CreateServerResp, error)
		SetServer(ctx context.Context, in *SetServerReq, opts ...grpc.CallOption) (*SetServerResp, error)
		GetServer(ctx context.Context, in *GetServerReq, opts ...grpc.CallOption) (*GetServerResp, error)
		LoginServer(ctx context.Context, in *LoginServerReq, opts ...grpc.CallOption) (*LoginServerResp, error)
		GetServerList(ctx context.Context, in *GetServerListReq, opts ...grpc.CallOption) (*GetServerListResp, error)
		GetServerCode(ctx context.Context, in *GetServerCodeReq, opts ...grpc.CallOption) (*GetServerCodeResp, error)
		EnterServer(ctx context.Context, in *EnterServerReq, opts ...grpc.CallOption) (*EnterServerResp, error)
	}

	defaultServermanager struct {
		cli zrpc.Client
	}
)

func NewServermanager(cli zrpc.Client) Servermanager {
	return &defaultServermanager{
		cli: cli,
	}
}

func (m *defaultServermanager) CreateServer(ctx context.Context, in *CreateServerReq, opts ...grpc.CallOption) (*CreateServerResp, error) {
	client := pb.NewServermanagerClient(m.cli.Conn())
	return client.CreateServer(ctx, in, opts...)
}

func (m *defaultServermanager) SetServer(ctx context.Context, in *SetServerReq, opts ...grpc.CallOption) (*SetServerResp, error) {
	client := pb.NewServermanagerClient(m.cli.Conn())
	return client.SetServer(ctx, in, opts...)
}

func (m *defaultServermanager) GetServer(ctx context.Context, in *GetServerReq, opts ...grpc.CallOption) (*GetServerResp, error) {
	client := pb.NewServermanagerClient(m.cli.Conn())
	return client.GetServer(ctx, in, opts...)
}

func (m *defaultServermanager) LoginServer(ctx context.Context, in *LoginServerReq, opts ...grpc.CallOption) (*LoginServerResp, error) {
	client := pb.NewServermanagerClient(m.cli.Conn())
	return client.LoginServer(ctx, in, opts...)
}

func (m *defaultServermanager) GetServerList(ctx context.Context, in *GetServerListReq, opts ...grpc.CallOption) (*GetServerListResp, error) {
	client := pb.NewServermanagerClient(m.cli.Conn())
	return client.GetServerList(ctx, in, opts...)
}

func (m *defaultServermanager) GetServerCode(ctx context.Context, in *GetServerCodeReq, opts ...grpc.CallOption) (*GetServerCodeResp, error) {
	client := pb.NewServermanagerClient(m.cli.Conn())
	return client.GetServerCode(ctx, in, opts...)
}

func (m *defaultServermanager) EnterServer(ctx context.Context, in *EnterServerReq, opts ...grpc.CallOption) (*EnterServerResp, error) {
	client := pb.NewServermanagerClient(m.cli.Conn())
	return client.EnterServer(ctx, in, opts...)
}
