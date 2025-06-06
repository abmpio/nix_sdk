// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: nix.proto

package proto

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
	Nix_HealthCheck_FullMethodName         = "/nix.Nix/HealthCheck"
	Nix_FindOneKVByKey_FullMethodName      = "/nix.Nix/FindOneKVByKey"
	Nix_FindKVListByKeyList_FullMethodName = "/nix.Nix/FindKVListByKeyList"
	Nix_FindKVListByTag_FullMethodName     = "/nix.Nix/FindKVListByTag"
)

// NixClient is the client API for Nix service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NixClient interface {
	HealthCheck(ctx context.Context, in *NixHealthCheckRequest, opts ...grpc.CallOption) (*NixHealthCheckResponse, error)
	FindOneKVByKey(ctx context.Context, in *FindKVOneByKeyRequest, opts ...grpc.CallOption) (*FindKVOneByKeyResponse, error)
	FindKVListByKeyList(ctx context.Context, in *FindKVListByKeyListRequest, opts ...grpc.CallOption) (*FindKVListByKeyListResponse, error)
	FindKVListByTag(ctx context.Context, in *FindKVListByTagRequest, opts ...grpc.CallOption) (*FindKVListByTagResponse, error)
}

type nixClient struct {
	cc grpc.ClientConnInterface
}

func NewNixClient(cc grpc.ClientConnInterface) NixClient {
	return &nixClient{cc}
}

func (c *nixClient) HealthCheck(ctx context.Context, in *NixHealthCheckRequest, opts ...grpc.CallOption) (*NixHealthCheckResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(NixHealthCheckResponse)
	err := c.cc.Invoke(ctx, Nix_HealthCheck_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nixClient) FindOneKVByKey(ctx context.Context, in *FindKVOneByKeyRequest, opts ...grpc.CallOption) (*FindKVOneByKeyResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(FindKVOneByKeyResponse)
	err := c.cc.Invoke(ctx, Nix_FindOneKVByKey_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nixClient) FindKVListByKeyList(ctx context.Context, in *FindKVListByKeyListRequest, opts ...grpc.CallOption) (*FindKVListByKeyListResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(FindKVListByKeyListResponse)
	err := c.cc.Invoke(ctx, Nix_FindKVListByKeyList_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nixClient) FindKVListByTag(ctx context.Context, in *FindKVListByTagRequest, opts ...grpc.CallOption) (*FindKVListByTagResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(FindKVListByTagResponse)
	err := c.cc.Invoke(ctx, Nix_FindKVListByTag_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NixServer is the server API for Nix service.
// All implementations must embed UnimplementedNixServer
// for forward compatibility.
type NixServer interface {
	HealthCheck(context.Context, *NixHealthCheckRequest) (*NixHealthCheckResponse, error)
	FindOneKVByKey(context.Context, *FindKVOneByKeyRequest) (*FindKVOneByKeyResponse, error)
	FindKVListByKeyList(context.Context, *FindKVListByKeyListRequest) (*FindKVListByKeyListResponse, error)
	FindKVListByTag(context.Context, *FindKVListByTagRequest) (*FindKVListByTagResponse, error)
	mustEmbedUnimplementedNixServer()
}

// UnimplementedNixServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedNixServer struct{}

func (UnimplementedNixServer) HealthCheck(context.Context, *NixHealthCheckRequest) (*NixHealthCheckResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HealthCheck not implemented")
}
func (UnimplementedNixServer) FindOneKVByKey(context.Context, *FindKVOneByKeyRequest) (*FindKVOneByKeyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindOneKVByKey not implemented")
}
func (UnimplementedNixServer) FindKVListByKeyList(context.Context, *FindKVListByKeyListRequest) (*FindKVListByKeyListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindKVListByKeyList not implemented")
}
func (UnimplementedNixServer) FindKVListByTag(context.Context, *FindKVListByTagRequest) (*FindKVListByTagResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindKVListByTag not implemented")
}
func (UnimplementedNixServer) mustEmbedUnimplementedNixServer() {}
func (UnimplementedNixServer) testEmbeddedByValue()             {}

// UnsafeNixServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NixServer will
// result in compilation errors.
type UnsafeNixServer interface {
	mustEmbedUnimplementedNixServer()
}

func RegisterNixServer(s grpc.ServiceRegistrar, srv NixServer) {
	// If the following call pancis, it indicates UnimplementedNixServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Nix_ServiceDesc, srv)
}

func _Nix_HealthCheck_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NixHealthCheckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NixServer).HealthCheck(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Nix_HealthCheck_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NixServer).HealthCheck(ctx, req.(*NixHealthCheckRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Nix_FindOneKVByKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindKVOneByKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NixServer).FindOneKVByKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Nix_FindOneKVByKey_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NixServer).FindOneKVByKey(ctx, req.(*FindKVOneByKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Nix_FindKVListByKeyList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindKVListByKeyListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NixServer).FindKVListByKeyList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Nix_FindKVListByKeyList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NixServer).FindKVListByKeyList(ctx, req.(*FindKVListByKeyListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Nix_FindKVListByTag_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindKVListByTagRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NixServer).FindKVListByTag(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Nix_FindKVListByTag_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NixServer).FindKVListByTag(ctx, req.(*FindKVListByTagRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Nix_ServiceDesc is the grpc.ServiceDesc for Nix service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Nix_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "nix.Nix",
	HandlerType: (*NixServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "HealthCheck",
			Handler:    _Nix_HealthCheck_Handler,
		},
		{
			MethodName: "FindOneKVByKey",
			Handler:    _Nix_FindOneKVByKey_Handler,
		},
		{
			MethodName: "FindKVListByKeyList",
			Handler:    _Nix_FindKVListByKeyList_Handler,
		},
		{
			MethodName: "FindKVListByTag",
			Handler:    _Nix_FindKVListByTag_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "nix.proto",
}
