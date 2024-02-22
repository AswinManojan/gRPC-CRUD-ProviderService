// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: pkg/PB/serviceprovider.proto

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

// ProviderServicesClient is the client API for ProviderServices service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProviderServicesClient interface {
	ProviderSignUp(ctx context.Context, in *ProviderDataRequest, opts ...grpc.CallOption) (*Result, error)
	ProviderLogin(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*Result, error)
	EditProvider(ctx context.Context, in *ProviderDataRequest, opts ...grpc.CallOption) (*Result, error)
	DeleteProviderById(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*Result, error)
	FindProviderById(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*Result, error)
	CreateProvider(ctx context.Context, in *ProviderDataRequest, opts ...grpc.CallOption) (*Result, error)
}

type providerServicesClient struct {
	cc grpc.ClientConnInterface
}

func NewProviderServicesClient(cc grpc.ClientConnInterface) ProviderServicesClient {
	return &providerServicesClient{cc}
}

func (c *providerServicesClient) ProviderSignUp(ctx context.Context, in *ProviderDataRequest, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/pb.ProviderServices/ProviderSignUp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *providerServicesClient) ProviderLogin(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/pb.ProviderServices/ProviderLogin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *providerServicesClient) EditProvider(ctx context.Context, in *ProviderDataRequest, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/pb.ProviderServices/EditProvider", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *providerServicesClient) DeleteProviderById(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/pb.ProviderServices/DeleteProviderById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *providerServicesClient) FindProviderById(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/pb.ProviderServices/FindProviderById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *providerServicesClient) CreateProvider(ctx context.Context, in *ProviderDataRequest, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/pb.ProviderServices/CreateProvider", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProviderServicesServer is the server API for ProviderServices service.
// All implementations must embed UnimplementedProviderServicesServer
// for forward compatibility
type ProviderServicesServer interface {
	ProviderSignUp(context.Context, *ProviderDataRequest) (*Result, error)
	ProviderLogin(context.Context, *LoginRequest) (*Result, error)
	EditProvider(context.Context, *ProviderDataRequest) (*Result, error)
	DeleteProviderById(context.Context, *IdRequest) (*Result, error)
	FindProviderById(context.Context, *IdRequest) (*Result, error)
	CreateProvider(context.Context, *ProviderDataRequest) (*Result, error)
	mustEmbedUnimplementedProviderServicesServer()
}

// UnimplementedProviderServicesServer must be embedded to have forward compatible implementations.
type UnimplementedProviderServicesServer struct {
}

func (UnimplementedProviderServicesServer) ProviderSignUp(context.Context, *ProviderDataRequest) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProviderSignUp not implemented")
}
func (UnimplementedProviderServicesServer) ProviderLogin(context.Context, *LoginRequest) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProviderLogin not implemented")
}
func (UnimplementedProviderServicesServer) EditProvider(context.Context, *ProviderDataRequest) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EditProvider not implemented")
}
func (UnimplementedProviderServicesServer) DeleteProviderById(context.Context, *IdRequest) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteProviderById not implemented")
}
func (UnimplementedProviderServicesServer) FindProviderById(context.Context, *IdRequest) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindProviderById not implemented")
}
func (UnimplementedProviderServicesServer) CreateProvider(context.Context, *ProviderDataRequest) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateProvider not implemented")
}
func (UnimplementedProviderServicesServer) mustEmbedUnimplementedProviderServicesServer() {}

// UnsafeProviderServicesServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProviderServicesServer will
// result in compilation errors.
type UnsafeProviderServicesServer interface {
	mustEmbedUnimplementedProviderServicesServer()
}

func RegisterProviderServicesServer(s grpc.ServiceRegistrar, srv ProviderServicesServer) {
	s.RegisterService(&ProviderServices_ServiceDesc, srv)
}

func _ProviderServices_ProviderSignUp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProviderDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProviderServicesServer).ProviderSignUp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ProviderServices/ProviderSignUp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProviderServicesServer).ProviderSignUp(ctx, req.(*ProviderDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProviderServices_ProviderLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProviderServicesServer).ProviderLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ProviderServices/ProviderLogin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProviderServicesServer).ProviderLogin(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProviderServices_EditProvider_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProviderDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProviderServicesServer).EditProvider(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ProviderServices/EditProvider",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProviderServicesServer).EditProvider(ctx, req.(*ProviderDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProviderServices_DeleteProviderById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProviderServicesServer).DeleteProviderById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ProviderServices/DeleteProviderById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProviderServicesServer).DeleteProviderById(ctx, req.(*IdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProviderServices_FindProviderById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProviderServicesServer).FindProviderById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ProviderServices/FindProviderById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProviderServicesServer).FindProviderById(ctx, req.(*IdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProviderServices_CreateProvider_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProviderDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProviderServicesServer).CreateProvider(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ProviderServices/CreateProvider",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProviderServicesServer).CreateProvider(ctx, req.(*ProviderDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ProviderServices_ServiceDesc is the grpc.ServiceDesc for ProviderServices service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProviderServices_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.ProviderServices",
	HandlerType: (*ProviderServicesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ProviderSignUp",
			Handler:    _ProviderServices_ProviderSignUp_Handler,
		},
		{
			MethodName: "ProviderLogin",
			Handler:    _ProviderServices_ProviderLogin_Handler,
		},
		{
			MethodName: "EditProvider",
			Handler:    _ProviderServices_EditProvider_Handler,
		},
		{
			MethodName: "DeleteProviderById",
			Handler:    _ProviderServices_DeleteProviderById_Handler,
		},
		{
			MethodName: "FindProviderById",
			Handler:    _ProviderServices_FindProviderById_Handler,
		},
		{
			MethodName: "CreateProvider",
			Handler:    _ProviderServices_CreateProvider_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/PB/serviceprovider.proto",
}
