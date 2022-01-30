package grpcserver

import (
	"dnsserver/gRPC/blockerserver"
	"dnsserver/gRPC/interceptor/auth"
	"dnsserver/generated/protos"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func Init() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", 8000))
	if err != nil {
		log.Fatalf("failed to listen on tcp port %v", 8000)
	}
	grpcBlockerServer := grpc.NewServer(grpc.UnaryInterceptor(auth.CheckAuth))
	reflection.Register(grpcBlockerServer)
	protos.RegisterBlockerServer(grpcBlockerServer, blockerserver.Server{})
	log.Printf("starting to listen on port %v", 8000)
	if err := grpcBlockerServer.Serve(lis); err != nil {
		log.Fatalf("Cannot serve: %v", err)
	}
}
