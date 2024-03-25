// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: friend.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Rolemanager_AddFriend_FullMethodName = "/pb.rolemanager/addFriend"
)

// RolemanagerClient is the client API for Rolemanager service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RolemanagerClient interface {
	AddFriend(ctx context.Context, opts ...grpc.CallOption) (Rolemanager_AddFriendClient, error)
}

type rolemanagerClient struct {
	cc grpc.ClientConnInterface
}

func NewRolemanagerClient(cc grpc.ClientConnInterface) RolemanagerClient {
	return &rolemanagerClient{cc}
}

func (c *rolemanagerClient) AddFriend(ctx context.Context, opts ...grpc.CallOption) (Rolemanager_AddFriendClient, error) {
	stream, err := c.cc.NewStream(ctx, &Rolemanager_ServiceDesc.Streams[0], Rolemanager_AddFriend_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &rolemanagerAddFriendClient{stream}
	return x, nil
}

type Rolemanager_AddFriendClient interface {
	Send(*AddFriendReq) error
	Recv() (*AddFriendResp, error)
	grpc.ClientStream
}

type rolemanagerAddFriendClient struct {
	grpc.ClientStream
}

func (x *rolemanagerAddFriendClient) Send(m *AddFriendReq) error {
	return x.ClientStream.SendMsg(m)
}

func (x *rolemanagerAddFriendClient) Recv() (*AddFriendResp, error) {
	m := new(AddFriendResp)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// RolemanagerServer is the server API for Rolemanager service.
// All implementations must embed UnimplementedRolemanagerServer
// for forward compatibility
type RolemanagerServer interface {
	AddFriend(Rolemanager_AddFriendServer) error
	mustEmbedUnimplementedRolemanagerServer()
}

// UnimplementedRolemanagerServer must be embedded to have forward compatible implementations.
type UnimplementedRolemanagerServer struct {
}

func (UnimplementedRolemanagerServer) AddFriend(Rolemanager_AddFriendServer) error {
	return status.Errorf(codes.Unimplemented, "method AddFriend not implemented")
}
func (UnimplementedRolemanagerServer) mustEmbedUnimplementedRolemanagerServer() {}

// UnsafeRolemanagerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RolemanagerServer will
// result in compilation errors.
type UnsafeRolemanagerServer interface {
	mustEmbedUnimplementedRolemanagerServer()
}

func RegisterRolemanagerServer(s grpc.ServiceRegistrar, srv RolemanagerServer) {
	s.RegisterService(&Rolemanager_ServiceDesc, srv)
}

func _Rolemanager_AddFriend_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(RolemanagerServer).AddFriend(&rolemanagerAddFriendServer{stream})
}

type Rolemanager_AddFriendServer interface {
	Send(*AddFriendResp) error
	Recv() (*AddFriendReq, error)
	grpc.ServerStream
}

type rolemanagerAddFriendServer struct {
	grpc.ServerStream
}

func (x *rolemanagerAddFriendServer) Send(m *AddFriendResp) error {
	return x.ServerStream.SendMsg(m)
}

func (x *rolemanagerAddFriendServer) Recv() (*AddFriendReq, error) {
	m := new(AddFriendReq)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Rolemanager_ServiceDesc is the grpc.ServiceDesc for Rolemanager service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Rolemanager_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.rolemanager",
	HandlerType: (*RolemanagerServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "addFriend",
			Handler:       _Rolemanager_AddFriend_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "friend.proto",
}