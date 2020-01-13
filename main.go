package main

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"grpc_test/middleware"
	"grpc_test/service"
	"grpc_test/pb"
	"log"
	"net"
)

func main() {
	listenPort, err := net.Listen("tcp", ":19003")
	if err != nil {
		log.Fatalln(err)
	}
	opt := []grpc.ServerOption{grpc.UnaryInterceptor(middleware.UnaryServerInterCeptor(log.Logger{}))}
	server := grpc.NewServer(opt...)
	catService := &service.MyCatService{}
	pb.RegisterCatServer(server, catService)
	reflection.Register(server) //grprurlを簡易に使えるように。
	server.Serve(listenPort)
}

// memo.
//　稼働確認コマンド
//grpcurl -plaintext -d '{"target_cat": "tama"}' localhost:19003 Cat.GetMyCat

