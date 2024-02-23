package main

import (
	"Real-Time-Chat/ent"
	"Real-Time-Chat/ent/proto/entpb"
	"Real-Time-Chat/model"
	"context"
	"fmt"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"os"
	"os/signal"

	mygrpc "Real-Time-Chat/pkg/grpc"
)

func main() {
	DB := model.DBOpen()
	defer model.DBClose(DB)

	// entクライアントの初期化
	client := model.EntOpen()
	defer model.EntClose(client)

	// マイグレーションツールの実行（テーブルの作成など）
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	port := 8080
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		panic(err)
	}

	// gRPCサーバーを作成
	server := grpc.NewServer()

	// サーバーにgrpc自作メソッドを登録
	registerGrpcServices(server)

	// サーバーにentメソッドを登録
	registerEntServices(server, client)

	go func() {
		log.Printf("start gRPC server port: %v", port)
		server.Serve(listener)
	}()

	reflection.Register(server)
	
	// Ctrl+Cが入力されたらGraceful shutdownされるようにする
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("stopping gRPC server...")
	server.GracefulStop()
}

func registerGrpcServices(server *grpc.Server) {
	mygrpc.RegisterAuthServiceServer(server, &mygrpc.AuthServer{})
}

func registerEntServices(server *grpc.Server, client *ent.Client) {
	entpb.RegisterChatServiceServer(server, entpb.NewChatService(client))
	entpb.RegisterUserServiceServer(server, entpb.NewUserService(client))
	entpb.RegisterUserRelationServiceServer(server, entpb.NewUserRelationService(client))
}
