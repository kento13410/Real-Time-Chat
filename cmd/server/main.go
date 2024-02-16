package main

import (
	"Real-Time-Chat/ent/proto/entpb"
	"Real-Time-Chat/model"
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

	port := 8080
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		panic(err)
	}

	// gRPCサーバーを作成
	server := grpc.NewServer()

	// サーバー登録用の構造体を呼び出す
	// ほんとは前方互換性のため埋め込みをする必要がある
	s := mygrpc.UnimplementedAuthServiceServer{}

	// サーバーにAuthメソッドを登録
	mygrpc.RegisterAuthServiceServer(server, s)

	svc := entpb.NewChatService(client)
	entpb.RegisterChatServiceServer(server, svc)

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
