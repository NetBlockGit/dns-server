package dnsblocker

import "github.com/NetBlockGit/dnsblocker/config"

var dnsBlocker *config.BlockerConfig

// Returns blocker instance and init it if it doesn't exist
func CheckInitAndGet() *config.BlockerConfig {
	if dnsBlocker != nil {
		return dnsBlocker
	}
	dnsBlocker = &config.BlockerConfig{
		UpstreamDns: "1.1.1.1:53",
		BlockList:   []string{},
		Addr:        "127.0.0.1:5301",
		Enabled:     true,
	}
	go dnsBlocker.StartDnsServer()
	return dnsBlocker
	// lis, err := net.Listen("tcp", fmt.Sprintf(":%v", 8000))
	// grpcServer := grpc.NewServer()
	// grpc.UnaryInterceptor(auth.CheckAuth)
}
