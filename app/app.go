package app

import (
	"dnsserver/config/dnsblocker"
	"dnsserver/config/env"
	"dnsserver/gRPC/grpcserver"
)

func Init() {
	env.Init(".env")
	dnsblocker.CheckInitAndGet()
	grpcserver.Init()
}
