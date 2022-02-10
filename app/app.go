package app

import (
	"dnsserver/config/dnsblocker"
	"dnsserver/config/env"
	"dnsserver/gRPC/grpcserver"
)

func Init() {
	env.Init("")
	dnsblocker.CheckInitAndGet()
	grpcserver.Init()
}
