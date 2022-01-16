package blockerserver

import (
	"context"
	"dnsserver/config/dnsblocker"
	"dnsserver/protos"
)

type Server struct {
	protos.UnimplementedBlockerServer
}

func (s Server) ToggleBlocker(c context.Context, request *protos.ToggleBlockerRequest) (*protos.ToggleBlockerResponse, error) {
	dnsConfig := dnsblocker.CheckInitAndGet()
	dnsConfig.ToggleBlocker()
	return &protos.ToggleBlockerResponse{}, nil
}
