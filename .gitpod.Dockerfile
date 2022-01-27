FROM gitpod/workspace-full
USER root
RUN sudo apt update && sudo apt install -y protobuf-compiler
RUN go install github.com/golang/protobuf/protoc-gen-go@latest
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest