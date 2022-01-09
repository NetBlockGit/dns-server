.PHONY: protos

.DEFAULT_GOAL := protos
protos:
	protoc protos/blocker.proto --go_out=. --go-grpc_out=.