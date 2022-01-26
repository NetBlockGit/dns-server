FROM gitpod/workspace-full
RUN sudo apt update && sudo apt install -y protobuf-compiler \
      go get github.com/golang/protobuf/protoc-gen-go \
      go get google.golang.org/grpc/cmd/protoc-gen-go-grpc \