.PHONY: protos

.DEFAULT_GOAL := protos
protos:
	mkdir -p out
	protoc protos/blocker.proto --go_out=out --go-grpc_out=out
	protoc protos/toggleblocker/toggleblocker.proto --go_out=out --go-grpc_out=out
	protoc protos/getstats/getstats.proto --go_out=out --go-grpc_out=out
	protoc protos/getauthtoken/getauthtoken.proto --go_out=out --go-grpc_out=out
	protoc protos/updateupstreamdns/updateupstreamdns.proto --go_out=out --go-grpc_out=out
	rm -rf generated/protos
	cp -r out/dnsserver/generated/protos generated/protos
	rm -rf out
