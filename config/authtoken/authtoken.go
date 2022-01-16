package authtoken

import (
	"encoding/hex"
	"time"

	"github.com/vk-rv/pvx"
)

func GenerateToken(privateKeyHex string,walletAddress string) (string, error) {
	privKey, err := hex.DecodeString(privateKeyHex)

	if err != nil {
		return "", err
	}
	symK := pvx.NewSymmetricKey(privKey, pvx.Version4)
	pv4 := pvx.NewPV4Local()

	expiration := pvx.TimePtr(time.Now().Add(24 * time.Hour))
	claims := DnsBlockerClaims{
		pvx.RegisteredClaims{Expiration: expiration},
		walletAddress,
	}
	token, err := pv4.Encrypt(symK, &claims)

	if err != nil {
		return "", err
	}
	return token, nil
}
