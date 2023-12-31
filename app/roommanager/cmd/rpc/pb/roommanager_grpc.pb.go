// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: roommanager.proto

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
	Roommanager_CreateRoom_FullMethodName     = "/pb.roommanager/createRoom"
	Roommanager_JoinRoom_FullMethodName       = "/pb.roommanager/joinRoom"
	Roommanager_JoinRoomStream_FullMethodName = "/pb.roommanager/joinRoomStream"
)

// RoommanagerClient is the client API for Roommanager service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RoommanagerClient interface {
	CreateRoom(ctx context.Context, in *CreateRoomReq, opts ...grpc.CallOption) (*CreateRoomResp, error)
	JoinRoom(ctx context.Context, in *JoinRoomReq, opts ...grpc.CallOption) (*JoinRoomResp, error)
	JoinRoomStream(ctx context.Context, opts ...grpc.CallOption) (Roommanager_JoinRoomStreamClient, error)
}

type roommanagerClient struct {
	cc grpc.ClientConnInterface
}

func NewRoommanagerClient(cc grpc.ClientConnInterface) RoommanagerClient {
	return &roommanagerClient{cc}
}

func (c *roommanagerClient) CreateRoom(ctx context.Context, in *CreateRoomReq, opts ...grpc.CallOption) (*CreateRoomResp, error) {
	out := new(CreateRoomResp)
	err := c.cc.Invoke(ctx, Roommanager_CreateRoom_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roommanagerClient) JoinRoom(ctx context.Context, in *JoinRoomReq, opts ...grpc.CallOption) (*JoinRoomResp, error) {
	out := new(JoinRoomResp)
	err := c.cc.Invoke(ctx, Roommanager_JoinRoom_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roommanagerClient) JoinRoomStream(ctx context.Context, opts ...grpc.CallOption) (Roommanager_JoinRoomStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &Roommanager_ServiceDesc.Streams[0], Roommanager_JoinRoomStream_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &roommanagerJoinRoomStreamClient{stream}
	return x, nil
}

type Roommanager_JoinRoomStreamClient interface {
	Send(*JoinRoomStreamReq) error
	Recv() (*JoinRoomStreamResp, error)
	grpc.ClientStream
}

type roommanagerJoinRoomStreamClient struct {
	grpc.ClientStream
}

func (x *roommanagerJoinRoomStreamClient) Send(m *JoinRoomStreamReq) error {
	return x.ClientStream.SendMsg(m)
}

func (x *roommanagerJoinRoomStreamClient) Recv() (*JoinRoomStreamResp, error) {
	m := new(JoinRoomStreamResp)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// RoommanagerServer is the server API for Roommanager service.
// All implementations must embed UnimplementedRoommanagerServer
// for forward compatibility
type RoommanagerServer interface {
	CreateRoom(context.Context, *CreateRoomReq) (*CreateRoomResp, error)
	JoinRoom(context.Context, *JoinRoomReq) (*JoinRoomResp, error)
	JoinRoomStream(Roommanager_JoinRoomStreamServer) error
	mustEmbedUnimplementedRoommanagerServer()
}

// UnimplementedRoommanagerServer must be embedded to have forward compatible implementations.
type UnimplementedRoommanagerServer struct {
}

func (UnimplementedRoommanagerServer) CreateRoom(context.Context, *CreateRoomReq) (*CreateRoomResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRoom not implemented")
}
func (UnimplementedRoommanagerServer) JoinRoom(context.Context, *JoinRoomReq) (*JoinRoomResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method JoinRoom not implemented")
}
func (UnimplementedRoommanagerServer) JoinRoomStream(Roommanager_JoinRoomStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method JoinRoomStream not implemented")
}
func (UnimplementedRoommanagerServer) mustEmbedUnimplementedRoommanagerServer() {}

// UnsafeRoommanagerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RoommanagerServer will
// result in compilation errors.
type UnsafeRoommanagerServer interface {
	mustEmbedUnimplementedRoommanagerServer()
}

func RegisterRoommanagerServer(s grpc.ServiceRegistrar, srv RoommanagerServer) {
	s.RegisterService(&Roommanager_ServiceDesc, srv)
}

func _Roommanager_CreateRoom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRoomReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoommanagerServer).CreateRoom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Roommanager_CreateRoom_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoommanagerServer).CreateRoom(ctx, req.(*CreateRoomReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Roommanager_JoinRoom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JoinRoomReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoommanagerServer).JoinRoom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Roommanager_JoinRoom_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoommanagerServer).JoinRoom(ctx, req.(*JoinRoomReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Roommanager_JoinRoomStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(RoommanagerServer).JoinRoomStream(&roommanagerJoinRoomStreamServer{stream})
}

type Roommanager_JoinRoomStreamServer interface {
	Send(*JoinRoomStreamResp) error
	Recv() (*JoinRoomStreamReq, error)
	grpc.ServerStream
}

type roommanagerJoinRoomStreamServer struct {
	grpc.ServerStream
}

func (x *roommanagerJoinRoomStreamServer) Send(m *JoinRoomStreamResp) error {
	return x.ServerStream.SendMsg(m)
}

func (x *roommanagerJoinRoomStreamServer) Recv() (*JoinRoomStreamReq, error) {
	m := new(JoinRoomStreamReq)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Roommanager_ServiceDesc is the grpc.ServiceDesc for Roommanager service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Roommanager_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.roommanager",
	HandlerType: (*RoommanagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "createRoom",
			Handler:    _Roommanager_CreateRoom_Handler,
		},
		{
			MethodName: "joinRoom",
			Handler:    _Roommanager_JoinRoom_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "joinRoomStream",
			Handler:       _Roommanager_JoinRoomStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "roommanager.proto",
}
