// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.15.8
// source: wallet/wallet.proto

package wallet

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

// WalletServiceClient is the client API for WalletService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WalletServiceClient interface {
	CreateWallet(ctx context.Context, in *CreateWalletReq, opts ...grpc.CallOption) (*CreateWalletRes, error)
	GetWallets(ctx context.Context, in *GetWalletsReq, opts ...grpc.CallOption) (*GetWalletsRes, error)
	DeleteWallet(ctx context.Context, in *DeleteWalletReq, opts ...grpc.CallOption) (*DeleteWalletRes, error)
	Transaction(ctx context.Context, in *TransactionReq, opts ...grpc.CallOption) (*TransactionRes, error)
	RollbackTransaction(ctx context.Context, in *RollbackTransactionReq, opts ...grpc.CallOption) (*RollbackTransactionRes, error)
	GetTransactionRecord(ctx context.Context, in *GetTransactionRecordReq, opts ...grpc.CallOption) (*GetTransactionRecordRes, error)
	GetTransactionRecords(ctx context.Context, in *GetTransactionRecordsReq, opts ...grpc.CallOption) (*GetTransactionRecordsRes, error)
}

type walletServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewWalletServiceClient(cc grpc.ClientConnInterface) WalletServiceClient {
	return &walletServiceClient{cc}
}

func (c *walletServiceClient) CreateWallet(ctx context.Context, in *CreateWalletReq, opts ...grpc.CallOption) (*CreateWalletRes, error) {
	out := new(CreateWalletRes)
	err := c.cc.Invoke(ctx, "/wallet.WalletService/CreateWallet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletServiceClient) GetWallets(ctx context.Context, in *GetWalletsReq, opts ...grpc.CallOption) (*GetWalletsRes, error) {
	out := new(GetWalletsRes)
	err := c.cc.Invoke(ctx, "/wallet.WalletService/GetWallets", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletServiceClient) DeleteWallet(ctx context.Context, in *DeleteWalletReq, opts ...grpc.CallOption) (*DeleteWalletRes, error) {
	out := new(DeleteWalletRes)
	err := c.cc.Invoke(ctx, "/wallet.WalletService/DeleteWallet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletServiceClient) Transaction(ctx context.Context, in *TransactionReq, opts ...grpc.CallOption) (*TransactionRes, error) {
	out := new(TransactionRes)
	err := c.cc.Invoke(ctx, "/wallet.WalletService/Transaction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletServiceClient) RollbackTransaction(ctx context.Context, in *RollbackTransactionReq, opts ...grpc.CallOption) (*RollbackTransactionRes, error) {
	out := new(RollbackTransactionRes)
	err := c.cc.Invoke(ctx, "/wallet.WalletService/RollbackTransaction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletServiceClient) GetTransactionRecord(ctx context.Context, in *GetTransactionRecordReq, opts ...grpc.CallOption) (*GetTransactionRecordRes, error) {
	out := new(GetTransactionRecordRes)
	err := c.cc.Invoke(ctx, "/wallet.WalletService/GetTransactionRecord", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletServiceClient) GetTransactionRecords(ctx context.Context, in *GetTransactionRecordsReq, opts ...grpc.CallOption) (*GetTransactionRecordsRes, error) {
	out := new(GetTransactionRecordsRes)
	err := c.cc.Invoke(ctx, "/wallet.WalletService/GetTransactionRecords", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WalletServiceServer is the server API for WalletService service.
// All implementations should embed UnimplementedWalletServiceServer
// for forward compatibility
type WalletServiceServer interface {
	CreateWallet(context.Context, *CreateWalletReq) (*CreateWalletRes, error)
	GetWallets(context.Context, *GetWalletsReq) (*GetWalletsRes, error)
	DeleteWallet(context.Context, *DeleteWalletReq) (*DeleteWalletRes, error)
	Transaction(context.Context, *TransactionReq) (*TransactionRes, error)
	RollbackTransaction(context.Context, *RollbackTransactionReq) (*RollbackTransactionRes, error)
	GetTransactionRecord(context.Context, *GetTransactionRecordReq) (*GetTransactionRecordRes, error)
	GetTransactionRecords(context.Context, *GetTransactionRecordsReq) (*GetTransactionRecordsRes, error)
}

// UnimplementedWalletServiceServer should be embedded to have forward compatible implementations.
type UnimplementedWalletServiceServer struct {
}

func (UnimplementedWalletServiceServer) CreateWallet(context.Context, *CreateWalletReq) (*CreateWalletRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateWallet not implemented")
}
func (UnimplementedWalletServiceServer) GetWallets(context.Context, *GetWalletsReq) (*GetWalletsRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWallets not implemented")
}
func (UnimplementedWalletServiceServer) DeleteWallet(context.Context, *DeleteWalletReq) (*DeleteWalletRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteWallet not implemented")
}
func (UnimplementedWalletServiceServer) Transaction(context.Context, *TransactionReq) (*TransactionRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Transaction not implemented")
}
func (UnimplementedWalletServiceServer) RollbackTransaction(context.Context, *RollbackTransactionReq) (*RollbackTransactionRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RollbackTransaction not implemented")
}
func (UnimplementedWalletServiceServer) GetTransactionRecord(context.Context, *GetTransactionRecordReq) (*GetTransactionRecordRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTransactionRecord not implemented")
}
func (UnimplementedWalletServiceServer) GetTransactionRecords(context.Context, *GetTransactionRecordsReq) (*GetTransactionRecordsRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTransactionRecords not implemented")
}

// UnsafeWalletServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WalletServiceServer will
// result in compilation errors.
type UnsafeWalletServiceServer interface {
	mustEmbedUnimplementedWalletServiceServer()
}

func RegisterWalletServiceServer(s grpc.ServiceRegistrar, srv WalletServiceServer) {
	s.RegisterService(&WalletService_ServiceDesc, srv)
}

func _WalletService_CreateWallet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateWalletReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletServiceServer).CreateWallet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wallet.WalletService/CreateWallet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletServiceServer).CreateWallet(ctx, req.(*CreateWalletReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _WalletService_GetWallets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetWalletsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletServiceServer).GetWallets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wallet.WalletService/GetWallets",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletServiceServer).GetWallets(ctx, req.(*GetWalletsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _WalletService_DeleteWallet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteWalletReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletServiceServer).DeleteWallet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wallet.WalletService/DeleteWallet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletServiceServer).DeleteWallet(ctx, req.(*DeleteWalletReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _WalletService_Transaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TransactionReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletServiceServer).Transaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wallet.WalletService/Transaction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletServiceServer).Transaction(ctx, req.(*TransactionReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _WalletService_RollbackTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RollbackTransactionReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletServiceServer).RollbackTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wallet.WalletService/RollbackTransaction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletServiceServer).RollbackTransaction(ctx, req.(*RollbackTransactionReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _WalletService_GetTransactionRecord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTransactionRecordReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletServiceServer).GetTransactionRecord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wallet.WalletService/GetTransactionRecord",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletServiceServer).GetTransactionRecord(ctx, req.(*GetTransactionRecordReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _WalletService_GetTransactionRecords_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTransactionRecordsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletServiceServer).GetTransactionRecords(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wallet.WalletService/GetTransactionRecords",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletServiceServer).GetTransactionRecords(ctx, req.(*GetTransactionRecordsReq))
	}
	return interceptor(ctx, in, info, handler)
}

// WalletService_ServiceDesc is the grpc.ServiceDesc for WalletService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var WalletService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "wallet.WalletService",
	HandlerType: (*WalletServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateWallet",
			Handler:    _WalletService_CreateWallet_Handler,
		},
		{
			MethodName: "GetWallets",
			Handler:    _WalletService_GetWallets_Handler,
		},
		{
			MethodName: "DeleteWallet",
			Handler:    _WalletService_DeleteWallet_Handler,
		},
		{
			MethodName: "Transaction",
			Handler:    _WalletService_Transaction_Handler,
		},
		{
			MethodName: "RollbackTransaction",
			Handler:    _WalletService_RollbackTransaction_Handler,
		},
		{
			MethodName: "GetTransactionRecord",
			Handler:    _WalletService_GetTransactionRecord_Handler,
		},
		{
			MethodName: "GetTransactionRecords",
			Handler:    _WalletService_GetTransactionRecords_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "wallet/wallet.proto",
}
