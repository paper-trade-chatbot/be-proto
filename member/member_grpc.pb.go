// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.15.8
// source: member/member.proto

package member

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

// MemberServiceClient is the client API for MemberService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MemberServiceClient interface {
	CreateMember(ctx context.Context, in *CreateMemberReq, opts ...grpc.CallOption) (*CreateMemberRes, error)
	GetMember(ctx context.Context, in *GetMemberReq, opts ...grpc.CallOption) (*GetMemberRes, error)
	GetMembers(ctx context.Context, in *GetMembersReq, opts ...grpc.CallOption) (*GetMembersRes, error)
	ModifyMember(ctx context.Context, in *ModifyMemberReq, opts ...grpc.CallOption) (*ModifyMemberRes, error)
	ResetPassword(ctx context.Context, in *ResetPasswordReq, opts ...grpc.CallOption) (*ResetPasswordRes, error)
	DeleteMember(ctx context.Context, in *DeleteMemberReq, opts ...grpc.CallOption) (*DeleteMemberRes, error)
	CreateMemberGroup(ctx context.Context, in *CreateMemberGroupReq, opts ...grpc.CallOption) (*CreateMemberGroupRes, error)
	GetMemberGroups(ctx context.Context, in *GetMemberGroupsReq, opts ...grpc.CallOption) (*GetMemberGroupsRes, error)
	DeleteMemberGroup(ctx context.Context, in *DeleteMemberGroupReq, opts ...grpc.CallOption) (*DeleteMemberGroupRes, error)
}

type memberServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMemberServiceClient(cc grpc.ClientConnInterface) MemberServiceClient {
	return &memberServiceClient{cc}
}

func (c *memberServiceClient) CreateMember(ctx context.Context, in *CreateMemberReq, opts ...grpc.CallOption) (*CreateMemberRes, error) {
	out := new(CreateMemberRes)
	err := c.cc.Invoke(ctx, "/member.MemberService/CreateMember", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *memberServiceClient) GetMember(ctx context.Context, in *GetMemberReq, opts ...grpc.CallOption) (*GetMemberRes, error) {
	out := new(GetMemberRes)
	err := c.cc.Invoke(ctx, "/member.MemberService/GetMember", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *memberServiceClient) GetMembers(ctx context.Context, in *GetMembersReq, opts ...grpc.CallOption) (*GetMembersRes, error) {
	out := new(GetMembersRes)
	err := c.cc.Invoke(ctx, "/member.MemberService/GetMembers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *memberServiceClient) ModifyMember(ctx context.Context, in *ModifyMemberReq, opts ...grpc.CallOption) (*ModifyMemberRes, error) {
	out := new(ModifyMemberRes)
	err := c.cc.Invoke(ctx, "/member.MemberService/ModifyMember", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *memberServiceClient) ResetPassword(ctx context.Context, in *ResetPasswordReq, opts ...grpc.CallOption) (*ResetPasswordRes, error) {
	out := new(ResetPasswordRes)
	err := c.cc.Invoke(ctx, "/member.MemberService/ResetPassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *memberServiceClient) DeleteMember(ctx context.Context, in *DeleteMemberReq, opts ...grpc.CallOption) (*DeleteMemberRes, error) {
	out := new(DeleteMemberRes)
	err := c.cc.Invoke(ctx, "/member.MemberService/DeleteMember", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *memberServiceClient) CreateMemberGroup(ctx context.Context, in *CreateMemberGroupReq, opts ...grpc.CallOption) (*CreateMemberGroupRes, error) {
	out := new(CreateMemberGroupRes)
	err := c.cc.Invoke(ctx, "/member.MemberService/CreateMemberGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *memberServiceClient) GetMemberGroups(ctx context.Context, in *GetMemberGroupsReq, opts ...grpc.CallOption) (*GetMemberGroupsRes, error) {
	out := new(GetMemberGroupsRes)
	err := c.cc.Invoke(ctx, "/member.MemberService/GetMemberGroups", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *memberServiceClient) DeleteMemberGroup(ctx context.Context, in *DeleteMemberGroupReq, opts ...grpc.CallOption) (*DeleteMemberGroupRes, error) {
	out := new(DeleteMemberGroupRes)
	err := c.cc.Invoke(ctx, "/member.MemberService/DeleteMemberGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MemberServiceServer is the server API for MemberService service.
// All implementations should embed UnimplementedMemberServiceServer
// for forward compatibility
type MemberServiceServer interface {
	CreateMember(context.Context, *CreateMemberReq) (*CreateMemberRes, error)
	GetMember(context.Context, *GetMemberReq) (*GetMemberRes, error)
	GetMembers(context.Context, *GetMembersReq) (*GetMembersRes, error)
	ModifyMember(context.Context, *ModifyMemberReq) (*ModifyMemberRes, error)
	ResetPassword(context.Context, *ResetPasswordReq) (*ResetPasswordRes, error)
	DeleteMember(context.Context, *DeleteMemberReq) (*DeleteMemberRes, error)
	CreateMemberGroup(context.Context, *CreateMemberGroupReq) (*CreateMemberGroupRes, error)
	GetMemberGroups(context.Context, *GetMemberGroupsReq) (*GetMemberGroupsRes, error)
	DeleteMemberGroup(context.Context, *DeleteMemberGroupReq) (*DeleteMemberGroupRes, error)
}

// UnimplementedMemberServiceServer should be embedded to have forward compatible implementations.
type UnimplementedMemberServiceServer struct {
}

func (UnimplementedMemberServiceServer) CreateMember(context.Context, *CreateMemberReq) (*CreateMemberRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateMember not implemented")
}
func (UnimplementedMemberServiceServer) GetMember(context.Context, *GetMemberReq) (*GetMemberRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMember not implemented")
}
func (UnimplementedMemberServiceServer) GetMembers(context.Context, *GetMembersReq) (*GetMembersRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMembers not implemented")
}
func (UnimplementedMemberServiceServer) ModifyMember(context.Context, *ModifyMemberReq) (*ModifyMemberRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ModifyMember not implemented")
}
func (UnimplementedMemberServiceServer) ResetPassword(context.Context, *ResetPasswordReq) (*ResetPasswordRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResetPassword not implemented")
}
func (UnimplementedMemberServiceServer) DeleteMember(context.Context, *DeleteMemberReq) (*DeleteMemberRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteMember not implemented")
}
func (UnimplementedMemberServiceServer) CreateMemberGroup(context.Context, *CreateMemberGroupReq) (*CreateMemberGroupRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateMemberGroup not implemented")
}
func (UnimplementedMemberServiceServer) GetMemberGroups(context.Context, *GetMemberGroupsReq) (*GetMemberGroupsRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMemberGroups not implemented")
}
func (UnimplementedMemberServiceServer) DeleteMemberGroup(context.Context, *DeleteMemberGroupReq) (*DeleteMemberGroupRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteMemberGroup not implemented")
}

// UnsafeMemberServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MemberServiceServer will
// result in compilation errors.
type UnsafeMemberServiceServer interface {
	mustEmbedUnimplementedMemberServiceServer()
}

func RegisterMemberServiceServer(s grpc.ServiceRegistrar, srv MemberServiceServer) {
	s.RegisterService(&MemberService_ServiceDesc, srv)
}

func _MemberService_CreateMember_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateMemberReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MemberServiceServer).CreateMember(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/member.MemberService/CreateMember",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MemberServiceServer).CreateMember(ctx, req.(*CreateMemberReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _MemberService_GetMember_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMemberReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MemberServiceServer).GetMember(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/member.MemberService/GetMember",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MemberServiceServer).GetMember(ctx, req.(*GetMemberReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _MemberService_GetMembers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMembersReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MemberServiceServer).GetMembers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/member.MemberService/GetMembers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MemberServiceServer).GetMembers(ctx, req.(*GetMembersReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _MemberService_ModifyMember_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ModifyMemberReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MemberServiceServer).ModifyMember(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/member.MemberService/ModifyMember",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MemberServiceServer).ModifyMember(ctx, req.(*ModifyMemberReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _MemberService_ResetPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResetPasswordReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MemberServiceServer).ResetPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/member.MemberService/ResetPassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MemberServiceServer).ResetPassword(ctx, req.(*ResetPasswordReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _MemberService_DeleteMember_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteMemberReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MemberServiceServer).DeleteMember(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/member.MemberService/DeleteMember",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MemberServiceServer).DeleteMember(ctx, req.(*DeleteMemberReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _MemberService_CreateMemberGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateMemberGroupReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MemberServiceServer).CreateMemberGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/member.MemberService/CreateMemberGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MemberServiceServer).CreateMemberGroup(ctx, req.(*CreateMemberGroupReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _MemberService_GetMemberGroups_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMemberGroupsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MemberServiceServer).GetMemberGroups(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/member.MemberService/GetMemberGroups",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MemberServiceServer).GetMemberGroups(ctx, req.(*GetMemberGroupsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _MemberService_DeleteMemberGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteMemberGroupReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MemberServiceServer).DeleteMemberGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/member.MemberService/DeleteMemberGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MemberServiceServer).DeleteMemberGroup(ctx, req.(*DeleteMemberGroupReq))
	}
	return interceptor(ctx, in, info, handler)
}

// MemberService_ServiceDesc is the grpc.ServiceDesc for MemberService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MemberService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "member.MemberService",
	HandlerType: (*MemberServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateMember",
			Handler:    _MemberService_CreateMember_Handler,
		},
		{
			MethodName: "GetMember",
			Handler:    _MemberService_GetMember_Handler,
		},
		{
			MethodName: "GetMembers",
			Handler:    _MemberService_GetMembers_Handler,
		},
		{
			MethodName: "ModifyMember",
			Handler:    _MemberService_ModifyMember_Handler,
		},
		{
			MethodName: "ResetPassword",
			Handler:    _MemberService_ResetPassword_Handler,
		},
		{
			MethodName: "DeleteMember",
			Handler:    _MemberService_DeleteMember_Handler,
		},
		{
			MethodName: "CreateMemberGroup",
			Handler:    _MemberService_CreateMemberGroup_Handler,
		},
		{
			MethodName: "GetMemberGroups",
			Handler:    _MemberService_GetMemberGroups_Handler,
		},
		{
			MethodName: "DeleteMemberGroup",
			Handler:    _MemberService_DeleteMemberGroup_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "member/member.proto",
}
