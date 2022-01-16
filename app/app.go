package app

import (
	"dnsserver/config/dnsblocker"
	"dnsserver/gRPC/interceptor/auth"
	"dnsserver/gRPC/server"
	"dnsserver/protos"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

func Init() {
	dnsblocker.CheckInitAndGet()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", 8000))
	if err != nil {
		log.Fatalf("failed to listen on tcp port %v", 8000)
	}
	grpcServer := grpc.NewServer()
	grpc.UnaryInterceptor(auth.CheckAuth)
	protos.RegisterBlockerServer(grpcServer, server.Server{})
	log.Printf("starting to listen on port %v", 8000)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Cannot serve: %v", err)
	}
}
