package dnsblocker

import (
	"context"
	"dnsserver/config/mongodb"
	"dnsserver/util/env"
	"log"

	"github.com/NetBlockGit/dnsblocker/config"
)

var dnsBlocker *config.BlockerConfig

// Returns blocker instance and init it if it doesn't exist
func CheckInitAndGet() *config.BlockerConfig {
	if dnsBlocker != nil {
		return dnsBlocker
	}
	var qCh chan config.QueryEvent
	qCh = make(chan config.QueryEvent)
	go saveStats(&qCh)

	dnsBlocker = &config.BlockerConfig{
		UpstreamDns:  "1.1.1.1:53",
		BlockList:    []string{},
		Addr:         env.Get("DNS_SERVER_URL"),
		Enabled:      true,
		QueryChannel: qCh,
	}
	go dnsBlocker.StartDnsServer()
	return dnsBlocker
}

var firstQueryHit = false

func saveStats(ch *chan config.QueryEvent) {
	for ev := range *ch {
		go addHostnameToLogDb(ev)
	}
}

func addHostnameToLogDb(ev config.QueryEvent) {
	if !firstQueryHit {
		mongodb.Disconnect()
		mongodb.Init()
		firstQueryHit = true
	}
	col := mongodb.StatsCollection
	_, err := col.InsertOne(context.Background(), ev)
	if err != nil {
		log.Printf("failed to insert, error %v", err.Error())
	}
}
