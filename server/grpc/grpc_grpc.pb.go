// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: grpc.proto

package grpc

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

// UsersServiceClient is the client API for UsersService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UsersServiceClient interface {
	Auth(ctx context.Context, in *AuthRequest, opts ...grpc.CallOption) (*User, error)
	GetUsers(ctx context.Context, in *Empty, opts ...grpc.CallOption) (UsersService_GetUsersClient, error)
	CreateUser(ctx context.Context, opts ...grpc.CallOption) (UsersService_CreateUserClient, error)
	DeleteUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*User, error)
}

type usersServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUsersServiceClient(cc grpc.ClientConnInterface) UsersServiceClient {
	return &usersServiceClient{cc}
}

func (c *usersServiceClient) Auth(ctx context.Context, in *AuthRequest, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/grpc.UsersService/Auth", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersServiceClient) GetUsers(ctx context.Context, in *Empty, opts ...grpc.CallOption) (UsersService_GetUsersClient, error) {
	stream, err := c.cc.NewStream(ctx, &UsersService_ServiceDesc.Streams[0], "/grpc.UsersService/GetUsers", opts...)
	if err != nil {
		return nil, err
	}
	x := &usersServiceGetUsersClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type UsersService_GetUsersClient interface {
	Recv() (*User, error)
	grpc.ClientStream
}

type usersServiceGetUsersClient struct {
	grpc.ClientStream
}

func (x *usersServiceGetUsersClient) Recv() (*User, error) {
	m := new(User)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *usersServiceClient) CreateUser(ctx context.Context, opts ...grpc.CallOption) (UsersService_CreateUserClient, error) {
	stream, err := c.cc.NewStream(ctx, &UsersService_ServiceDesc.Streams[1], "/grpc.UsersService/CreateUser", opts...)
	if err != nil {
		return nil, err
	}
	x := &usersServiceCreateUserClient{stream}
	return x, nil
}

type UsersService_CreateUserClient interface {
	Send(*User) error
	Recv() (*User, error)
	grpc.ClientStream
}

type usersServiceCreateUserClient struct {
	grpc.ClientStream
}

func (x *usersServiceCreateUserClient) Send(m *User) error {
	return x.ClientStream.SendMsg(m)
}

func (x *usersServiceCreateUserClient) Recv() (*User, error) {
	m := new(User)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *usersServiceClient) DeleteUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/grpc.UsersService/DeleteUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UsersServiceServer is the server API for UsersService service.
// All implementations must embed UnimplementedUsersServiceServer
// for forward compatibility
type UsersServiceServer interface {
	Auth(context.Context, *AuthRequest) (*User, error)
	GetUsers(*Empty, UsersService_GetUsersServer) error
	CreateUser(UsersService_CreateUserServer) error
	DeleteUser(context.Context, *User) (*User, error)
	mustEmbedUnimplementedUsersServiceServer()
}

// UnimplementedUsersServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUsersServiceServer struct {
}

func (UnimplementedUsersServiceServer) Auth(context.Context, *AuthRequest) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Auth not implemented")
}
func (UnimplementedUsersServiceServer) GetUsers(*Empty, UsersService_GetUsersServer) error {
	return status.Errorf(codes.Unimplemented, "method GetUsers not implemented")
}
func (UnimplementedUsersServiceServer) CreateUser(UsersService_CreateUserServer) error {
	return status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (UnimplementedUsersServiceServer) DeleteUser(context.Context, *User) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUser not implemented")
}
func (UnimplementedUsersServiceServer) mustEmbedUnimplementedUsersServiceServer() {}

// UnsafeUsersServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UsersServiceServer will
// result in compilation errors.
type UnsafeUsersServiceServer interface {
	mustEmbedUnimplementedUsersServiceServer()
}

func RegisterUsersServiceServer(s grpc.ServiceRegistrar, srv UsersServiceServer) {
	s.RegisterService(&UsersService_ServiceDesc, srv)
}

func _UsersService_Auth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServiceServer).Auth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.UsersService/Auth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServiceServer).Auth(ctx, req.(*AuthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UsersService_GetUsers_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(UsersServiceServer).GetUsers(m, &usersServiceGetUsersServer{stream})
}

type UsersService_GetUsersServer interface {
	Send(*User) error
	grpc.ServerStream
}

type usersServiceGetUsersServer struct {
	grpc.ServerStream
}

func (x *usersServiceGetUsersServer) Send(m *User) error {
	return x.ServerStream.SendMsg(m)
}

func _UsersService_CreateUser_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(UsersServiceServer).CreateUser(&usersServiceCreateUserServer{stream})
}

type UsersService_CreateUserServer interface {
	Send(*User) error
	Recv() (*User, error)
	grpc.ServerStream
}

type usersServiceCreateUserServer struct {
	grpc.ServerStream
}

func (x *usersServiceCreateUserServer) Send(m *User) error {
	return x.ServerStream.SendMsg(m)
}

func (x *usersServiceCreateUserServer) Recv() (*User, error) {
	m := new(User)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _UsersService_DeleteUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServiceServer).DeleteUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.UsersService/DeleteUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServiceServer).DeleteUser(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

// UsersService_ServiceDesc is the grpc.ServiceDesc for UsersService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UsersService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.UsersService",
	HandlerType: (*UsersServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Auth",
			Handler:    _UsersService_Auth_Handler,
		},
		{
			MethodName: "DeleteUser",
			Handler:    _UsersService_DeleteUser_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetUsers",
			Handler:       _UsersService_GetUsers_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "CreateUser",
			Handler:       _UsersService_CreateUser_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "grpc.proto",
}

// FilesServiceClient is the client API for FilesService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FilesServiceClient interface {
	GetFiles(ctx context.Context, in *TargetId, opts ...grpc.CallOption) (FilesService_GetFilesClient, error)
	CreateFile(ctx context.Context, in *File, opts ...grpc.CallOption) (*File, error)
	DeleteFile(ctx context.Context, in *FileId, opts ...grpc.CallOption) (*File, error)
}

type filesServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFilesServiceClient(cc grpc.ClientConnInterface) FilesServiceClient {
	return &filesServiceClient{cc}
}

func (c *filesServiceClient) GetFiles(ctx context.Context, in *TargetId, opts ...grpc.CallOption) (FilesService_GetFilesClient, error) {
	stream, err := c.cc.NewStream(ctx, &FilesService_ServiceDesc.Streams[0], "/grpc.FilesService/GetFiles", opts...)
	if err != nil {
		return nil, err
	}
	x := &filesServiceGetFilesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type FilesService_GetFilesClient interface {
	Recv() (*File, error)
	grpc.ClientStream
}

type filesServiceGetFilesClient struct {
	grpc.ClientStream
}

func (x *filesServiceGetFilesClient) Recv() (*File, error) {
	m := new(File)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *filesServiceClient) CreateFile(ctx context.Context, in *File, opts ...grpc.CallOption) (*File, error) {
	out := new(File)
	err := c.cc.Invoke(ctx, "/grpc.FilesService/CreateFile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *filesServiceClient) DeleteFile(ctx context.Context, in *FileId, opts ...grpc.CallOption) (*File, error) {
	out := new(File)
	err := c.cc.Invoke(ctx, "/grpc.FilesService/DeleteFile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FilesServiceServer is the server API for FilesService service.
// All implementations must embed UnimplementedFilesServiceServer
// for forward compatibility
type FilesServiceServer interface {
	GetFiles(*TargetId, FilesService_GetFilesServer) error
	CreateFile(context.Context, *File) (*File, error)
	DeleteFile(context.Context, *FileId) (*File, error)
	mustEmbedUnimplementedFilesServiceServer()
}

// UnimplementedFilesServiceServer must be embedded to have forward compatible implementations.
type UnimplementedFilesServiceServer struct {
}

func (UnimplementedFilesServiceServer) GetFiles(*TargetId, FilesService_GetFilesServer) error {
	return status.Errorf(codes.Unimplemented, "method GetFiles not implemented")
}
func (UnimplementedFilesServiceServer) CreateFile(context.Context, *File) (*File, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateFile not implemented")
}
func (UnimplementedFilesServiceServer) DeleteFile(context.Context, *FileId) (*File, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteFile not implemented")
}
func (UnimplementedFilesServiceServer) mustEmbedUnimplementedFilesServiceServer() {}

// UnsafeFilesServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FilesServiceServer will
// result in compilation errors.
type UnsafeFilesServiceServer interface {
	mustEmbedUnimplementedFilesServiceServer()
}

func RegisterFilesServiceServer(s grpc.ServiceRegistrar, srv FilesServiceServer) {
	s.RegisterService(&FilesService_ServiceDesc, srv)
}

func _FilesService_GetFiles_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(TargetId)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(FilesServiceServer).GetFiles(m, &filesServiceGetFilesServer{stream})
}

type FilesService_GetFilesServer interface {
	Send(*File) error
	grpc.ServerStream
}

type filesServiceGetFilesServer struct {
	grpc.ServerStream
}

func (x *filesServiceGetFilesServer) Send(m *File) error {
	return x.ServerStream.SendMsg(m)
}

func _FilesService_CreateFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(File)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FilesServiceServer).CreateFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.FilesService/CreateFile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FilesServiceServer).CreateFile(ctx, req.(*File))
	}
	return interceptor(ctx, in, info, handler)
}

func _FilesService_DeleteFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FileId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FilesServiceServer).DeleteFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.FilesService/DeleteFile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FilesServiceServer).DeleteFile(ctx, req.(*FileId))
	}
	return interceptor(ctx, in, info, handler)
}

// FilesService_ServiceDesc is the grpc.ServiceDesc for FilesService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FilesService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.FilesService",
	HandlerType: (*FilesServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateFile",
			Handler:    _FilesService_CreateFile_Handler,
		},
		{
			MethodName: "DeleteFile",
			Handler:    _FilesService_DeleteFile_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetFiles",
			Handler:       _FilesService_GetFiles_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "grpc.proto",
}
