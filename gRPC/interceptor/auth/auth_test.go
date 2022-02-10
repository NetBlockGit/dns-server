package auth

import (
	"context"
	"dnsserver/config/authtoken"
	envconfig "dnsserver/config/env"
	"dnsserver/util/env"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func TestCheckAuth(t *testing.T) {
	envconfig.Init("../../../.env")
	grpcInfo := &grpc.UnaryServerInfo{
		FullMethod: "/BlockerService/Protected",
	}
	t.Run("should continue running next handlers if token is valid", func(t *testing.T) {
		privateKeyHex := env.Get("PASETO_PRIVATE_KEY")
		wAddr := env.Get("WALLET_ADDRESS_OF_AUTHORIZED")
		cxt := createContext(privateKeyHex, wAddr)
		resp, err := CheckAuth(cxt, nil, grpcInfo, handler)
		if err != nil {
			t.Fatalf("got error response from CheckAuth, error: %v", err)
		}
		assert.Equal(t, true, resp)
	})

	t.Run("should return error if token is invalid", func(t *testing.T) {
		privateKeyHex := "68656c6c6f20686f772061726520796f75"
		wAddr := env.Get("WALLET_ADDRESS_OF_AUTHORIZED")
		cxt := createContext(privateKeyHex, wAddr)
		resp, err := CheckAuth(cxt, nil, grpcInfo, handler)
		assert.Equal(t, nil, resp)
		assert.Contains(t, err.Error(), "failed to verify token")
	})

	t.Run("should return error if user is not authorized in block chain", func(t *testing.T) {
		privateKeyHex := env.Get("PASETO_PRIVATE_KEY")
		wAddr := "0x23F680e406CFda323Ef23d0594486811eE6b8587"
		cxt := createContext(privateKeyHex, wAddr)
		resp, err := CheckAuth(cxt, nil, grpcInfo, handler)
		assert.Equal(t, nil, resp)
		assert.Contains(t, err.Error(), "failed to verify token")
	})

}

func createContext(privateKeyHex string, walletAddr string) context.Context {
	token, err := authtoken.GenerateToken(privateKeyHex, walletAddr)
	if err != nil {
		log.Fatalf("failed to generate token, error: %v", err.Error())
	}
	return metadata.NewIncomingContext(context.Background(), metadata.Pairs(
		"authorization", token,
	))
}

func handler(ctx context.Context, req interface{}) (success interface{}, err error) {
	return true, nil
}
