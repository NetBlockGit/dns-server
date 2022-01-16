package authtoken

import (
	"dnsserver/config/smartcontract"
	"dnsserver/util/env"
	"encoding/hex"

	"github.com/ethereum/go-ethereum/common"
	"github.com/vk-rv/pvx"
)

func VerifyToken(token string) (bool, error) {
	var claims DnsBlockerClaims
	privateKeyHex := env.Get("PASETO_PRIVATE_KEY")
	privKey, err := hex.DecodeString(privateKeyHex)

	if err != nil {
		return false, err
	}
	symK := pvx.NewSymmetricKey(privKey, pvx.Version4)
	pv4 := pvx.NewPV4Local()
	err = pv4.Decrypt(token, symK).ScanClaims(&claims)
	instance := smartcontract.GetBlockListInstance(smartcontract.GetClient())
	addr := common.HexToAddress(claims.WalletAddress)
	isAuthorized, err := instance.AuthorizedUsers(nil, addr)
	if err != nil {
		return false, err
	}
	if !isAuthorized {
		return false, err
	}

	return true, nil
}
