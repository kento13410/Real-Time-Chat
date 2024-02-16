// Code generated by entproto. DO NOT EDIT.

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: entpb/entpb.proto

package entpb

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	ChatService_Create_FullMethodName      = "/entpb.ChatService/Create"
	ChatService_Get_FullMethodName         = "/entpb.ChatService/Get"
	ChatService_Update_FullMethodName      = "/entpb.ChatService/Update"
	ChatService_Delete_FullMethodName      = "/entpb.ChatService/Delete"
	ChatService_List_FullMethodName        = "/entpb.ChatService/List"
	ChatService_BatchCreate_FullMethodName = "/entpb.ChatService/BatchCreate"
)

// ChatServiceClient is the client API for ChatService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ChatServiceClient interface {
	Create(ctx context.Context, in *CreateChatRequest, opts ...grpc.CallOption) (*Chat, error)
	Get(ctx context.Context, in *GetChatRequest, opts ...grpc.CallOption) (*Chat, error)
	Update(ctx context.Context, in *UpdateChatRequest, opts ...grpc.CallOption) (*Chat, error)
	Delete(ctx context.Context, in *DeleteChatRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	List(ctx context.Context, in *ListChatRequest, opts ...grpc.CallOption) (*ListChatResponse, error)
	BatchCreate(ctx context.Context, in *BatchCreateChatsRequest, opts ...grpc.CallOption) (*BatchCreateChatsResponse, error)
}

type chatServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewChatServiceClient(cc grpc.ClientConnInterface) ChatServiceClient {
	return &chatServiceClient{cc}
}

func (c *chatServiceClient) Create(ctx context.Context, in *CreateChatRequest, opts ...grpc.CallOption) (*Chat, error) {
	out := new(Chat)
	err := c.cc.Invoke(ctx, ChatService_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) Get(ctx context.Context, in *GetChatRequest, opts ...grpc.CallOption) (*Chat, error) {
	out := new(Chat)
	err := c.cc.Invoke(ctx, ChatService_Get_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) Update(ctx context.Context, in *UpdateChatRequest, opts ...grpc.CallOption) (*Chat, error) {
	out := new(Chat)
	err := c.cc.Invoke(ctx, ChatService_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) Delete(ctx context.Context, in *DeleteChatRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, ChatService_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) List(ctx context.Context, in *ListChatRequest, opts ...grpc.CallOption) (*ListChatResponse, error) {
	out := new(ListChatResponse)
	err := c.cc.Invoke(ctx, ChatService_List_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) BatchCreate(ctx context.Context, in *BatchCreateChatsRequest, opts ...grpc.CallOption) (*BatchCreateChatsResponse, error) {
	out := new(BatchCreateChatsResponse)
	err := c.cc.Invoke(ctx, ChatService_BatchCreate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChatServiceServer is the server API for ChatService service.
// All implementations must embed UnimplementedChatServiceServer
// for forward compatibility
type ChatServiceServer interface {
	Create(context.Context, *CreateChatRequest) (*Chat, error)
	Get(context.Context, *GetChatRequest) (*Chat, error)
	Update(context.Context, *UpdateChatRequest) (*Chat, error)
	Delete(context.Context, *DeleteChatRequest) (*empty.Empty, error)
	List(context.Context, *ListChatRequest) (*ListChatResponse, error)
	BatchCreate(context.Context, *BatchCreateChatsRequest) (*BatchCreateChatsResponse, error)
	mustEmbedUnimplementedChatServiceServer()
}

// UnimplementedChatServiceServer must be embedded to have forward compatible implementations.
type UnimplementedChatServiceServer struct {
}

func (UnimplementedChatServiceServer) Create(context.Context, *CreateChatRequest) (*Chat, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedChatServiceServer) Get(context.Context, *GetChatRequest) (*Chat, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedChatServiceServer) Update(context.Context, *UpdateChatRequest) (*Chat, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedChatServiceServer) Delete(context.Context, *DeleteChatRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedChatServiceServer) List(context.Context, *ListChatRequest) (*ListChatResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedChatServiceServer) BatchCreate(context.Context, *BatchCreateChatsRequest) (*BatchCreateChatsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BatchCreate not implemented")
}
func (UnimplementedChatServiceServer) mustEmbedUnimplementedChatServiceServer() {}

// UnsafeChatServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ChatServiceServer will
// result in compilation errors.
type UnsafeChatServiceServer interface {
	mustEmbedUnimplementedChatServiceServer()
}

func RegisterChatServiceServer(s grpc.ServiceRegistrar, srv ChatServiceServer) {
	s.RegisterService(&ChatService_ServiceDesc, srv)
}

func _ChatService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateChatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChatService_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).Create(ctx, req.(*CreateChatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetChatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChatService_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).Get(ctx, req.(*GetChatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateChatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChatService_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).Update(ctx, req.(*UpdateChatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteChatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChatService_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).Delete(ctx, req.(*DeleteChatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListChatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChatService_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).List(ctx, req.(*ListChatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_BatchCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BatchCreateChatsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).BatchCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChatService_BatchCreate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).BatchCreate(ctx, req.(*BatchCreateChatsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ChatService_ServiceDesc is the grpc.ServiceDesc for ChatService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ChatService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "entpb.ChatService",
	HandlerType: (*ChatServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _ChatService_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _ChatService_Get_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _ChatService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _ChatService_Delete_Handler,
		},
		{
			MethodName: "List",
			Handler:    _ChatService_List_Handler,
		},
		{
			MethodName: "BatchCreate",
			Handler:    _ChatService_BatchCreate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "entpb/entpb.proto",
}

const (
	UserService_Create_FullMethodName      = "/entpb.UserService/Create"
	UserService_Get_FullMethodName         = "/entpb.UserService/Get"
	UserService_Update_FullMethodName      = "/entpb.UserService/Update"
	UserService_Delete_FullMethodName      = "/entpb.UserService/Delete"
	UserService_List_FullMethodName        = "/entpb.UserService/List"
	UserService_BatchCreate_FullMethodName = "/entpb.UserService/BatchCreate"
)

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserServiceClient interface {
	Create(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*User, error)
	Get(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*User, error)
	Update(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*User, error)
	Delete(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	List(ctx context.Context, in *ListUserRequest, opts ...grpc.CallOption) (*ListUserResponse, error)
	BatchCreate(ctx context.Context, in *BatchCreateUsersRequest, opts ...grpc.CallOption) (*BatchCreateUsersResponse, error)
}

type userServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) Create(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, UserService_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Get(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, UserService_Get_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Update(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, UserService_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Delete(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, UserService_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) List(ctx context.Context, in *ListUserRequest, opts ...grpc.CallOption) (*ListUserResponse, error) {
	out := new(ListUserResponse)
	err := c.cc.Invoke(ctx, UserService_List_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) BatchCreate(ctx context.Context, in *BatchCreateUsersRequest, opts ...grpc.CallOption) (*BatchCreateUsersResponse, error) {
	out := new(BatchCreateUsersResponse)
	err := c.cc.Invoke(ctx, UserService_BatchCreate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
// All implementations must embed UnimplementedUserServiceServer
// for forward compatibility
type UserServiceServer interface {
	Create(context.Context, *CreateUserRequest) (*User, error)
	Get(context.Context, *GetUserRequest) (*User, error)
	Update(context.Context, *UpdateUserRequest) (*User, error)
	Delete(context.Context, *DeleteUserRequest) (*empty.Empty, error)
	List(context.Context, *ListUserRequest) (*ListUserResponse, error)
	BatchCreate(context.Context, *BatchCreateUsersRequest) (*BatchCreateUsersResponse, error)
	mustEmbedUnimplementedUserServiceServer()
}

// UnimplementedUserServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (UnimplementedUserServiceServer) Create(context.Context, *CreateUserRequest) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedUserServiceServer) Get(context.Context, *GetUserRequest) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedUserServiceServer) Update(context.Context, *UpdateUserRequest) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedUserServiceServer) Delete(context.Context, *DeleteUserRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedUserServiceServer) List(context.Context, *ListUserRequest) (*ListUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedUserServiceServer) BatchCreate(context.Context, *BatchCreateUsersRequest) (*BatchCreateUsersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BatchCreate not implemented")
}
func (UnimplementedUserServiceServer) mustEmbedUnimplementedUserServiceServer() {}

// UnsafeUserServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServiceServer will
// result in compilation errors.
type UnsafeUserServiceServer interface {
	mustEmbedUnimplementedUserServiceServer()
}

func RegisterUserServiceServer(s grpc.ServiceRegistrar, srv UserServiceServer) {
	s.RegisterService(&UserService_ServiceDesc, srv)
}

func _UserService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Create(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Get(ctx, req.(*GetUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Update(ctx, req.(*UpdateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Delete(ctx, req.(*DeleteUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).List(ctx, req.(*ListUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_BatchCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BatchCreateUsersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).BatchCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_BatchCreate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).BatchCreate(ctx, req.(*BatchCreateUsersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserService_ServiceDesc is the grpc.ServiceDesc for UserService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "entpb.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _UserService_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _UserService_Get_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _UserService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _UserService_Delete_Handler,
		},
		{
			MethodName: "List",
			Handler:    _UserService_List_Handler,
		},
		{
			MethodName: "BatchCreate",
			Handler:    _UserService_BatchCreate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "entpb/entpb.proto",
}

const (
	UserRelationsService_Create_FullMethodName      = "/entpb.UserRelationsService/Create"
	UserRelationsService_Get_FullMethodName         = "/entpb.UserRelationsService/Get"
	UserRelationsService_Update_FullMethodName      = "/entpb.UserRelationsService/Update"
	UserRelationsService_Delete_FullMethodName      = "/entpb.UserRelationsService/Delete"
	UserRelationsService_List_FullMethodName        = "/entpb.UserRelationsService/List"
	UserRelationsService_BatchCreate_FullMethodName = "/entpb.UserRelationsService/BatchCreate"
)

// UserRelationsServiceClient is the client API for UserRelationsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserRelationsServiceClient interface {
	Create(ctx context.Context, in *CreateUserRelationsRequest, opts ...grpc.CallOption) (*UserRelations, error)
	Get(ctx context.Context, in *GetUserRelationsRequest, opts ...grpc.CallOption) (*UserRelations, error)
	Update(ctx context.Context, in *UpdateUserRelationsRequest, opts ...grpc.CallOption) (*UserRelations, error)
	Delete(ctx context.Context, in *DeleteUserRelationsRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	List(ctx context.Context, in *ListUserRelationsRequest, opts ...grpc.CallOption) (*ListUserRelationsResponse, error)
	BatchCreate(ctx context.Context, in *BatchCreateUserRelationsSliceRequest, opts ...grpc.CallOption) (*BatchCreateUserRelationsSliceResponse, error)
}

type userRelationsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserRelationsServiceClient(cc grpc.ClientConnInterface) UserRelationsServiceClient {
	return &userRelationsServiceClient{cc}
}

func (c *userRelationsServiceClient) Create(ctx context.Context, in *CreateUserRelationsRequest, opts ...grpc.CallOption) (*UserRelations, error) {
	out := new(UserRelations)
	err := c.cc.Invoke(ctx, UserRelationsService_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userRelationsServiceClient) Get(ctx context.Context, in *GetUserRelationsRequest, opts ...grpc.CallOption) (*UserRelations, error) {
	out := new(UserRelations)
	err := c.cc.Invoke(ctx, UserRelationsService_Get_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userRelationsServiceClient) Update(ctx context.Context, in *UpdateUserRelationsRequest, opts ...grpc.CallOption) (*UserRelations, error) {
	out := new(UserRelations)
	err := c.cc.Invoke(ctx, UserRelationsService_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userRelationsServiceClient) Delete(ctx context.Context, in *DeleteUserRelationsRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, UserRelationsService_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userRelationsServiceClient) List(ctx context.Context, in *ListUserRelationsRequest, opts ...grpc.CallOption) (*ListUserRelationsResponse, error) {
	out := new(ListUserRelationsResponse)
	err := c.cc.Invoke(ctx, UserRelationsService_List_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userRelationsServiceClient) BatchCreate(ctx context.Context, in *BatchCreateUserRelationsSliceRequest, opts ...grpc.CallOption) (*BatchCreateUserRelationsSliceResponse, error) {
	out := new(BatchCreateUserRelationsSliceResponse)
	err := c.cc.Invoke(ctx, UserRelationsService_BatchCreate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserRelationsServiceServer is the server API for UserRelationsService service.
// All implementations must embed UnimplementedUserRelationsServiceServer
// for forward compatibility
type UserRelationsServiceServer interface {
	Create(context.Context, *CreateUserRelationsRequest) (*UserRelations, error)
	Get(context.Context, *GetUserRelationsRequest) (*UserRelations, error)
	Update(context.Context, *UpdateUserRelationsRequest) (*UserRelations, error)
	Delete(context.Context, *DeleteUserRelationsRequest) (*empty.Empty, error)
	List(context.Context, *ListUserRelationsRequest) (*ListUserRelationsResponse, error)
	BatchCreate(context.Context, *BatchCreateUserRelationsSliceRequest) (*BatchCreateUserRelationsSliceResponse, error)
	mustEmbedUnimplementedUserRelationsServiceServer()
}

// UnimplementedUserRelationsServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUserRelationsServiceServer struct {
}

func (UnimplementedUserRelationsServiceServer) Create(context.Context, *CreateUserRelationsRequest) (*UserRelations, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedUserRelationsServiceServer) Get(context.Context, *GetUserRelationsRequest) (*UserRelations, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedUserRelationsServiceServer) Update(context.Context, *UpdateUserRelationsRequest) (*UserRelations, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedUserRelationsServiceServer) Delete(context.Context, *DeleteUserRelationsRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedUserRelationsServiceServer) List(context.Context, *ListUserRelationsRequest) (*ListUserRelationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedUserRelationsServiceServer) BatchCreate(context.Context, *BatchCreateUserRelationsSliceRequest) (*BatchCreateUserRelationsSliceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BatchCreate not implemented")
}
func (UnimplementedUserRelationsServiceServer) mustEmbedUnimplementedUserRelationsServiceServer() {}

// UnsafeUserRelationsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserRelationsServiceServer will
// result in compilation errors.
type UnsafeUserRelationsServiceServer interface {
	mustEmbedUnimplementedUserRelationsServiceServer()
}

func RegisterUserRelationsServiceServer(s grpc.ServiceRegistrar, srv UserRelationsServiceServer) {
	s.RegisterService(&UserRelationsService_ServiceDesc, srv)
}

func _UserRelationsService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRelationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserRelationsServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserRelationsService_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserRelationsServiceServer).Create(ctx, req.(*CreateUserRelationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserRelationsService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserRelationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserRelationsServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserRelationsService_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserRelationsServiceServer).Get(ctx, req.(*GetUserRelationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserRelationsService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserRelationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserRelationsServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserRelationsService_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserRelationsServiceServer).Update(ctx, req.(*UpdateUserRelationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserRelationsService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteUserRelationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserRelationsServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserRelationsService_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserRelationsServiceServer).Delete(ctx, req.(*DeleteUserRelationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserRelationsService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListUserRelationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserRelationsServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserRelationsService_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserRelationsServiceServer).List(ctx, req.(*ListUserRelationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserRelationsService_BatchCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BatchCreateUserRelationsSliceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserRelationsServiceServer).BatchCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserRelationsService_BatchCreate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserRelationsServiceServer).BatchCreate(ctx, req.(*BatchCreateUserRelationsSliceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserRelationsService_ServiceDesc is the grpc.ServiceDesc for UserRelationsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserRelationsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "entpb.UserRelationsService",
	HandlerType: (*UserRelationsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _UserRelationsService_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _UserRelationsService_Get_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _UserRelationsService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _UserRelationsService_Delete_Handler,
		},
		{
			MethodName: "List",
			Handler:    _UserRelationsService_List_Handler,
		},
		{
			MethodName: "BatchCreate",
			Handler:    _UserRelationsService_BatchCreate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "entpb/entpb.proto",
}
