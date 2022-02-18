package app

import (
	"dnsserver/app/routine/syncblocklist"
	"dnsserver/config/dnsblocker"
	"dnsserver/config/env"
	"dnsserver/config/mongodb"
	"dnsserver/gRPC/grpcserver"
	"net/http"
	"regexp"
)

func Init() {
	env.Init(".env")
	mongodb.Init()
	dnsblocker.CheckInitAndGet()
	syncblocklist.Init()

	fileServer := http.FileServer(http.Dir("public"))
	fileMatcher := regexp.MustCompile(`\.[a-zA-Z]*$`)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if !fileMatcher.MatchString(r.URL.Path) {
			http.ServeFile(w, r, "public/index.html")
		} else {
			fileServer.ServeHTTP(w, r)
		}
	})
	go http.ListenAndServe(":8001", nil)

	grpcserver.Init()
}
