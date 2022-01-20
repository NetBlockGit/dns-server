package dnsblocker

import (
	"context"
	"dnsserver/config/mongodb"
	"log"

	"github.com/NetBlockGit/dnsblocker/config"
)

var dnsBlocker *config.BlockerConfig

// Returns blocker instance and init it if it doesn't exist
func CheckInitAndGet() *config.BlockerConfig {
	if dnsBlocker != nil {
		return dnsBlocker
	}
	qCh := make(chan config.QueryEvent)
	dnsBlocker = &config.BlockerConfig{
		UpstreamDns:  "1.1.1.1:53",
		BlockList:    []string{},
		Addr:         "127.0.0.1:5302",
		Enabled:      true,
		QueryChannel: qCh,
	}
	go saveStats(&qCh)
	go dnsBlocker.StartDnsServer()
	return dnsBlocker
}

func saveStats(ch *chan config.QueryEvent) {
	for ev := range *ch {
		col := mongodb.StatsCollection
		_, err := col.InsertOne(context.Background(), ev)
		if err != nil {
			log.Printf("failed to insert, error %v", err.Error())
		}
	}
}
