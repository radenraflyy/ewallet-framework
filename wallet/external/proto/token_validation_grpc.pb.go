// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v6.31.1
// source: cmd/proto/token_validation.proto

package tokenvalidation

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	TokenValidationService_ValidateToken_FullMethodName = "/tokenvalidation.TokenValidationService/ValidateToken"
)

// TokenValidationServiceClient is the client API for TokenValidationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Defien the request message for token validation
type TokenValidationServiceClient interface {
	// the method to validate a token
	ValidateToken(ctx context.Context, in *TokenRequest, opts ...grpc.CallOption) (*TokenResponse, error)
}

type tokenValidationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTokenValidationServiceClient(cc grpc.ClientConnInterface) TokenValidationServiceClient {
	return &tokenValidationServiceClient{cc}
}

func (c *tokenValidationServiceClient) ValidateToken(ctx context.Context, in *TokenRequest, opts ...grpc.CallOption) (*TokenResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TokenResponse)
	err := c.cc.Invoke(ctx, TokenValidationService_ValidateToken_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TokenValidationServiceServer is the server API for TokenValidationService service.
// All implementations must embed UnimplementedTokenValidationServiceServer
// for forward compatibility.
//
// Defien the request message for token validation
type TokenValidationServiceServer interface {
	// the method to validate a token
	ValidateToken(context.Context, *TokenRequest) (*TokenResponse, error)
	mustEmbedUnimplementedTokenValidationServiceServer()
}

// UnimplementedTokenValidationServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedTokenValidationServiceServer struct{}

func (UnimplementedTokenValidationServiceServer) ValidateToken(context.Context, *TokenRequest) (*TokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidateToken not implemented")
}
func (UnimplementedTokenValidationServiceServer) mustEmbedUnimplementedTokenValidationServiceServer() {
}
func (UnimplementedTokenValidationServiceServer) testEmbeddedByValue() {}

// UnsafeTokenValidationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TokenValidationServiceServer will
// result in compilation errors.
type UnsafeTokenValidationServiceServer interface {
	mustEmbedUnimplementedTokenValidationServiceServer()
}

func RegisterTokenValidationServiceServer(s grpc.ServiceRegistrar, srv TokenValidationServiceServer) {
	// If the following call pancis, it indicates UnimplementedTokenValidationServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&TokenValidationService_ServiceDesc, srv)
}

func _TokenValidationService_ValidateToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TokenValidationServiceServer).ValidateToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TokenValidationService_ValidateToken_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TokenValidationServiceServer).ValidateToken(ctx, req.(*TokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TokenValidationService_ServiceDesc is the grpc.ServiceDesc for TokenValidationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TokenValidationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "tokenvalidation.TokenValidationService",
	HandlerType: (*TokenValidationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ValidateToken",
			Handler:    _TokenValidationService_ValidateToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cmd/proto/token_validation.proto",
}
