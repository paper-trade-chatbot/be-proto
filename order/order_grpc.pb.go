// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.15.8
// source: order/order.proto

package order

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

// OrderServiceClient is the client API for OrderService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OrderServiceClient interface {
	StartOpenPositionOrder(ctx context.Context, in *StartOpenPositionOrderReq, opts ...grpc.CallOption) (*StartOpenPositionOrderRes, error)
	FinishOpenPositionOrder(ctx context.Context, in *FinishOpenPositionOrderReq, opts ...grpc.CallOption) (*FinishOpenPositionOrderRes, error)
	StartClosePositionOrder(ctx context.Context, in *StartClosePositionOrderReq, opts ...grpc.CallOption) (*StartClosePositionOrderRes, error)
	FinishClosePositionOrder(ctx context.Context, in *FinishClosePositionOrderReq, opts ...grpc.CallOption) (*FinishClosePositionOrderRes, error)
	FailOrder(ctx context.Context, in *FailOrderReq, opts ...grpc.CallOption) (*FailOrderRes, error)
	RollbackOrder(ctx context.Context, in *RollbackOrderReq, opts ...grpc.CallOption) (*RollbackOrderRes, error)
	GetOrders(ctx context.Context, in *GetOrdersReq, opts ...grpc.CallOption) (*GetOrdersRes, error)
	UpdateOrderProcess(ctx context.Context, in *UpdateOrderProcessReq, opts ...grpc.CallOption) (*UpdateOrderProcessRes, error)
	GetOrderProcess(ctx context.Context, in *GetOrderProcessReq, opts ...grpc.CallOption) (*GetOrderProcessRes, error)
}

type orderServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewOrderServiceClient(cc grpc.ClientConnInterface) OrderServiceClient {
	return &orderServiceClient{cc}
}

func (c *orderServiceClient) StartOpenPositionOrder(ctx context.Context, in *StartOpenPositionOrderReq, opts ...grpc.CallOption) (*StartOpenPositionOrderRes, error) {
	out := new(StartOpenPositionOrderRes)
	err := c.cc.Invoke(ctx, "/order.OrderService/StartOpenPositionOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) FinishOpenPositionOrder(ctx context.Context, in *FinishOpenPositionOrderReq, opts ...grpc.CallOption) (*FinishOpenPositionOrderRes, error) {
	out := new(FinishOpenPositionOrderRes)
	err := c.cc.Invoke(ctx, "/order.OrderService/FinishOpenPositionOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) StartClosePositionOrder(ctx context.Context, in *StartClosePositionOrderReq, opts ...grpc.CallOption) (*StartClosePositionOrderRes, error) {
	out := new(StartClosePositionOrderRes)
	err := c.cc.Invoke(ctx, "/order.OrderService/StartClosePositionOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) FinishClosePositionOrder(ctx context.Context, in *FinishClosePositionOrderReq, opts ...grpc.CallOption) (*FinishClosePositionOrderRes, error) {
	out := new(FinishClosePositionOrderRes)
	err := c.cc.Invoke(ctx, "/order.OrderService/FinishClosePositionOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) FailOrder(ctx context.Context, in *FailOrderReq, opts ...grpc.CallOption) (*FailOrderRes, error) {
	out := new(FailOrderRes)
	err := c.cc.Invoke(ctx, "/order.OrderService/FailOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) RollbackOrder(ctx context.Context, in *RollbackOrderReq, opts ...grpc.CallOption) (*RollbackOrderRes, error) {
	out := new(RollbackOrderRes)
	err := c.cc.Invoke(ctx, "/order.OrderService/RollbackOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) GetOrders(ctx context.Context, in *GetOrdersReq, opts ...grpc.CallOption) (*GetOrdersRes, error) {
	out := new(GetOrdersRes)
	err := c.cc.Invoke(ctx, "/order.OrderService/GetOrders", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) UpdateOrderProcess(ctx context.Context, in *UpdateOrderProcessReq, opts ...grpc.CallOption) (*UpdateOrderProcessRes, error) {
	out := new(UpdateOrderProcessRes)
	err := c.cc.Invoke(ctx, "/order.OrderService/UpdateOrderProcess", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) GetOrderProcess(ctx context.Context, in *GetOrderProcessReq, opts ...grpc.CallOption) (*GetOrderProcessRes, error) {
	out := new(GetOrderProcessRes)
	err := c.cc.Invoke(ctx, "/order.OrderService/GetOrderProcess", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OrderServiceServer is the server API for OrderService service.
// All implementations should embed UnimplementedOrderServiceServer
// for forward compatibility
type OrderServiceServer interface {
	StartOpenPositionOrder(context.Context, *StartOpenPositionOrderReq) (*StartOpenPositionOrderRes, error)
	FinishOpenPositionOrder(context.Context, *FinishOpenPositionOrderReq) (*FinishOpenPositionOrderRes, error)
	StartClosePositionOrder(context.Context, *StartClosePositionOrderReq) (*StartClosePositionOrderRes, error)
	FinishClosePositionOrder(context.Context, *FinishClosePositionOrderReq) (*FinishClosePositionOrderRes, error)
	FailOrder(context.Context, *FailOrderReq) (*FailOrderRes, error)
	RollbackOrder(context.Context, *RollbackOrderReq) (*RollbackOrderRes, error)
	GetOrders(context.Context, *GetOrdersReq) (*GetOrdersRes, error)
	UpdateOrderProcess(context.Context, *UpdateOrderProcessReq) (*UpdateOrderProcessRes, error)
	GetOrderProcess(context.Context, *GetOrderProcessReq) (*GetOrderProcessRes, error)
}

// UnimplementedOrderServiceServer should be embedded to have forward compatible implementations.
type UnimplementedOrderServiceServer struct {
}

func (UnimplementedOrderServiceServer) StartOpenPositionOrder(context.Context, *StartOpenPositionOrderReq) (*StartOpenPositionOrderRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartOpenPositionOrder not implemented")
}
func (UnimplementedOrderServiceServer) FinishOpenPositionOrder(context.Context, *FinishOpenPositionOrderReq) (*FinishOpenPositionOrderRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FinishOpenPositionOrder not implemented")
}
func (UnimplementedOrderServiceServer) StartClosePositionOrder(context.Context, *StartClosePositionOrderReq) (*StartClosePositionOrderRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartClosePositionOrder not implemented")
}
func (UnimplementedOrderServiceServer) FinishClosePositionOrder(context.Context, *FinishClosePositionOrderReq) (*FinishClosePositionOrderRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FinishClosePositionOrder not implemented")
}
func (UnimplementedOrderServiceServer) FailOrder(context.Context, *FailOrderReq) (*FailOrderRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FailOrder not implemented")
}
func (UnimplementedOrderServiceServer) RollbackOrder(context.Context, *RollbackOrderReq) (*RollbackOrderRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RollbackOrder not implemented")
}
func (UnimplementedOrderServiceServer) GetOrders(context.Context, *GetOrdersReq) (*GetOrdersRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrders not implemented")
}
func (UnimplementedOrderServiceServer) UpdateOrderProcess(context.Context, *UpdateOrderProcessReq) (*UpdateOrderProcessRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateOrderProcess not implemented")
}
func (UnimplementedOrderServiceServer) GetOrderProcess(context.Context, *GetOrderProcessReq) (*GetOrderProcessRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrderProcess not implemented")
}

// UnsafeOrderServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OrderServiceServer will
// result in compilation errors.
type UnsafeOrderServiceServer interface {
	mustEmbedUnimplementedOrderServiceServer()
}

func RegisterOrderServiceServer(s grpc.ServiceRegistrar, srv OrderServiceServer) {
	s.RegisterService(&OrderService_ServiceDesc, srv)
}

func _OrderService_StartOpenPositionOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartOpenPositionOrderReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).StartOpenPositionOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/order.OrderService/StartOpenPositionOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).StartOpenPositionOrder(ctx, req.(*StartOpenPositionOrderReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_FinishOpenPositionOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FinishOpenPositionOrderReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).FinishOpenPositionOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/order.OrderService/FinishOpenPositionOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).FinishOpenPositionOrder(ctx, req.(*FinishOpenPositionOrderReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_StartClosePositionOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartClosePositionOrderReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).StartClosePositionOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/order.OrderService/StartClosePositionOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).StartClosePositionOrder(ctx, req.(*StartClosePositionOrderReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_FinishClosePositionOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FinishClosePositionOrderReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).FinishClosePositionOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/order.OrderService/FinishClosePositionOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).FinishClosePositionOrder(ctx, req.(*FinishClosePositionOrderReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_FailOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FailOrderReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).FailOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/order.OrderService/FailOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).FailOrder(ctx, req.(*FailOrderReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_RollbackOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RollbackOrderReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).RollbackOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/order.OrderService/RollbackOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).RollbackOrder(ctx, req.(*RollbackOrderReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_GetOrders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOrdersReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).GetOrders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/order.OrderService/GetOrders",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).GetOrders(ctx, req.(*GetOrdersReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_UpdateOrderProcess_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateOrderProcessReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).UpdateOrderProcess(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/order.OrderService/UpdateOrderProcess",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).UpdateOrderProcess(ctx, req.(*UpdateOrderProcessReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_GetOrderProcess_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOrderProcessReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).GetOrderProcess(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/order.OrderService/GetOrderProcess",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).GetOrderProcess(ctx, req.(*GetOrderProcessReq))
	}
	return interceptor(ctx, in, info, handler)
}

// OrderService_ServiceDesc is the grpc.ServiceDesc for OrderService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OrderService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "order.OrderService",
	HandlerType: (*OrderServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "StartOpenPositionOrder",
			Handler:    _OrderService_StartOpenPositionOrder_Handler,
		},
		{
			MethodName: "FinishOpenPositionOrder",
			Handler:    _OrderService_FinishOpenPositionOrder_Handler,
		},
		{
			MethodName: "StartClosePositionOrder",
			Handler:    _OrderService_StartClosePositionOrder_Handler,
		},
		{
			MethodName: "FinishClosePositionOrder",
			Handler:    _OrderService_FinishClosePositionOrder_Handler,
		},
		{
			MethodName: "FailOrder",
			Handler:    _OrderService_FailOrder_Handler,
		},
		{
			MethodName: "RollbackOrder",
			Handler:    _OrderService_RollbackOrder_Handler,
		},
		{
			MethodName: "GetOrders",
			Handler:    _OrderService_GetOrders_Handler,
		},
		{
			MethodName: "UpdateOrderProcess",
			Handler:    _OrderService_UpdateOrderProcess_Handler,
		},
		{
			MethodName: "GetOrderProcess",
			Handler:    _OrderService_GetOrderProcess_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "order/order.proto",
}
