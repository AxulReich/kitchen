// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.9
// source: kitchen_api.proto

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

// KitchenClient is the client API for Kitchen service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type KitchenClient interface {
	GetCookOrders(ctx context.Context, in *GetCookOrdersRequest, opts ...grpc.CallOption) (*GetCookOrdersResponse, error)
	CookingStart(ctx context.Context, in *CookingStartRequest, opts ...grpc.CallOption) (*CookingStartResponse, error)
	CookingEnd(ctx context.Context, in *CookingEndRequest, opts ...grpc.CallOption) (*CookingEndResponse, error)
}

type kitchenClient struct {
	cc grpc.ClientConnInterface
}

func NewKitchenClient(cc grpc.ClientConnInterface) KitchenClient {
	return &kitchenClient{cc}
}

func (c *kitchenClient) GetCookOrders(ctx context.Context, in *GetCookOrdersRequest, opts ...grpc.CallOption) (*GetCookOrdersResponse, error) {
	out := new(GetCookOrdersResponse)
	err := c.cc.Invoke(ctx, "/kitchen_api.v1.Kitchen/GetCookOrders", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kitchenClient) CookingStart(ctx context.Context, in *CookingStartRequest, opts ...grpc.CallOption) (*CookingStartResponse, error) {
	out := new(CookingStartResponse)
	err := c.cc.Invoke(ctx, "/kitchen_api.v1.Kitchen/CookingStart", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kitchenClient) CookingEnd(ctx context.Context, in *CookingEndRequest, opts ...grpc.CallOption) (*CookingEndResponse, error) {
	out := new(CookingEndResponse)
	err := c.cc.Invoke(ctx, "/kitchen_api.v1.Kitchen/CookingEnd", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// KitchenServer is the server API for Kitchen service.
// All implementations must embed UnimplementedKitchenServer
// for forward compatibility
type KitchenServer interface {
	GetCookOrders(context.Context, *GetCookOrdersRequest) (*GetCookOrdersResponse, error)
	CookingStart(context.Context, *CookingStartRequest) (*CookingStartResponse, error)
	CookingEnd(context.Context, *CookingEndRequest) (*CookingEndResponse, error)
	mustEmbedUnimplementedKitchenServer()
}

// UnimplementedKitchenServer must be embedded to have forward compatible implementations.
type UnimplementedKitchenServer struct {
}

func (UnimplementedKitchenServer) GetCookOrders(context.Context, *GetCookOrdersRequest) (*GetCookOrdersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCookOrders not implemented")
}
func (UnimplementedKitchenServer) CookingStart(context.Context, *CookingStartRequest) (*CookingStartResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CookingStart not implemented")
}
func (UnimplementedKitchenServer) CookingEnd(context.Context, *CookingEndRequest) (*CookingEndResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CookingEnd not implemented")
}
func (UnimplementedKitchenServer) mustEmbedUnimplementedKitchenServer() {}

// UnsafeKitchenServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to KitchenServer will
// result in compilation errors.
type UnsafeKitchenServer interface {
	mustEmbedUnimplementedKitchenServer()
}

func RegisterKitchenServer(s grpc.ServiceRegistrar, srv KitchenServer) {
	s.RegisterService(&Kitchen_ServiceDesc, srv)
}

func _Kitchen_GetCookOrders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCookOrdersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KitchenServer).GetCookOrders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kitchen_api.v1.Kitchen/GetCookOrders",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KitchenServer).GetCookOrders(ctx, req.(*GetCookOrdersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Kitchen_CookingStart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CookingStartRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KitchenServer).CookingStart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kitchen_api.v1.Kitchen/CookingStart",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KitchenServer).CookingStart(ctx, req.(*CookingStartRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Kitchen_CookingEnd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CookingEndRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KitchenServer).CookingEnd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kitchen_api.v1.Kitchen/CookingEnd",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KitchenServer).CookingEnd(ctx, req.(*CookingEndRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Kitchen_ServiceDesc is the grpc.ServiceDesc for Kitchen service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Kitchen_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "kitchen_api.v1.Kitchen",
	HandlerType: (*KitchenServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCookOrders",
			Handler:    _Kitchen_GetCookOrders_Handler,
		},
		{
			MethodName: "CookingStart",
			Handler:    _Kitchen_CookingStart_Handler,
		},
		{
			MethodName: "CookingEnd",
			Handler:    _Kitchen_CookingEnd_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "kitchen_api.proto",
}