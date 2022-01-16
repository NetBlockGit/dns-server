package auth

import (
	"context"
	"dnsserver/config/authtoken"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

var (
	ErrVerificationFailed = status.Error(codes.Unauthenticated, "failed to verify token")
)

func CheckAuth(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		log.Println("failed to get metadata from incoming context")
		return nil, ErrVerificationFailed
	}
	var token string
	if authmd := md["authorization"]; len(authmd) == 1 {
		token = md["authorization"][0]
	} else {
		return nil, ErrVerificationFailed
	}
	valid, err := authtoken.VerifyToken(token)
	if err != nil {
		return nil, ErrVerificationFailed
	}
	if valid {
		return handler(ctx, req)
	} else {
		return nil, ErrVerificationFailed
	}
}
