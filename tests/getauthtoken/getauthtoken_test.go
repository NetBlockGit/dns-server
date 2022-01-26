package getauthtoken

import (
	"context"
	"crypto/ecdsa"
	"dnsserver/config/dnsblocker"
	"dnsserver/config/env"
	"dnsserver/config/mongodb"
	"dnsserver/gRPC/blockerserver"
	"dnsserver/generated/protos/getauthtoken"
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"go.mongodb.org/mongo-driver/bson"
)

func Test_GetAuthToken(t *testing.T) {
	env.Init()
	dnsblocker.CheckInitAndGet()
	time.Sleep(2 * time.Second)
	s := blockerserver.Server{}
	mongodb.StatsCollection.DeleteMany(context.Background(), bson.D{})
	testWallet := GenerateWallet()
	signature := getSignature(time.Now().Format("2006-Jan-02"), testWallet.PrivateKey)
	res, err := s.GetAuthToken(context.Background(), &getauthtoken.GetAuthTokenRequest{Signature: signature, WalletAddress: testWallet.WalletAddress})
	if err != nil {
		t.Fatal(err)
	}
	if len(res.Token) < 1 {
		t.Fatal("Token length is less than 1")
	}
}

func getSignature(message string, privateKey string) string {
	newMsg := fmt.Sprintf("\x19Ethereum Signed Message:\n%v%v", len(message), message)

	privateKeyECDSA, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		log.Fatal("HexToECDSA failed ", err)
	}

	// keccak256 hash of the data
	dataBytes := []byte(newMsg)
	hashData := crypto.Keccak256Hash(dataBytes)

	signatureBytes, err := crypto.Sign(hashData.Bytes(), privateKeyECDSA)
	if err != nil {
		log.Fatal("len", err)
	}

	signature := hexutil.Encode(signatureBytes)

	return signature
}

type TestWallet struct {
	PrivateKey    string
	WalletAddress string
}

func GenerateWallet() *TestWallet {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}

	privateKeyBytes := crypto.FromECDSA(privateKey)

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	privateKeyHex := hexutil.Encode(privateKeyBytes)
	testWallet := TestWallet{
		PrivateKey:    privateKeyHex[2:],
		WalletAddress: address,
	}
	return &testWallet
}
