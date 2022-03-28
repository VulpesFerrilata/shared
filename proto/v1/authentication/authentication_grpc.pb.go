// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.1
// source: proto/v1/authentication/authentication.proto

package authentication

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// AuthenticationClient is the client API for Authentication service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthenticationClient interface {
	CreateUserCredential(ctx context.Context, in *CreateUserCredentialRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	RemoveUserCredential(ctx context.Context, in *RemoveUserCredentialRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	RevokeToken(ctx context.Context, in *RevokeTokenRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetTokenByClaim(ctx context.Context, in *GetTokenByClaimRequest, opts ...grpc.CallOption) (*TokenResponse, error)
	GetClaimByAccessToken(ctx context.Context, in *GetClaimByAccessTokenRequest, opts ...grpc.CallOption) (*ClaimResponse, error)
	GetAccessTokenByRefreshToken(ctx context.Context, in *GetAccessTokenByRefreshTokenRequest, opts ...grpc.CallOption) (*TokenResponse, error)
}

type authenticationClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthenticationClient(cc grpc.ClientConnInterface) AuthenticationClient {
	return &authenticationClient{cc}
}

func (c *authenticationClient) CreateUserCredential(ctx context.Context, in *CreateUserCredentialRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/authentication.Authentication/CreateUserCredential", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authenticationClient) RemoveUserCredential(ctx context.Context, in *RemoveUserCredentialRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/authentication.Authentication/RemoveUserCredential", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authenticationClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/authentication.Authentication/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authenticationClient) RevokeToken(ctx context.Context, in *RevokeTokenRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/authentication.Authentication/RevokeToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authenticationClient) GetTokenByClaim(ctx context.Context, in *GetTokenByClaimRequest, opts ...grpc.CallOption) (*TokenResponse, error) {
	out := new(TokenResponse)
	err := c.cc.Invoke(ctx, "/authentication.Authentication/GetTokenByClaim", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authenticationClient) GetClaimByAccessToken(ctx context.Context, in *GetClaimByAccessTokenRequest, opts ...grpc.CallOption) (*ClaimResponse, error) {
	out := new(ClaimResponse)
	err := c.cc.Invoke(ctx, "/authentication.Authentication/GetClaimByAccessToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authenticationClient) GetAccessTokenByRefreshToken(ctx context.Context, in *GetAccessTokenByRefreshTokenRequest, opts ...grpc.CallOption) (*TokenResponse, error) {
	out := new(TokenResponse)
	err := c.cc.Invoke(ctx, "/authentication.Authentication/GetAccessTokenByRefreshToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthenticationServer is the server API for Authentication service.
// All implementations must embed UnimplementedAuthenticationServer
// for forward compatibility
type AuthenticationServer interface {
	CreateUserCredential(context.Context, *CreateUserCredentialRequest) (*emptypb.Empty, error)
	RemoveUserCredential(context.Context, *RemoveUserCredentialRequest) (*emptypb.Empty, error)
	Login(context.Context, *LoginRequest) (*emptypb.Empty, error)
	RevokeToken(context.Context, *RevokeTokenRequest) (*emptypb.Empty, error)
	GetTokenByClaim(context.Context, *GetTokenByClaimRequest) (*TokenResponse, error)
	GetClaimByAccessToken(context.Context, *GetClaimByAccessTokenRequest) (*ClaimResponse, error)
	GetAccessTokenByRefreshToken(context.Context, *GetAccessTokenByRefreshTokenRequest) (*TokenResponse, error)
	mustEmbedUnimplementedAuthenticationServer()
}

// UnimplementedAuthenticationServer must be embedded to have forward compatible implementations.
type UnimplementedAuthenticationServer struct {
}

func (UnimplementedAuthenticationServer) CreateUserCredential(context.Context, *CreateUserCredentialRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUserCredential not implemented")
}
func (UnimplementedAuthenticationServer) RemoveUserCredential(context.Context, *RemoveUserCredentialRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveUserCredential not implemented")
}
func (UnimplementedAuthenticationServer) Login(context.Context, *LoginRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedAuthenticationServer) RevokeToken(context.Context, *RevokeTokenRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RevokeToken not implemented")
}
func (UnimplementedAuthenticationServer) GetTokenByClaim(context.Context, *GetTokenByClaimRequest) (*TokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTokenByClaim not implemented")
}
func (UnimplementedAuthenticationServer) GetClaimByAccessToken(context.Context, *GetClaimByAccessTokenRequest) (*ClaimResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetClaimByAccessToken not implemented")
}
func (UnimplementedAuthenticationServer) GetAccessTokenByRefreshToken(context.Context, *GetAccessTokenByRefreshTokenRequest) (*TokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAccessTokenByRefreshToken not implemented")
}
func (UnimplementedAuthenticationServer) mustEmbedUnimplementedAuthenticationServer() {}

// UnsafeAuthenticationServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthenticationServer will
// result in compilation errors.
type UnsafeAuthenticationServer interface {
	mustEmbedUnimplementedAuthenticationServer()
}

func RegisterAuthenticationServer(s grpc.ServiceRegistrar, srv AuthenticationServer) {
	s.RegisterService(&Authentication_ServiceDesc, srv)
}

func _Authentication_CreateUserCredential_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserCredentialRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticationServer).CreateUserCredential(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/authentication.Authentication/CreateUserCredential",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticationServer).CreateUserCredential(ctx, req.(*CreateUserCredentialRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Authentication_RemoveUserCredential_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveUserCredentialRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticationServer).RemoveUserCredential(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/authentication.Authentication/RemoveUserCredential",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticationServer).RemoveUserCredential(ctx, req.(*RemoveUserCredentialRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Authentication_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticationServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/authentication.Authentication/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticationServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Authentication_RevokeToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RevokeTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticationServer).RevokeToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/authentication.Authentication/RevokeToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticationServer).RevokeToken(ctx, req.(*RevokeTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Authentication_GetTokenByClaim_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTokenByClaimRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticationServer).GetTokenByClaim(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/authentication.Authentication/GetTokenByClaim",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticationServer).GetTokenByClaim(ctx, req.(*GetTokenByClaimRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Authentication_GetClaimByAccessToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetClaimByAccessTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticationServer).GetClaimByAccessToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/authentication.Authentication/GetClaimByAccessToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticationServer).GetClaimByAccessToken(ctx, req.(*GetClaimByAccessTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Authentication_GetAccessTokenByRefreshToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAccessTokenByRefreshTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticationServer).GetAccessTokenByRefreshToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/authentication.Authentication/GetAccessTokenByRefreshToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticationServer).GetAccessTokenByRefreshToken(ctx, req.(*GetAccessTokenByRefreshTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Authentication_ServiceDesc is the grpc.ServiceDesc for Authentication service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Authentication_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "authentication.Authentication",
	HandlerType: (*AuthenticationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateUserCredential",
			Handler:    _Authentication_CreateUserCredential_Handler,
		},
		{
			MethodName: "RemoveUserCredential",
			Handler:    _Authentication_RemoveUserCredential_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _Authentication_Login_Handler,
		},
		{
			MethodName: "RevokeToken",
			Handler:    _Authentication_RevokeToken_Handler,
		},
		{
			MethodName: "GetTokenByClaim",
			Handler:    _Authentication_GetTokenByClaim_Handler,
		},
		{
			MethodName: "GetClaimByAccessToken",
			Handler:    _Authentication_GetClaimByAccessToken_Handler,
		},
		{
			MethodName: "GetAccessTokenByRefreshToken",
			Handler:    _Authentication_GetAccessTokenByRefreshToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/v1/authentication/authentication.proto",
}
