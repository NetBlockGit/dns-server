package authtoken

import (
	"github.com/vk-rv/pvx"
)

type DnsBlockerClaims struct {
	pvx.RegisteredClaims
	WalletAddress string
}

func (c *DnsBlockerClaims) Valid() error {

	if err := c.RegisteredClaims.Valid(); err != nil {
		return err
	}

	return nil

}
