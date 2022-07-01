// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.1
// source: connection-service.proto

package connection

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

// ConnectionServiceClient is the client API for ConnectionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ConnectionServiceClient interface {
	MakeConnectionWithPublicProfile(ctx context.Context, in *ConnectionBody, opts ...grpc.CallOption) (*ConnectionResponse, error)
	MakeConnectionRequest(ctx context.Context, in *ConnectionBody, opts ...grpc.CallOption) (*ConnectionResponse, error)
	ApproveConnectionRequest(ctx context.Context, in *ConnectionBody, opts ...grpc.CallOption) (*ConnectionResponse, error)
	GetConnectionsUsernamesFor(ctx context.Context, in *GetConnectionsUsernamesRequest, opts ...grpc.CallOption) (*GetConnectionsUsernamesResponse, error)
	GetRequestsUsernamesFor(ctx context.Context, in *GetConnectionsUsernamesRequest, opts ...grpc.CallOption) (*GetConnectionsUsernamesResponse, error)
	InsertUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*Status, error)
	GetSuggestionIdsFor(ctx context.Context, in *GetSuggestionIdsRequest, opts ...grpc.CallOption) (*GetSuggestionIdsResponse, error)
	BlockConnection(ctx context.Context, in *ConnectionBody, opts ...grpc.CallOption) (*ConnectionResponse, error)
}

type connectionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewConnectionServiceClient(cc grpc.ClientConnInterface) ConnectionServiceClient {
	return &connectionServiceClient{cc}
}

func (c *connectionServiceClient) MakeConnectionWithPublicProfile(ctx context.Context, in *ConnectionBody, opts ...grpc.CallOption) (*ConnectionResponse, error) {
	out := new(ConnectionResponse)
	err := c.cc.Invoke(ctx, "/connection.ConnectionService/MakeConnectionWithPublicProfile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *connectionServiceClient) MakeConnectionRequest(ctx context.Context, in *ConnectionBody, opts ...grpc.CallOption) (*ConnectionResponse, error) {
	out := new(ConnectionResponse)
	err := c.cc.Invoke(ctx, "/connection.ConnectionService/MakeConnectionRequest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *connectionServiceClient) ApproveConnectionRequest(ctx context.Context, in *ConnectionBody, opts ...grpc.CallOption) (*ConnectionResponse, error) {
	out := new(ConnectionResponse)
	err := c.cc.Invoke(ctx, "/connection.ConnectionService/ApproveConnectionRequest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *connectionServiceClient) GetConnectionsUsernamesFor(ctx context.Context, in *GetConnectionsUsernamesRequest, opts ...grpc.CallOption) (*GetConnectionsUsernamesResponse, error) {
	out := new(GetConnectionsUsernamesResponse)
	err := c.cc.Invoke(ctx, "/connection.ConnectionService/GetConnectionsUsernamesFor", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *connectionServiceClient) GetRequestsUsernamesFor(ctx context.Context, in *GetConnectionsUsernamesRequest, opts ...grpc.CallOption) (*GetConnectionsUsernamesResponse, error) {
	out := new(GetConnectionsUsernamesResponse)
	err := c.cc.Invoke(ctx, "/connection.ConnectionService/GetRequestsUsernamesFor", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *connectionServiceClient) InsertUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/connection.ConnectionService/InsertUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *connectionServiceClient) GetSuggestionIdsFor(ctx context.Context, in *GetSuggestionIdsRequest, opts ...grpc.CallOption) (*GetSuggestionIdsResponse, error) {
	out := new(GetSuggestionIdsResponse)
	err := c.cc.Invoke(ctx, "/connection.ConnectionService/GetSuggestionIdsFor", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *connectionServiceClient) BlockConnection(ctx context.Context, in *ConnectionBody, opts ...grpc.CallOption) (*ConnectionResponse, error) {
	out := new(ConnectionResponse)
	err := c.cc.Invoke(ctx, "/connection.ConnectionService/BlockConnection", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ConnectionServiceServer is the server API for ConnectionService service.
// All implementations must embed UnimplementedConnectionServiceServer
// for forward compatibility
type ConnectionServiceServer interface {
	MakeConnectionWithPublicProfile(context.Context, *ConnectionBody) (*ConnectionResponse, error)
	MakeConnectionRequest(context.Context, *ConnectionBody) (*ConnectionResponse, error)
	ApproveConnectionRequest(context.Context, *ConnectionBody) (*ConnectionResponse, error)
	GetConnectionsUsernamesFor(context.Context, *GetConnectionsUsernamesRequest) (*GetConnectionsUsernamesResponse, error)
	GetRequestsUsernamesFor(context.Context, *GetConnectionsUsernamesRequest) (*GetConnectionsUsernamesResponse, error)
	InsertUser(context.Context, *User) (*Status, error)
	GetSuggestionIdsFor(context.Context, *GetSuggestionIdsRequest) (*GetSuggestionIdsResponse, error)
	BlockConnection(context.Context, *ConnectionBody) (*ConnectionResponse, error)
	mustEmbedUnimplementedConnectionServiceServer()
}

// UnimplementedConnectionServiceServer must be embedded to have forward compatible implementations.
type UnimplementedConnectionServiceServer struct {
}

func (UnimplementedConnectionServiceServer) MakeConnectionWithPublicProfile(context.Context, *ConnectionBody) (*ConnectionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MakeConnectionWithPublicProfile not implemented")
}
func (UnimplementedConnectionServiceServer) MakeConnectionRequest(context.Context, *ConnectionBody) (*ConnectionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MakeConnectionRequest not implemented")
}
func (UnimplementedConnectionServiceServer) ApproveConnectionRequest(context.Context, *ConnectionBody) (*ConnectionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ApproveConnectionRequest not implemented")
}
func (UnimplementedConnectionServiceServer) GetConnectionsUsernamesFor(context.Context, *GetConnectionsUsernamesRequest) (*GetConnectionsUsernamesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetConnectionsUsernamesFor not implemented")
}
func (UnimplementedConnectionServiceServer) GetRequestsUsernamesFor(context.Context, *GetConnectionsUsernamesRequest) (*GetConnectionsUsernamesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRequestsUsernamesFor not implemented")
}
func (UnimplementedConnectionServiceServer) InsertUser(context.Context, *User) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InsertUser not implemented")
}
func (UnimplementedConnectionServiceServer) GetSuggestionIdsFor(context.Context, *GetSuggestionIdsRequest) (*GetSuggestionIdsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSuggestionIdsFor not implemented")
}
func (UnimplementedConnectionServiceServer) BlockConnection(context.Context, *ConnectionBody) (*ConnectionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BlockConnection not implemented")
}
func (UnimplementedConnectionServiceServer) mustEmbedUnimplementedConnectionServiceServer() {}

// UnsafeConnectionServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ConnectionServiceServer will
// result in compilation errors.
type UnsafeConnectionServiceServer interface {
	mustEmbedUnimplementedConnectionServiceServer()
}

func RegisterConnectionServiceServer(s grpc.ServiceRegistrar, srv ConnectionServiceServer) {
	s.RegisterService(&ConnectionService_ServiceDesc, srv)
}

func _ConnectionService_MakeConnectionWithPublicProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConnectionBody)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectionServiceServer).MakeConnectionWithPublicProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/connection.ConnectionService/MakeConnectionWithPublicProfile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectionServiceServer).MakeConnectionWithPublicProfile(ctx, req.(*ConnectionBody))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConnectionService_MakeConnectionRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConnectionBody)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectionServiceServer).MakeConnectionRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/connection.ConnectionService/MakeConnectionRequest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectionServiceServer).MakeConnectionRequest(ctx, req.(*ConnectionBody))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConnectionService_ApproveConnectionRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConnectionBody)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectionServiceServer).ApproveConnectionRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/connection.ConnectionService/ApproveConnectionRequest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectionServiceServer).ApproveConnectionRequest(ctx, req.(*ConnectionBody))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConnectionService_GetConnectionsUsernamesFor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetConnectionsUsernamesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectionServiceServer).GetConnectionsUsernamesFor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/connection.ConnectionService/GetConnectionsUsernamesFor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectionServiceServer).GetConnectionsUsernamesFor(ctx, req.(*GetConnectionsUsernamesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConnectionService_GetRequestsUsernamesFor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetConnectionsUsernamesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectionServiceServer).GetRequestsUsernamesFor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/connection.ConnectionService/GetRequestsUsernamesFor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectionServiceServer).GetRequestsUsernamesFor(ctx, req.(*GetConnectionsUsernamesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConnectionService_InsertUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectionServiceServer).InsertUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/connection.ConnectionService/InsertUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectionServiceServer).InsertUser(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConnectionService_GetSuggestionIdsFor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSuggestionIdsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectionServiceServer).GetSuggestionIdsFor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/connection.ConnectionService/GetSuggestionIdsFor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectionServiceServer).GetSuggestionIdsFor(ctx, req.(*GetSuggestionIdsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConnectionService_BlockConnection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConnectionBody)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectionServiceServer).BlockConnection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/connection.ConnectionService/BlockConnection",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectionServiceServer).BlockConnection(ctx, req.(*ConnectionBody))
	}
	return interceptor(ctx, in, info, handler)
}

// ConnectionService_ServiceDesc is the grpc.ServiceDesc for ConnectionService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ConnectionService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "connection.ConnectionService",
	HandlerType: (*ConnectionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MakeConnectionWithPublicProfile",
			Handler:    _ConnectionService_MakeConnectionWithPublicProfile_Handler,
		},
		{
			MethodName: "MakeConnectionRequest",
			Handler:    _ConnectionService_MakeConnectionRequest_Handler,
		},
		{
			MethodName: "ApproveConnectionRequest",
			Handler:    _ConnectionService_ApproveConnectionRequest_Handler,
		},
		{
			MethodName: "GetConnectionsUsernamesFor",
			Handler:    _ConnectionService_GetConnectionsUsernamesFor_Handler,
		},
		{
			MethodName: "GetRequestsUsernamesFor",
			Handler:    _ConnectionService_GetRequestsUsernamesFor_Handler,
		},
		{
			MethodName: "InsertUser",
			Handler:    _ConnectionService_InsertUser_Handler,
		},
		{
			MethodName: "GetSuggestionIdsFor",
			Handler:    _ConnectionService_GetSuggestionIdsFor_Handler,
		},
		{
			MethodName: "BlockConnection",
			Handler:    _ConnectionService_BlockConnection_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "connection-service.proto",
}
