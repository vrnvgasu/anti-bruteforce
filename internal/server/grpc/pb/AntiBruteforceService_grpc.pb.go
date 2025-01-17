// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: api/AntiBruteforceService.proto

package pb

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
	AntiBruteforce_Check_FullMethodName          = "/event.AntiBruteforce/Check"
	AntiBruteforce_Clear_FullMethodName          = "/event.AntiBruteforce/Clear"
	AntiBruteforce_AddToList_FullMethodName      = "/event.AntiBruteforce/AddToList"
	AntiBruteforce_RemoveFromList_FullMethodName = "/event.AntiBruteforce/RemoveFromList"
)

// AntiBruteforceClient is the client API for AntiBruteforce service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AntiBruteforceClient interface {
	Check(ctx context.Context, in *CheckRequest, opts ...grpc.CallOption) (*CheckResponse, error)
	Clear(ctx context.Context, in *ClearRequest, opts ...grpc.CallOption) (*OK, error)
	AddToList(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*OK, error)
	RemoveFromList(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*OK, error)
}

type antiBruteforceClient struct {
	cc grpc.ClientConnInterface
}

func NewAntiBruteforceClient(cc grpc.ClientConnInterface) AntiBruteforceClient {
	return &antiBruteforceClient{cc}
}

func (c *antiBruteforceClient) Check(ctx context.Context, in *CheckRequest, opts ...grpc.CallOption) (*CheckResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CheckResponse)
	err := c.cc.Invoke(ctx, AntiBruteforce_Check_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *antiBruteforceClient) Clear(ctx context.Context, in *ClearRequest, opts ...grpc.CallOption) (*OK, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(OK)
	err := c.cc.Invoke(ctx, AntiBruteforce_Clear_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *antiBruteforceClient) AddToList(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*OK, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(OK)
	err := c.cc.Invoke(ctx, AntiBruteforce_AddToList_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *antiBruteforceClient) RemoveFromList(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*OK, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(OK)
	err := c.cc.Invoke(ctx, AntiBruteforce_RemoveFromList_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AntiBruteforceServer is the server API for AntiBruteforce service.
// All implementations must embed UnimplementedAntiBruteforceServer
// for forward compatibility.
type AntiBruteforceServer interface {
	Check(context.Context, *CheckRequest) (*CheckResponse, error)
	Clear(context.Context, *ClearRequest) (*OK, error)
	AddToList(context.Context, *ListRequest) (*OK, error)
	RemoveFromList(context.Context, *ListRequest) (*OK, error)
	mustEmbedUnimplementedAntiBruteforceServer()
}

// UnimplementedAntiBruteforceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedAntiBruteforceServer struct{}

func (UnimplementedAntiBruteforceServer) Check(context.Context, *CheckRequest) (*CheckResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Check not implemented")
}
func (UnimplementedAntiBruteforceServer) Clear(context.Context, *ClearRequest) (*OK, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Clear not implemented")
}
func (UnimplementedAntiBruteforceServer) AddToList(context.Context, *ListRequest) (*OK, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddToList not implemented")
}
func (UnimplementedAntiBruteforceServer) RemoveFromList(context.Context, *ListRequest) (*OK, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveFromList not implemented")
}
func (UnimplementedAntiBruteforceServer) mustEmbedUnimplementedAntiBruteforceServer() {}
func (UnimplementedAntiBruteforceServer) testEmbeddedByValue()                        {}

// UnsafeAntiBruteforceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AntiBruteforceServer will
// result in compilation errors.
type UnsafeAntiBruteforceServer interface {
	mustEmbedUnimplementedAntiBruteforceServer()
}

func RegisterAntiBruteforceServer(s grpc.ServiceRegistrar, srv AntiBruteforceServer) {
	// If the following call pancis, it indicates UnimplementedAntiBruteforceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&AntiBruteforce_ServiceDesc, srv)
}

func _AntiBruteforce_Check_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AntiBruteforceServer).Check(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AntiBruteforce_Check_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AntiBruteforceServer).Check(ctx, req.(*CheckRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AntiBruteforce_Clear_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClearRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AntiBruteforceServer).Clear(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AntiBruteforce_Clear_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AntiBruteforceServer).Clear(ctx, req.(*ClearRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AntiBruteforce_AddToList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AntiBruteforceServer).AddToList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AntiBruteforce_AddToList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AntiBruteforceServer).AddToList(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AntiBruteforce_RemoveFromList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AntiBruteforceServer).RemoveFromList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AntiBruteforce_RemoveFromList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AntiBruteforceServer).RemoveFromList(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AntiBruteforce_ServiceDesc is the grpc.ServiceDesc for AntiBruteforce service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AntiBruteforce_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "event.AntiBruteforce",
	HandlerType: (*AntiBruteforceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Check",
			Handler:    _AntiBruteforce_Check_Handler,
		},
		{
			MethodName: "Clear",
			Handler:    _AntiBruteforce_Clear_Handler,
		},
		{
			MethodName: "AddToList",
			Handler:    _AntiBruteforce_AddToList_Handler,
		},
		{
			MethodName: "RemoveFromList",
			Handler:    _AntiBruteforce_RemoveFromList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/AntiBruteforceService.proto",
}
