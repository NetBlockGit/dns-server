package auth

import (
	"context"

	"google.golang.org/grpc"
)

func CheckAuth(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	return handler(ctx, req)
}
