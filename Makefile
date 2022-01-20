.PHONY: protos

.DEFAULT_GOAL := protos
protos:
	protoc protos/blocker.proto --go_out=. --go-grpc_out=.
	protoc protos/models/toggleblocker.proto --go_out=. --go-grpc_out=.
	protoc protos/models/getstats.proto --go_out=. --go-grpc_out=.