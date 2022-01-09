package main

import (
	"dnsserver/gRPC/interceptor/auth"

	dnsblockerpackage "github.com/NetBlockGit/dnsblocker"
	"google.golang.org/grpc"
)

func main() {
	dnsblockerpackage.StartDnsServer()
	// lis, err := net.Listen("tcp", fmt.Sprintf(":%v", 8000))
	// grpcServer := grpc.NewServer()
	grpc.UnaryInterceptor(auth.CheckAuth)
}
