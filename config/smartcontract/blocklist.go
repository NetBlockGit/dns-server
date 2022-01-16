package smartcontract

import (
	"dnsserver/generated/smartcontract/blocklist"
	"dnsserver/util/env"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var instance *blocklist.Blocklist

func GetBlockListInstance(client *ethclient.Client) *blocklist.Blocklist {
	if instance != nil {
		return instance
	}
	blocklistAddressHex := env.Get("SMART_CONTRACT_ADDRESS")
	addr := common.HexToAddress(blocklistAddressHex)
	var err error
	instance, err = blocklist.NewBlocklist(addr, client)
	if err != nil {
		log.Fatalf("failed to create blocklist instance at address %v, error: %v", blocklistAddressHex, err.Error())
	}
	return instance
}
