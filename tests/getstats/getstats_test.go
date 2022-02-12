package getstats

import (
	"context"
	"dnsserver/config/dnsblocker"
	"dnsserver/config/env"
	"dnsserver/config/mongodb"
	"dnsserver/gRPC/blockerserver"
	"dnsserver/generated/protos/getstats"
	"testing"
	"time"

	"github.com/NetBlockGit/dnsblocker/config"
	"github.com/miekg/dns"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
)

func Test_GetStats(t *testing.T) {
	env.Init("../../.env")
	blockerConfig := dnsblocker.CheckInitAndGet()
	time.Sleep(2 * time.Second)
	s := blockerserver.Server{}
	mongodb.StatsCollection.DeleteMany(context.Background(), bson.D{})
	blockerConfig.AddHostToBlockList("toast.com")

	queries := []config.QueryEvent{
		{Hostname: "toast.com", Blocked: true},
		{Hostname: "ommore.me", Blocked: false},
		{Hostname: "toast.xyz", Blocked: false},
	}
	for _, q := range queries {
		tryAQuery(t, q.Hostname, blockerConfig.Addr)
	}

	time.Sleep(2 * time.Second)
	res, err := s.GetStats(context.Background(), &getstats.GetStatsRequest{})
	if err != nil {
		t.Fatal(err)
	}

	for i, s := range res.Stats {
		assert.Equal(t, queries[i].Hostname, s.Hostname)
		assert.Equal(t, queries[i].Blocked, s.Blocked)
	}
}

func tryAQuery(t *testing.T, host string, addr string) *dns.Msg {
	msg := new(dns.Msg)
	msg.SetQuestion(host+".", dns.TypeA)
	r, e := dns.Exchange(msg, addr)
	if e != nil {
		t.Fatal(e)
	}
	return r
}
