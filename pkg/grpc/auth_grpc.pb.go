// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: auth.proto

package grpc

import (
	"Real-Time-Chat/ent/proto/entpb"
	context "context"
	"crypto/rand"
	"encoding/base64"
	"github.com/go-redis/redis/v8"
	empty "github.com/golang/protobuf/ptypes/empty"
	"golang.org/x/crypto/bcrypt"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	status "google.golang.org/grpc/status"
	"io"
	"log"
	"time"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	AuthService_Signup_FullMethodName = "/chat.AuthService/Signup"
	AuthService_Login_FullMethodName  = "/chat.AuthService/Login"
	AuthService_Logout_FullMethodName = "/chat.AuthService/Logout"
)

// AuthServiceClient is the client API for AuthService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthServiceClient interface {
	Signup(ctx context.Context, in *SignupRequest, opts ...grpc.CallOption) (*SignupResponse, error)
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
	Logout(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*LogoutResponse, error)
}

type authServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthServiceClient(cc grpc.ClientConnInterface) AuthServiceClient {
	return &authServiceClient{cc}
}

func (c *authServiceClient) Signup(ctx context.Context, in *SignupRequest, opts ...grpc.CallOption) (*SignupResponse, error) {
	out := new(SignupResponse)
	err := c.cc.Invoke(ctx, AuthService_Signup_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, AuthService_Login_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) Logout(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*LogoutResponse, error) {
	out := new(LogoutResponse)
	err := c.cc.Invoke(ctx, AuthService_Logout_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthServiceServer is the server API for AuthService service.
// All implementations must embed UnimplementedAuthServiceServer
// for forward compatibility
type AuthServiceServer interface {
	Signup(context.Context, *SignupRequest) (*SignupResponse, error)
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
	Logout(context.Context, *empty.Empty) (*LogoutResponse, error)
	mustEmbedUnimplementedAuthServiceServer()
}

var (
	redisConn *redis.Client
	grpConn *grpc.ClientConn
	client entpb.UserServiceClient
)

func init() {
	// Redisとのコネクションを開く
	redisConn = redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})

	// gRPCサーバーとのコネクションを開く
	var err error
	grpConn, err = grpc.Dial(":8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed connecting to server: %s", err)
	}

	// Create a User service Client on the connection.
	client = entpb.NewUserServiceClient(grpConn)
}

type AuthServer struct {
	UnimplementedAuthServiceServer
}

// UnimplementedAuthServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAuthServiceServer struct {
}

func (UnimplementedAuthServiceServer) Signup(ctx context.Context, req *SignupRequest) (*SignupResponse, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "パスワードのハッシュ化に失敗しました: %v", err)
	}
	_, err = client.Create(ctx, &entpb.CreateUserRequest{
		User: &entpb.User{
			Username: req.Username,
			PasswordHash: string(hash),
		},
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "ユーザー作成時にエラーが発生しました: %v", err)
	}
	return &SignupResponse{Success: true}, nil
}
func (UnimplementedAuthServiceServer) Login(ctx context.Context, req *LoginRequest) (*LoginResponse, error) {
	b := make([]byte, 64)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return &LoginResponse{Success: false}, status.Errorf(codes.Internal, "ランダムな文字作成時にエラーが発生しました: %v", err)
	}
	newRedisKey := base64.URLEncoding.EncodeToString(b)
	redisValue := "user_id" // 実際にはリクエストから得られるユーザーIDやその他のユーザー情報を使う

	// Redisにセッションキーを保存（例: 24時間の有効期限を設定）
	if err := redisConn.Set(ctx, newRedisKey, redisValue, 24*time.Hour).Err(); err != nil {
		return &LoginResponse{Success: false}, status.Errorf(codes.Internal, "Session登録時にエラーが発生: %v", err)
	}

	md := metadata.New(map[string]string{"session-id": newRedisKey})
	if err := grpc.SendHeader(ctx, md); err != nil {
		return &LoginResponse{Success: false}, status.Errorf(codes.Internal, "SendHeader時にエラーが発生: %v", err)
	}

	return &LoginResponse{Success: true}, nil
}
func (UnimplementedAuthServiceServer) Logout(ctx context.Context, req *empty.Empty) (*LogoutResponse, error) {
	// メタデータからセッションIDを取得
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Errorf(codes.Internal, "メタデータの取得に失敗しました")
	}
	sessionIDs, ok := md["session-id"]
	if !ok || len(sessionIDs) == 0 {
		return nil, status.Errorf(codes.Unauthenticated, "セッションIDが見つかりません")
	}
	sessionID := sessionIDs[0]

	// Redisからセッションキーを削除
	if err := redisConn.Del(ctx, sessionID).Err(); err != nil {
		return nil, status.Errorf(codes.Internal, "セッション削除時にエラーが発生: %v", err)
	}

	return &LogoutResponse{Success: true}, nil
}
func (UnimplementedAuthServiceServer) mustEmbedUnimplementedAuthServiceServer() {}

// UnsafeAuthServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthServiceServer will
// result in compilation errors.
type UnsafeAuthServiceServer interface {
	mustEmbedUnimplementedAuthServiceServer()
}

func RegisterAuthServiceServer(s grpc.ServiceRegistrar, srv AuthServiceServer) {
	s.RegisterService(&AuthService_ServiceDesc, srv)
}

func _AuthService_Signup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).Signup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_Signup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).Signup(ctx, req.(*SignupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_Login_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_Logout_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).Logout(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_Logout_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).Logout(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// AuthService_ServiceDesc is the grpc.ServiceDesc for AuthService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AuthService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "chat.AuthService",
	HandlerType: (*AuthServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Signup",
			Handler:    _AuthService_Signup_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _AuthService_Login_Handler,
		},
		{
			MethodName: "Logout",
			Handler:    _AuthService_Logout_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "auth.proto",
}
