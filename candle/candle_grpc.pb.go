// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.15.8
// source: candle/candle.proto

package candle

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

// CandleServiceClient is the client API for CandleService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CandleServiceClient interface {
	CreateCandles(ctx context.Context, in *CreateCandlesReq, opts ...grpc.CallOption) (*CreateCandlesRes, error)
	GetCandles(ctx context.Context, in *GetCandlesReq, opts ...grpc.CallOption) (*GetCandlesRes, error)
}

type candleServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCandleServiceClient(cc grpc.ClientConnInterface) CandleServiceClient {
	return &candleServiceClient{cc}
}

func (c *candleServiceClient) CreateCandles(ctx context.Context, in *CreateCandlesReq, opts ...grpc.CallOption) (*CreateCandlesRes, error) {
	out := new(CreateCandlesRes)
	err := c.cc.Invoke(ctx, "/candle.CandleService/CreateCandles", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *candleServiceClient) GetCandles(ctx context.Context, in *GetCandlesReq, opts ...grpc.CallOption) (*GetCandlesRes, error) {
	out := new(GetCandlesRes)
	err := c.cc.Invoke(ctx, "/candle.CandleService/GetCandles", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CandleServiceServer is the server API for CandleService service.
// All implementations should embed UnimplementedCandleServiceServer
// for forward compatibility
type CandleServiceServer interface {
	CreateCandles(context.Context, *CreateCandlesReq) (*CreateCandlesRes, error)
	GetCandles(context.Context, *GetCandlesReq) (*GetCandlesRes, error)
}

// UnimplementedCandleServiceServer should be embedded to have forward compatible implementations.
type UnimplementedCandleServiceServer struct {
}

func (UnimplementedCandleServiceServer) CreateCandles(context.Context, *CreateCandlesReq) (*CreateCandlesRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCandles not implemented")
}
func (UnimplementedCandleServiceServer) GetCandles(context.Context, *GetCandlesReq) (*GetCandlesRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCandles not implemented")
}

// UnsafeCandleServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CandleServiceServer will
// result in compilation errors.
type UnsafeCandleServiceServer interface {
	mustEmbedUnimplementedCandleServiceServer()
}

func RegisterCandleServiceServer(s grpc.ServiceRegistrar, srv CandleServiceServer) {
	s.RegisterService(&CandleService_ServiceDesc, srv)
}

func _CandleService_CreateCandles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCandlesReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CandleServiceServer).CreateCandles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/candle.CandleService/CreateCandles",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CandleServiceServer).CreateCandles(ctx, req.(*CreateCandlesReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _CandleService_GetCandles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCandlesReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CandleServiceServer).GetCandles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/candle.CandleService/GetCandles",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CandleServiceServer).GetCandles(ctx, req.(*GetCandlesReq))
	}
	return interceptor(ctx, in, info, handler)
}

// CandleService_ServiceDesc is the grpc.ServiceDesc for CandleService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CandleService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "candle.CandleService",
	HandlerType: (*CandleServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateCandles",
			Handler:    _CandleService_CreateCandles_Handler,
		},
		{
			MethodName: "GetCandles",
			Handler:    _CandleService_GetCandles_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "candle/candle.proto",
}
