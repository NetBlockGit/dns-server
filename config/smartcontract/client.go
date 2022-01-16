package smartcontract

import (
	"dnsserver/util/env"
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
)

var client *ethclient.Client

func GetClient() *ethclient.Client {
	if client != nil {
		return client
	}

	rpcUrl := env.Get("RPC_URL")
	client, err := ethclient.Dial(rpcUrl)
	if err != nil {
		log.Fatalf("failed to dial rpc url %v", rpcUrl)
	}
	return client
}
