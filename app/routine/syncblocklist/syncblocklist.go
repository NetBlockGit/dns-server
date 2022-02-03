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

	go syncAdd()
	go syncDelete()
}

func syncDelete() {
	blockerConfig := dnsblocker.CheckInitAndGet()
	client := smartcontract.GetClient()
	bl := smartcontract.GetBlockListInstance(client)
	res, err := bl.GetHostList(nil)
	if err != nil {
		log.Fatal(err)
	}
	blockerConfig.BlockList = res

	ch := make(chan *blocklist.BlocklistHostNameDeleted)
	_, err = bl.WatchHostNameDeleted(nil, ch)
	if err != nil {
		log.Fatal(err)
	}
	for e := range ch {
		blockerConfig.BlockList[e.Index.Int64()] = blockerConfig.BlockList[len(blockerConfig.BlockList)-1]
		blockerConfig.BlockList = blockerConfig.BlockList[:len(blockerConfig.BlockList)-1]
	}
}

func syncAdd() {
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
	if err != nil {
		log.Fatal(err)
	}
	for e := range ch {
		blockerConfig.BlockList = append(blockerConfig.BlockList, e.Hostname)
	}
}
