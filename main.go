package main

import (
	"google.golang.org/grpc"
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
	server.Serve(listenPort)
}



