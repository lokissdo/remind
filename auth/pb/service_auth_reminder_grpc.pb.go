// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.26.1
// source: service_auth_reminder.proto

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

// AuthRemindClient is the client API for AuthRemind service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthRemindClient interface {
	CreateToken(ctx context.Context, in *CreateTokenRequest, opts ...grpc.CallOption) (*CreateTokenResponse, error)
	VerifyToken(ctx context.Context, in *VerifyTokenRequest, opts ...grpc.CallOption) (*VerifyTokenResponse, error)
}

type authRemindClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthRemindClient(cc grpc.ClientConnInterface) AuthRemindClient {
	return &authRemindClient{cc}
}

func (c *authRemindClient) CreateToken(ctx context.Context, in *CreateTokenRequest, opts ...grpc.CallOption) (*CreateTokenResponse, error) {
	out := new(CreateTokenResponse)
	err := c.cc.Invoke(ctx, "/pb.AuthRemind/CreateToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authRemindClient) VerifyToken(ctx context.Context, in *VerifyTokenRequest, opts ...grpc.CallOption) (*VerifyTokenResponse, error) {
	out := new(VerifyTokenResponse)
	err := c.cc.Invoke(ctx, "/pb.AuthRemind/VerifyToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthRemindServer is the server API for AuthRemind service.
// All implementations must embed UnimplementedAuthRemindServer
// for forward compatibility
type AuthRemindServer interface {
	CreateToken(context.Context, *CreateTokenRequest) (*CreateTokenResponse, error)
	VerifyToken(context.Context, *VerifyTokenRequest) (*VerifyTokenResponse, error)
	mustEmbedUnimplementedAuthRemindServer()
}

// UnimplementedAuthRemindServer must be embedded to have forward compatible implementations.
type UnimplementedAuthRemindServer struct {
}

func (UnimplementedAuthRemindServer) CreateToken(context.Context, *CreateTokenRequest) (*CreateTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateToken not implemented")
}
func (UnimplementedAuthRemindServer) VerifyToken(context.Context, *VerifyTokenRequest) (*VerifyTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyToken not implemented")
}
func (UnimplementedAuthRemindServer) mustEmbedUnimplementedAuthRemindServer() {}

// UnsafeAuthRemindServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthRemindServer will
// result in compilation errors.
type UnsafeAuthRemindServer interface {
	mustEmbedUnimplementedAuthRemindServer()
}

func RegisterAuthRemindServer(s grpc.ServiceRegistrar, srv AuthRemindServer) {
	s.RegisterService(&AuthRemind_ServiceDesc, srv)
}

func _AuthRemind_CreateToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthRemindServer).CreateToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.AuthRemind/CreateToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthRemindServer).CreateToken(ctx, req.(*CreateTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthRemind_VerifyToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifyTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthRemindServer).VerifyToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.AuthRemind/VerifyToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthRemindServer).VerifyToken(ctx, req.(*VerifyTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AuthRemind_ServiceDesc is the grpc.ServiceDesc for AuthRemind service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AuthRemind_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.AuthRemind",
	HandlerType: (*AuthRemindServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateToken",
			Handler:    _AuthRemind_CreateToken_Handler,
		},
		{
			MethodName: "VerifyToken",
			Handler:    _AuthRemind_VerifyToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service_auth_reminder.proto",
}
