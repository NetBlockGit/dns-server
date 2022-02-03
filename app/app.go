package app

import (
	"dnsserver/app/routine/syncblocklist"
	"dnsserver/config/dnsblocker"
	"dnsserver/config/env"
	"dnsserver/gRPC/grpcserver"
)

func Init() {
	env.Init()
	dnsblocker.CheckInitAndGet()
	go syncblocklist.Init()
	grpcserver.Init()
}
