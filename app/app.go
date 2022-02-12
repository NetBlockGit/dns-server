package app

import (
	"dnsserver/app/routine/syncblocklist"
	"dnsserver/config/dnsblocker"
	"dnsserver/config/env"
	"dnsserver/config/mongodb"
	"dnsserver/gRPC/grpcserver"
)

func Init() {
	env.Init(".env")
	mongodb.Init()
	dnsblocker.CheckInitAndGet()
	syncblocklist.Init()
	grpcserver.Init()
}
