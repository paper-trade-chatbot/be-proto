// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.15.8
// source: position/position.proto

package position

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

// PositionServiceClient is the client API for PositionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PositionServiceClient interface {
	OpenPosition(ctx context.Context, in *OpenPositionReq, opts ...grpc.CallOption) (*OpenPositionRes, error)
	ClosePosition(ctx context.Context, in *ClosePositionReq, opts ...grpc.CallOption) (*ClosePositionRes, error)
	GetPositions(ctx context.Context, in *GetPositionsReq, opts ...grpc.CallOption) (*GetPositionsRes, error)
	ModifyPosition(ctx context.Context, in *ModifyPositionReq, opts ...grpc.CallOption) (*ModifyPositionRes, error)
}

type positionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPositionServiceClient(cc grpc.ClientConnInterface) PositionServiceClient {
	return &positionServiceClient{cc}
}

func (c *positionServiceClient) OpenPosition(ctx context.Context, in *OpenPositionReq, opts ...grpc.CallOption) (*OpenPositionRes, error) {
	out := new(OpenPositionRes)
	err := c.cc.Invoke(ctx, "/position.PositionService/OpenPosition", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *positionServiceClient) ClosePosition(ctx context.Context, in *ClosePositionReq, opts ...grpc.CallOption) (*ClosePositionRes, error) {
	out := new(ClosePositionRes)
	err := c.cc.Invoke(ctx, "/position.PositionService/ClosePosition", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *positionServiceClient) GetPositions(ctx context.Context, in *GetPositionsReq, opts ...grpc.CallOption) (*GetPositionsRes, error) {
	out := new(GetPositionsRes)
	err := c.cc.Invoke(ctx, "/position.PositionService/GetPositions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *positionServiceClient) ModifyPosition(ctx context.Context, in *ModifyPositionReq, opts ...grpc.CallOption) (*ModifyPositionRes, error) {
	out := new(ModifyPositionRes)
	err := c.cc.Invoke(ctx, "/position.PositionService/ModifyPosition", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PositionServiceServer is the server API for PositionService service.
// All implementations should embed UnimplementedPositionServiceServer
// for forward compatibility
type PositionServiceServer interface {
	OpenPosition(context.Context, *OpenPositionReq) (*OpenPositionRes, error)
	ClosePosition(context.Context, *ClosePositionReq) (*ClosePositionRes, error)
	GetPositions(context.Context, *GetPositionsReq) (*GetPositionsRes, error)
	ModifyPosition(context.Context, *ModifyPositionReq) (*ModifyPositionRes, error)
}

// UnimplementedPositionServiceServer should be embedded to have forward compatible implementations.
type UnimplementedPositionServiceServer struct {
}

func (UnimplementedPositionServiceServer) OpenPosition(context.Context, *OpenPositionReq) (*OpenPositionRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OpenPosition not implemented")
}
func (UnimplementedPositionServiceServer) ClosePosition(context.Context, *ClosePositionReq) (*ClosePositionRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ClosePosition not implemented")
}
func (UnimplementedPositionServiceServer) GetPositions(context.Context, *GetPositionsReq) (*GetPositionsRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPositions not implemented")
}
func (UnimplementedPositionServiceServer) ModifyPosition(context.Context, *ModifyPositionReq) (*ModifyPositionRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ModifyPosition not implemented")
}

// UnsafePositionServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PositionServiceServer will
// result in compilation errors.
type UnsafePositionServiceServer interface {
	mustEmbedUnimplementedPositionServiceServer()
}

func RegisterPositionServiceServer(s grpc.ServiceRegistrar, srv PositionServiceServer) {
	s.RegisterService(&PositionService_ServiceDesc, srv)
}

func _PositionService_OpenPosition_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OpenPositionReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PositionServiceServer).OpenPosition(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/position.PositionService/OpenPosition",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PositionServiceServer).OpenPosition(ctx, req.(*OpenPositionReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _PositionService_ClosePosition_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClosePositionReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PositionServiceServer).ClosePosition(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/position.PositionService/ClosePosition",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PositionServiceServer).ClosePosition(ctx, req.(*ClosePositionReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _PositionService_GetPositions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPositionsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PositionServiceServer).GetPositions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/position.PositionService/GetPositions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PositionServiceServer).GetPositions(ctx, req.(*GetPositionsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _PositionService_ModifyPosition_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ModifyPositionReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PositionServiceServer).ModifyPosition(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/position.PositionService/ModifyPosition",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PositionServiceServer).ModifyPosition(ctx, req.(*ModifyPositionReq))
	}
	return interceptor(ctx, in, info, handler)
}

// PositionService_ServiceDesc is the grpc.ServiceDesc for PositionService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PositionService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "position.PositionService",
	HandlerType: (*PositionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "OpenPosition",
			Handler:    _PositionService_OpenPosition_Handler,
		},
		{
			MethodName: "ClosePosition",
			Handler:    _PositionService_ClosePosition_Handler,
		},
		{
			MethodName: "GetPositions",
			Handler:    _PositionService_GetPositions_Handler,
		},
		{
			MethodName: "ModifyPosition",
			Handler:    _PositionService_ModifyPosition_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "position/position.proto",
}