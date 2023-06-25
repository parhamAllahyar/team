// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.10
// source: proto.proto

package grpcpkg

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

// AdminClient is the client API for Admin service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AdminClient interface {
	SignUp(ctx context.Context, in *SignUpRequest, opts ...grpc.CallOption) (*SignUpRequest, error)
	SignIn(ctx context.Context, in *SignInRequest, opts ...grpc.CallOption) (*SignInResponse, error)
	Logout(ctx context.Context, in *LogoutRequest, opts ...grpc.CallOption) (*LogoutResponse, error)
	Index(ctx context.Context, in *IndexAdminRequest, opts ...grpc.CallOption) (*IndexAdminResponse, error)
}

type adminClient struct {
	cc grpc.ClientConnInterface
}

func NewAdminClient(cc grpc.ClientConnInterface) AdminClient {
	return &adminClient{cc}
}

func (c *adminClient) SignUp(ctx context.Context, in *SignUpRequest, opts ...grpc.CallOption) (*SignUpRequest, error) {
	out := new(SignUpRequest)
	err := c.cc.Invoke(ctx, "/grpcpkg.Admin/SignUp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminClient) SignIn(ctx context.Context, in *SignInRequest, opts ...grpc.CallOption) (*SignInResponse, error) {
	out := new(SignInResponse)
	err := c.cc.Invoke(ctx, "/grpcpkg.Admin/SignIn", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminClient) Logout(ctx context.Context, in *LogoutRequest, opts ...grpc.CallOption) (*LogoutResponse, error) {
	out := new(LogoutResponse)
	err := c.cc.Invoke(ctx, "/grpcpkg.Admin/Logout", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminClient) Index(ctx context.Context, in *IndexAdminRequest, opts ...grpc.CallOption) (*IndexAdminResponse, error) {
	out := new(IndexAdminResponse)
	err := c.cc.Invoke(ctx, "/grpcpkg.Admin/Index", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AdminServer is the server API for Admin service.
// All implementations must embed UnimplementedAdminServer
// for forward compatibility
type AdminServer interface {
	SignUp(context.Context, *SignUpRequest) (*SignUpRequest, error)
	SignIn(context.Context, *SignInRequest) (*SignInResponse, error)
	Logout(context.Context, *LogoutRequest) (*LogoutResponse, error)
	Index(context.Context, *IndexAdminRequest) (*IndexAdminResponse, error)
	mustEmbedUnimplementedAdminServer()
}

// UnimplementedAdminServer must be embedded to have forward compatible implementations.
type UnimplementedAdminServer struct {
}

func (UnimplementedAdminServer) SignUp(context.Context, *SignUpRequest) (*SignUpRequest, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignUp not implemented")
}
func (UnimplementedAdminServer) SignIn(context.Context, *SignInRequest) (*SignInResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignIn not implemented")
}
func (UnimplementedAdminServer) Logout(context.Context, *LogoutRequest) (*LogoutResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Logout not implemented")
}
func (UnimplementedAdminServer) Index(context.Context, *IndexAdminRequest) (*IndexAdminResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Index not implemented")
}
func (UnimplementedAdminServer) mustEmbedUnimplementedAdminServer() {}

// UnsafeAdminServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AdminServer will
// result in compilation errors.
type UnsafeAdminServer interface {
	mustEmbedUnimplementedAdminServer()
}

func RegisterAdminServer(s grpc.ServiceRegistrar, srv AdminServer) {
	s.RegisterService(&Admin_ServiceDesc, srv)
}

func _Admin_SignUp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignUpRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).SignUp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpcpkg.Admin/SignUp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).SignUp(ctx, req.(*SignUpRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Admin_SignIn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignInRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).SignIn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpcpkg.Admin/SignIn",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).SignIn(ctx, req.(*SignInRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Admin_Logout_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogoutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).Logout(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpcpkg.Admin/Logout",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).Logout(ctx, req.(*LogoutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Admin_Index_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IndexAdminRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).Index(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpcpkg.Admin/Index",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).Index(ctx, req.(*IndexAdminRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Admin_ServiceDesc is the grpc.ServiceDesc for Admin service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Admin_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpcpkg.Admin",
	HandlerType: (*AdminServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SignUp",
			Handler:    _Admin_SignUp_Handler,
		},
		{
			MethodName: "SignIn",
			Handler:    _Admin_SignIn_Handler,
		},
		{
			MethodName: "Logout",
			Handler:    _Admin_Logout_Handler,
		},
		{
			MethodName: "Index",
			Handler:    _Admin_Index_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto.proto",
}

// BoardClient is the client API for Board service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BoardClient interface {
	GetById(ctx context.Context, in *GetBoardRequest, opts ...grpc.CallOption) (*GetBoardResponse, error)
	Index(ctx context.Context, in *IndexBoardRequest, opts ...grpc.CallOption) (*IndexBoardResponse, error)
}

type boardClient struct {
	cc grpc.ClientConnInterface
}

func NewBoardClient(cc grpc.ClientConnInterface) BoardClient {
	return &boardClient{cc}
}

func (c *boardClient) GetById(ctx context.Context, in *GetBoardRequest, opts ...grpc.CallOption) (*GetBoardResponse, error) {
	out := new(GetBoardResponse)
	err := c.cc.Invoke(ctx, "/grpcpkg.Board/GetById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *boardClient) Index(ctx context.Context, in *IndexBoardRequest, opts ...grpc.CallOption) (*IndexBoardResponse, error) {
	out := new(IndexBoardResponse)
	err := c.cc.Invoke(ctx, "/grpcpkg.Board/Index", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BoardServer is the server API for Board service.
// All implementations must embed UnimplementedBoardServer
// for forward compatibility
type BoardServer interface {
	GetById(context.Context, *GetBoardRequest) (*GetBoardResponse, error)
	Index(context.Context, *IndexBoardRequest) (*IndexBoardResponse, error)
	mustEmbedUnimplementedBoardServer()
}

// UnimplementedBoardServer must be embedded to have forward compatible implementations.
type UnimplementedBoardServer struct {
}

func (UnimplementedBoardServer) GetById(context.Context, *GetBoardRequest) (*GetBoardResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetById not implemented")
}
func (UnimplementedBoardServer) Index(context.Context, *IndexBoardRequest) (*IndexBoardResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Index not implemented")
}
func (UnimplementedBoardServer) mustEmbedUnimplementedBoardServer() {}

// UnsafeBoardServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BoardServer will
// result in compilation errors.
type UnsafeBoardServer interface {
	mustEmbedUnimplementedBoardServer()
}

func RegisterBoardServer(s grpc.ServiceRegistrar, srv BoardServer) {
	s.RegisterService(&Board_ServiceDesc, srv)
}

func _Board_GetById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBoardRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BoardServer).GetById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpcpkg.Board/GetById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BoardServer).GetById(ctx, req.(*GetBoardRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Board_Index_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IndexBoardRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BoardServer).Index(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpcpkg.Board/Index",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BoardServer).Index(ctx, req.(*IndexBoardRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Board_ServiceDesc is the grpc.ServiceDesc for Board service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Board_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpcpkg.Board",
	HandlerType: (*BoardServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetById",
			Handler:    _Board_GetById_Handler,
		},
		{
			MethodName: "Index",
			Handler:    _Board_Index_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto.proto",
}

// SectionClient is the client API for Section service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SectionClient interface {
	GetById(ctx context.Context, in *GetSectionRequest, opts ...grpc.CallOption) (*GetSectionResponse, error)
	Index(ctx context.Context, in *IndexSectionRequest, opts ...grpc.CallOption) (*IndexSectionResponse, error)
}

type sectionClient struct {
	cc grpc.ClientConnInterface
}

func NewSectionClient(cc grpc.ClientConnInterface) SectionClient {
	return &sectionClient{cc}
}

func (c *sectionClient) GetById(ctx context.Context, in *GetSectionRequest, opts ...grpc.CallOption) (*GetSectionResponse, error) {
	out := new(GetSectionResponse)
	err := c.cc.Invoke(ctx, "/grpcpkg.Section/GetById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sectionClient) Index(ctx context.Context, in *IndexSectionRequest, opts ...grpc.CallOption) (*IndexSectionResponse, error) {
	out := new(IndexSectionResponse)
	err := c.cc.Invoke(ctx, "/grpcpkg.Section/Index", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SectionServer is the server API for Section service.
// All implementations must embed UnimplementedSectionServer
// for forward compatibility
type SectionServer interface {
	GetById(context.Context, *GetSectionRequest) (*GetSectionResponse, error)
	Index(context.Context, *IndexSectionRequest) (*IndexSectionResponse, error)
	mustEmbedUnimplementedSectionServer()
}

// UnimplementedSectionServer must be embedded to have forward compatible implementations.
type UnimplementedSectionServer struct {
}

func (UnimplementedSectionServer) GetById(context.Context, *GetSectionRequest) (*GetSectionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetById not implemented")
}
func (UnimplementedSectionServer) Index(context.Context, *IndexSectionRequest) (*IndexSectionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Index not implemented")
}
func (UnimplementedSectionServer) mustEmbedUnimplementedSectionServer() {}

// UnsafeSectionServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SectionServer will
// result in compilation errors.
type UnsafeSectionServer interface {
	mustEmbedUnimplementedSectionServer()
}

func RegisterSectionServer(s grpc.ServiceRegistrar, srv SectionServer) {
	s.RegisterService(&Section_ServiceDesc, srv)
}

func _Section_GetById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSectionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SectionServer).GetById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpcpkg.Section/GetById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SectionServer).GetById(ctx, req.(*GetSectionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Section_Index_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IndexSectionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SectionServer).Index(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpcpkg.Section/Index",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SectionServer).Index(ctx, req.(*IndexSectionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Section_ServiceDesc is the grpc.ServiceDesc for Section service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Section_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpcpkg.Section",
	HandlerType: (*SectionServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetById",
			Handler:    _Section_GetById_Handler,
		},
		{
			MethodName: "Index",
			Handler:    _Section_Index_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto.proto",
}

// TaskClient is the client API for Task service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TaskClient interface {
	GetById(ctx context.Context, in *GetBoardRequest, opts ...grpc.CallOption) (*GetBoardResponse, error)
	Index(ctx context.Context, in *IndexTaskRequest, opts ...grpc.CallOption) (*IndexTaskResponse, error)
}

type taskClient struct {
	cc grpc.ClientConnInterface
}

func NewTaskClient(cc grpc.ClientConnInterface) TaskClient {
	return &taskClient{cc}
}

func (c *taskClient) GetById(ctx context.Context, in *GetBoardRequest, opts ...grpc.CallOption) (*GetBoardResponse, error) {
	out := new(GetBoardResponse)
	err := c.cc.Invoke(ctx, "/grpcpkg.Task/GetById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskClient) Index(ctx context.Context, in *IndexTaskRequest, opts ...grpc.CallOption) (*IndexTaskResponse, error) {
	out := new(IndexTaskResponse)
	err := c.cc.Invoke(ctx, "/grpcpkg.Task/Index", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TaskServer is the server API for Task service.
// All implementations must embed UnimplementedTaskServer
// for forward compatibility
type TaskServer interface {
	GetById(context.Context, *GetBoardRequest) (*GetBoardResponse, error)
	Index(context.Context, *IndexTaskRequest) (*IndexTaskResponse, error)
	mustEmbedUnimplementedTaskServer()
}

// UnimplementedTaskServer must be embedded to have forward compatible implementations.
type UnimplementedTaskServer struct {
}

func (UnimplementedTaskServer) GetById(context.Context, *GetBoardRequest) (*GetBoardResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetById not implemented")
}
func (UnimplementedTaskServer) Index(context.Context, *IndexTaskRequest) (*IndexTaskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Index not implemented")
}
func (UnimplementedTaskServer) mustEmbedUnimplementedTaskServer() {}

// UnsafeTaskServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TaskServer will
// result in compilation errors.
type UnsafeTaskServer interface {
	mustEmbedUnimplementedTaskServer()
}

func RegisterTaskServer(s grpc.ServiceRegistrar, srv TaskServer) {
	s.RegisterService(&Task_ServiceDesc, srv)
}

func _Task_GetById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBoardRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServer).GetById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpcpkg.Task/GetById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServer).GetById(ctx, req.(*GetBoardRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Task_Index_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IndexTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServer).Index(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpcpkg.Task/Index",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServer).Index(ctx, req.(*IndexTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Task_ServiceDesc is the grpc.ServiceDesc for Task service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Task_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpcpkg.Task",
	HandlerType: (*TaskServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetById",
			Handler:    _Task_GetById_Handler,
		},
		{
			MethodName: "Index",
			Handler:    _Task_Index_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto.proto",
}

// CustomerClient is the client API for Customer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CustomerClient interface {
	GetById(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
	Index(ctx context.Context, in *IndexRequest, opts ...grpc.CallOption) (*IndexResponse, error)
	UpdateStatus(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error)
}

type customerClient struct {
	cc grpc.ClientConnInterface
}

func NewCustomerClient(cc grpc.ClientConnInterface) CustomerClient {
	return &customerClient{cc}
}

func (c *customerClient) GetById(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, "/grpcpkg.Customer/GetById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerClient) Index(ctx context.Context, in *IndexRequest, opts ...grpc.CallOption) (*IndexResponse, error) {
	out := new(IndexResponse)
	err := c.cc.Invoke(ctx, "/grpcpkg.Customer/Index", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerClient) UpdateStatus(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error) {
	out := new(UpdateResponse)
	err := c.cc.Invoke(ctx, "/grpcpkg.Customer/UpdateStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CustomerServer is the server API for Customer service.
// All implementations must embed UnimplementedCustomerServer
// for forward compatibility
type CustomerServer interface {
	GetById(context.Context, *GetRequest) (*GetResponse, error)
	Index(context.Context, *IndexRequest) (*IndexResponse, error)
	UpdateStatus(context.Context, *UpdateRequest) (*UpdateResponse, error)
	mustEmbedUnimplementedCustomerServer()
}

// UnimplementedCustomerServer must be embedded to have forward compatible implementations.
type UnimplementedCustomerServer struct {
}

func (UnimplementedCustomerServer) GetById(context.Context, *GetRequest) (*GetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetById not implemented")
}
func (UnimplementedCustomerServer) Index(context.Context, *IndexRequest) (*IndexResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Index not implemented")
}
func (UnimplementedCustomerServer) UpdateStatus(context.Context, *UpdateRequest) (*UpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateStatus not implemented")
}
func (UnimplementedCustomerServer) mustEmbedUnimplementedCustomerServer() {}

// UnsafeCustomerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CustomerServer will
// result in compilation errors.
type UnsafeCustomerServer interface {
	mustEmbedUnimplementedCustomerServer()
}

func RegisterCustomerServer(s grpc.ServiceRegistrar, srv CustomerServer) {
	s.RegisterService(&Customer_ServiceDesc, srv)
}

func _Customer_GetById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerServer).GetById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpcpkg.Customer/GetById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerServer).GetById(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Customer_Index_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IndexRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerServer).Index(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpcpkg.Customer/Index",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerServer).Index(ctx, req.(*IndexRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Customer_UpdateStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerServer).UpdateStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpcpkg.Customer/UpdateStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerServer).UpdateStatus(ctx, req.(*UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Customer_ServiceDesc is the grpc.ServiceDesc for Customer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Customer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpcpkg.Customer",
	HandlerType: (*CustomerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetById",
			Handler:    _Customer_GetById_Handler,
		},
		{
			MethodName: "Index",
			Handler:    _Customer_Index_Handler,
		},
		{
			MethodName: "UpdateStatus",
			Handler:    _Customer_UpdateStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto.proto",
}