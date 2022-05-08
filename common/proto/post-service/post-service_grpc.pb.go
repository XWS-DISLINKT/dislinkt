// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.1
// source: post-service.proto

package post

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

// PostServiceClient is the client API for PostService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PostServiceClient interface {
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
	GetAll(ctx context.Context, in *GetAllRequest, opts ...grpc.CallOption) (*GetAllResponse, error)
	Post(ctx context.Context, in *PostRequest, opts ...grpc.CallOption) (*PostResponse, error)
	GetByUser(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetAllResponse, error)
	LikePost(ctx context.Context, in *ReactionRequest, opts ...grpc.CallOption) (*ReactionResponse, error)
	DislikePost(ctx context.Context, in *ReactionRequest, opts ...grpc.CallOption) (*ReactionResponse, error)
	CommentPost(ctx context.Context, in *CommentRequest, opts ...grpc.CallOption) (*CommentResponse, error)
	GetAllReactionsByPost(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*AllReactionsResponse, error)
	GetAllCommentsByPost(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*AllCommentsResponse, error)
	GetFeed(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetAllResponse, error)
}

type postServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPostServiceClient(cc grpc.ClientConnInterface) PostServiceClient {
	return &postServiceClient{cc}
}

func (c *postServiceClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, "/post.PostService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postServiceClient) GetAll(ctx context.Context, in *GetAllRequest, opts ...grpc.CallOption) (*GetAllResponse, error) {
	out := new(GetAllResponse)
	err := c.cc.Invoke(ctx, "/post.PostService/GetAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postServiceClient) Post(ctx context.Context, in *PostRequest, opts ...grpc.CallOption) (*PostResponse, error) {
	out := new(PostResponse)
	err := c.cc.Invoke(ctx, "/post.PostService/Post", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postServiceClient) GetByUser(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetAllResponse, error) {
	out := new(GetAllResponse)
	err := c.cc.Invoke(ctx, "/post.PostService/GetByUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postServiceClient) LikePost(ctx context.Context, in *ReactionRequest, opts ...grpc.CallOption) (*ReactionResponse, error) {
	out := new(ReactionResponse)
	err := c.cc.Invoke(ctx, "/post.PostService/LikePost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postServiceClient) DislikePost(ctx context.Context, in *ReactionRequest, opts ...grpc.CallOption) (*ReactionResponse, error) {
	out := new(ReactionResponse)
	err := c.cc.Invoke(ctx, "/post.PostService/DislikePost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postServiceClient) CommentPost(ctx context.Context, in *CommentRequest, opts ...grpc.CallOption) (*CommentResponse, error) {
	out := new(CommentResponse)
	err := c.cc.Invoke(ctx, "/post.PostService/CommentPost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postServiceClient) GetAllReactionsByPost(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*AllReactionsResponse, error) {
	out := new(AllReactionsResponse)
	err := c.cc.Invoke(ctx, "/post.PostService/GetAllReactionsByPost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postServiceClient) GetAllCommentsByPost(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*AllCommentsResponse, error) {
	out := new(AllCommentsResponse)
	err := c.cc.Invoke(ctx, "/post.PostService/GetAllCommentsByPost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postServiceClient) GetFeed(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetAllResponse, error) {
	out := new(GetAllResponse)
	err := c.cc.Invoke(ctx, "/post.PostService/GetFeed", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PostServiceServer is the server API for PostService service.
// All implementations must embed UnimplementedPostServiceServer
// for forward compatibility
type PostServiceServer interface {
	Get(context.Context, *GetRequest) (*GetResponse, error)
	GetAll(context.Context, *GetAllRequest) (*GetAllResponse, error)
	Post(context.Context, *PostRequest) (*PostResponse, error)
	GetByUser(context.Context, *GetRequest) (*GetAllResponse, error)
	LikePost(context.Context, *ReactionRequest) (*ReactionResponse, error)
	DislikePost(context.Context, *ReactionRequest) (*ReactionResponse, error)
	CommentPost(context.Context, *CommentRequest) (*CommentResponse, error)
	GetAllReactionsByPost(context.Context, *GetRequest) (*AllReactionsResponse, error)
	GetAllCommentsByPost(context.Context, *GetRequest) (*AllCommentsResponse, error)
	GetFeed(context.Context, *GetRequest) (*GetAllResponse, error)
	mustEmbedUnimplementedPostServiceServer()
}

// UnimplementedPostServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPostServiceServer struct {
}

func (UnimplementedPostServiceServer) Get(context.Context, *GetRequest) (*GetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedPostServiceServer) GetAll(context.Context, *GetAllRequest) (*GetAllResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
func (UnimplementedPostServiceServer) Post(context.Context, *PostRequest) (*PostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Post not implemented")
}
func (UnimplementedPostServiceServer) GetByUser(context.Context, *GetRequest) (*GetAllResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByUser not implemented")
}
func (UnimplementedPostServiceServer) LikePost(context.Context, *ReactionRequest) (*ReactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LikePost not implemented")
}
func (UnimplementedPostServiceServer) DislikePost(context.Context, *ReactionRequest) (*ReactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DislikePost not implemented")
}
func (UnimplementedPostServiceServer) CommentPost(context.Context, *CommentRequest) (*CommentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CommentPost not implemented")
}
func (UnimplementedPostServiceServer) GetAllReactionsByPost(context.Context, *GetRequest) (*AllReactionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllReactionsByPost not implemented")
}
func (UnimplementedPostServiceServer) GetAllCommentsByPost(context.Context, *GetRequest) (*AllCommentsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllCommentsByPost not implemented")
}
func (UnimplementedPostServiceServer) GetFeed(context.Context, *GetRequest) (*GetAllResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFeed not implemented")
}
func (UnimplementedPostServiceServer) mustEmbedUnimplementedPostServiceServer() {}

// UnsafePostServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PostServiceServer will
// result in compilation errors.
type UnsafePostServiceServer interface {
	mustEmbedUnimplementedPostServiceServer()
}

func RegisterPostServiceServer(s grpc.ServiceRegistrar, srv PostServiceServer) {
	s.RegisterService(&PostService_ServiceDesc, srv)
}

func _PostService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/post.PostService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServiceServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostService_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServiceServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/post.PostService/GetAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServiceServer).GetAll(ctx, req.(*GetAllRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostService_Post_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServiceServer).Post(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/post.PostService/Post",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServiceServer).Post(ctx, req.(*PostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostService_GetByUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServiceServer).GetByUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/post.PostService/GetByUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServiceServer).GetByUser(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostService_LikePost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServiceServer).LikePost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/post.PostService/LikePost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServiceServer).LikePost(ctx, req.(*ReactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostService_DislikePost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServiceServer).DislikePost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/post.PostService/DislikePost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServiceServer).DislikePost(ctx, req.(*ReactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostService_CommentPost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServiceServer).CommentPost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/post.PostService/CommentPost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServiceServer).CommentPost(ctx, req.(*CommentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostService_GetAllReactionsByPost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServiceServer).GetAllReactionsByPost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/post.PostService/GetAllReactionsByPost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServiceServer).GetAllReactionsByPost(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostService_GetAllCommentsByPost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServiceServer).GetAllCommentsByPost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/post.PostService/GetAllCommentsByPost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServiceServer).GetAllCommentsByPost(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostService_GetFeed_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServiceServer).GetFeed(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/post.PostService/GetFeed",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServiceServer).GetFeed(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PostService_ServiceDesc is the grpc.ServiceDesc for PostService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PostService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "post.PostService",
	HandlerType: (*PostServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _PostService_Get_Handler,
		},
		{
			MethodName: "GetAll",
			Handler:    _PostService_GetAll_Handler,
		},
		{
			MethodName: "Post",
			Handler:    _PostService_Post_Handler,
		},
		{
			MethodName: "GetByUser",
			Handler:    _PostService_GetByUser_Handler,
		},
		{
			MethodName: "LikePost",
			Handler:    _PostService_LikePost_Handler,
		},
		{
			MethodName: "DislikePost",
			Handler:    _PostService_DislikePost_Handler,
		},
		{
			MethodName: "CommentPost",
			Handler:    _PostService_CommentPost_Handler,
		},
		{
			MethodName: "GetAllReactionsByPost",
			Handler:    _PostService_GetAllReactionsByPost_Handler,
		},
		{
			MethodName: "GetAllCommentsByPost",
			Handler:    _PostService_GetAllCommentsByPost_Handler,
		},
		{
			MethodName: "GetFeed",
			Handler:    _PostService_GetFeed_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "post-service.proto",
}
