package blockerserver

import (
	"context"
	"dnsserver/config/dnsblocker"
	"dnsserver/config/mongodb"
	"dnsserver/generated/protos"
	"dnsserver/generated/protos/toggleblocker"
	"dnsserver/generated/protos/getstats"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
	protos.UnimplementedBlockerServer
}

var (
	ErrFailedToGetStats = status.Error(codes.Internal, "unexpected error occured")
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
		return nil, ErrFailedToGetStats
	}
	var stats []*getstats.Stats
	res.All(context.Background(), &stats)
	return &getstats.GetStatsResponse{
		Stats: stats,
	}, nil
}
