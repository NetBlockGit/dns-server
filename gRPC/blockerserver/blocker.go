package blockerserver

import (
	"context"
	"dnsserver/config/authtoken"
	"dnsserver/config/dnsblocker"
	"dnsserver/config/mongodb"
	"dnsserver/config/smartcontract"
	"dnsserver/generated/protos"
	"dnsserver/generated/protos/getauthtoken"
	"dnsserver/generated/protos/getstats"
	"dnsserver/generated/protos/toggleblocker"
	"dnsserver/generated/protos/updateupstreamdns"
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/vk-rv/pvx"
	"go.mongodb.org/mongo-driver/bson"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
	protos.UnimplementedBlockerServer
}

var (
	ErrUnexpected           = status.Error(codes.Internal, "unexpected error occured")
	ErrInvalidWalletAddress = status.Error(codes.PermissionDenied, "wallet address is invalid")
	ErrNotAuthorized        = status.Error(codes.PermissionDenied, "wallet address is not authorized")
)

func (s Server) ToggleBlocker(c context.Context, request *toggleblocker.ToggleBlockerRequest) (*toggleblocker.ToggleBlockerResponse, error) {
	dnsConfig := dnsblocker.CheckInitAndGet()
	dnsConfig.ToggleBlocker()
	return &toggleblocker.ToggleBlockerResponse{}, nil
}

func (s Server) GetStats(c context.Context, request *getstats.GetStatsRequest) (*getstats.GetStatsResponse, error) {
	col := mongodb.StatsCollection
	res, err := col.Find(context.Background(), bson.D{})
	if err != nil {
		log.Printf("failed to find objects, error: %v", err.Error())
		return nil, ErrUnexpected
	}
	var stats []*getstats.Stats
	res.All(context.Background(), &stats)
	return &getstats.GetStatsResponse{
		Stats: stats,
	}, nil
}

func (s Server) UpdateUpstreamDns(c context.Context, request *updateupstreamdns.UpdateUpstreamDnsRequest) (*updateupstreamdns.UpdateUpstreamDnsResponse, error) {
	dnsConfig := dnsblocker.CheckInitAndGet()
	dnsConfig.UpstreamDns = request.NewDnsServer
	return &updateupstreamdns.UpdateUpstreamDnsResponse{}, nil
}
func (s Server) GetAuthToken(c context.Context, request *getauthtoken.GetAuthTokenRequest) (*getauthtoken.GetAuthTokenResponse, error) {
	message := time.Now().Format("2006-Jan-02")
	newMsg := fmt.Sprintf("\x19Ethereum Signed Message:\n%v%v", len(message), message)
	newMsgHash := crypto.Keccak256Hash([]byte(newMsg))
	signatureInBytes, err := hexutil.Decode(request.Signature)
	if err != nil {
		log.Printf("failed to decode, error: %v", err.Error())
		return nil, status.Error(codes.Internal, "unexpected error occured")
	}
	if signatureInBytes[64] == 27 || signatureInBytes[64] == 28 {
		signatureInBytes[64] -= 27
	}
	pubKey, err := crypto.SigToPub(newMsgHash.Bytes(), signatureInBytes)

	if err != nil {
		log.Printf("failed to decode, error: %v", err.Error())
		return nil, status.Error(codes.Internal, "unexpected error occured")
	}

	//Get address from public key
	walletAddress := crypto.PubkeyToAddress(*pubKey)

	if walletAddress.String() != request.WalletAddress {
		return nil, ErrInvalidWalletAddress
	}

	blocklistInstance := smartcontract.GetBlockListInstance(smartcontract.GetClient())
	isUserAuthorized, err := blocklistInstance.AuthorizedUsers(nil, walletAddress)
	if err != nil {
		return nil, ErrUnexpected
	}
	if !isUserAuthorized {
		return nil, ErrNotAuthorized
	}
	pasetoToken, err := generateToken(walletAddress.String())
	if err != nil {
		return nil, status.Error(codes.Internal, "unexpected error occured")
	}
	return &getauthtoken.GetAuthTokenResponse{
		Token: pasetoToken,
	}, nil
}

func generateToken(walletAddress string) (string, error) {

	privKey, err := hex.DecodeString(os.Getenv("PASETO_PRIVATE_KEY"))

	if err != nil {
		return "", err
	}
	symK := pvx.NewSymmetricKey(privKey, pvx.Version4)
	pv4 := pvx.NewPV4Local()
	claims := authtoken.DnsBlockerClaims{
		RegisteredClaims: pvx.RegisteredClaims{Expiration: pvx.TimePtr(time.Now().Add(24 * time.Hour))},
		WalletAddress:    walletAddress,
	}
	token, err := pv4.Encrypt(symK, &claims)

	if err != nil {
		return "", err
	}
	return token, nil
}
