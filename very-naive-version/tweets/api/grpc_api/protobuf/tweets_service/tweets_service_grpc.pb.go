// Code generated by protoc-gen-go-grpc_api. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc_api v1.2.0
// - protoc             v3.6.1
// source: protobuf/tweets_service/tweets_service.proto

package tweets_service

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc_api package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// TweetsServiceClient is the client API for TweetsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TweetsServiceClient interface {
	PostTweet(ctx context.Context, in *PostTweetRequest, opts ...grpc.CallOption) (*PostTweetResponse, error)
	DeleteTweetById(ctx context.Context, in *DeleteTweetByIdRequest, opts ...grpc.CallOption) (*DeleteTweetByIdResponse, error)
	GetTweetById(ctx context.Context, in *GetTweetByIdRequest, opts ...grpc.CallOption) (*GetTweetByIdResponse, error)
	GetAllTweets(ctx context.Context, in *GetAllTweetsRequest, opts ...grpc.CallOption) (*GetAllTweetsResponse, error)
	GetTweetsOfUser(ctx context.Context, in *GetTweetsOfUserRequest, opts ...grpc.CallOption) (*GetTweetsOfUserResponse, error)
}

type tweetsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTweetsServiceClient(cc grpc.ClientConnInterface) TweetsServiceClient {
	return &tweetsServiceClient{cc}
}

func (c *tweetsServiceClient) PostTweet(ctx context.Context, in *PostTweetRequest, opts ...grpc.CallOption) (*PostTweetResponse, error) {
	out := new(PostTweetResponse)
	err := c.cc.Invoke(ctx, "/TweetsService/PostTweet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tweetsServiceClient) DeleteTweetById(ctx context.Context, in *DeleteTweetByIdRequest, opts ...grpc.CallOption) (*DeleteTweetByIdResponse, error) {
	out := new(DeleteTweetByIdResponse)
	err := c.cc.Invoke(ctx, "/TweetsService/DeleteTweetById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tweetsServiceClient) GetTweetById(ctx context.Context, in *GetTweetByIdRequest, opts ...grpc.CallOption) (*GetTweetByIdResponse, error) {
	out := new(GetTweetByIdResponse)
	err := c.cc.Invoke(ctx, "/TweetsService/GetTweetById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tweetsServiceClient) GetAllTweets(ctx context.Context, in *GetAllTweetsRequest, opts ...grpc.CallOption) (*GetAllTweetsResponse, error) {
	out := new(GetAllTweetsResponse)
	err := c.cc.Invoke(ctx, "/TweetsService/GetAllTweets", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tweetsServiceClient) GetTweetsOfUser(ctx context.Context, in *GetTweetsOfUserRequest, opts ...grpc.CallOption) (*GetTweetsOfUserResponse, error) {
	out := new(GetTweetsOfUserResponse)
	err := c.cc.Invoke(ctx, "/TweetsService/GetTweetsOfUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TweetsServiceServer is the server API for TweetsService service.
// All implementations must embed UnimplementedTweetsServiceServer
// for forward compatibility
type TweetsServiceServer interface {
	PostTweet(context.Context, *PostTweetRequest) (*PostTweetResponse, error)
	DeleteTweetById(context.Context, *DeleteTweetByIdRequest) (*DeleteTweetByIdResponse, error)
	GetTweetById(context.Context, *GetTweetByIdRequest) (*GetTweetByIdResponse, error)
	GetAllTweets(context.Context, *GetAllTweetsRequest) (*GetAllTweetsResponse, error)
	GetTweetsOfUser(context.Context, *GetTweetsOfUserRequest) (*GetTweetsOfUserResponse, error)
	mustEmbedUnimplementedTweetsServiceServer()
}

// UnimplementedTweetsServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTweetsServiceServer struct {
}

func (UnimplementedTweetsServiceServer) PostTweet(context.Context, *PostTweetRequest) (*PostTweetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PostTweet not implemented")
}
func (UnimplementedTweetsServiceServer) DeleteTweetById(context.Context, *DeleteTweetByIdRequest) (*DeleteTweetByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTweetById not implemented")
}
func (UnimplementedTweetsServiceServer) GetTweetById(context.Context, *GetTweetByIdRequest) (*GetTweetByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTweetById not implemented")
}
func (UnimplementedTweetsServiceServer) GetAllTweets(context.Context, *GetAllTweetsRequest) (*GetAllTweetsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllTweets not implemented")
}
func (UnimplementedTweetsServiceServer) GetTweetsOfUser(context.Context, *GetTweetsOfUserRequest) (*GetTweetsOfUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTweetsOfUser not implemented")
}
func (UnimplementedTweetsServiceServer) mustEmbedUnimplementedTweetsServiceServer() {}

// UnsafeTweetsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TweetsServiceServer will
// result in compilation errors.
type UnsafeTweetsServiceServer interface {
	mustEmbedUnimplementedTweetsServiceServer()
}

func RegisterTweetsServiceServer(s grpc.ServiceRegistrar, srv TweetsServiceServer) {
	s.RegisterService(&TweetsService_ServiceDesc, srv)
}

func _TweetsService_PostTweet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PostTweetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TweetsServiceServer).PostTweet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/TweetsService/PostTweet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TweetsServiceServer).PostTweet(ctx, req.(*PostTweetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TweetsService_DeleteTweetById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTweetByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TweetsServiceServer).DeleteTweetById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/TweetsService/DeleteTweetById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TweetsServiceServer).DeleteTweetById(ctx, req.(*DeleteTweetByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TweetsService_GetTweetById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTweetByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TweetsServiceServer).GetTweetById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/TweetsService/GetTweetById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TweetsServiceServer).GetTweetById(ctx, req.(*GetTweetByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TweetsService_GetAllTweets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllTweetsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TweetsServiceServer).GetAllTweets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/TweetsService/GetAllTweets",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TweetsServiceServer).GetAllTweets(ctx, req.(*GetAllTweetsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TweetsService_GetTweetsOfUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTweetsOfUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TweetsServiceServer).GetTweetsOfUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/TweetsService/GetTweetsOfUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TweetsServiceServer).GetTweetsOfUser(ctx, req.(*GetTweetsOfUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TweetsService_ServiceDesc is the grpc.ServiceDesc for TweetsService service.
// It's only intended for direct use with grpc_api.RegisterService,
// and not to be introspected or modified (even as a copy)
var TweetsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "TweetsService",
	HandlerType: (*TweetsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PostTweet",
			Handler:    _TweetsService_PostTweet_Handler,
		},
		{
			MethodName: "DeleteTweetById",
			Handler:    _TweetsService_DeleteTweetById_Handler,
		},
		{
			MethodName: "GetTweetById",
			Handler:    _TweetsService_GetTweetById_Handler,
		},
		{
			MethodName: "GetAllTweets",
			Handler:    _TweetsService_GetAllTweets_Handler,
		},
		{
			MethodName: "GetTweetsOfUser",
			Handler:    _TweetsService_GetTweetsOfUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protobuf/tweets_service/tweets_service.proto",
}
