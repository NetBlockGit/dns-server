package syncblocklist

import (
	"dnsserver/config/dnsblocker"
	"dnsserver/config/smartcontract"
	"dnsserver/generated/smartcontract/blocklist"
	"log"
)

func Init() {
	blockerConfig := dnsblocker.CheckInitAndGet()
	client := smartcontract.GetClient()
	bl := smartcontract.GetBlockListInstance(client)
	res, err := bl.GetHostList(nil)
	if err != nil {
		log.Fatal(err)
	}
	blockerConfig.BlockList = res

	ch := make(chan *blocklist.BlocklistHostNameAdded)
	_, err = bl.WatchHostNameAdded(nil, ch)
	log.Print(err)
	if err != nil {
		log.Fatal(err)
	}
	for e := range ch {
		blockerConfig.BlockList = append(blockerConfig.BlockList, e.Hostname)
		log.Println(blockerConfig.BlockList)
	}
}
