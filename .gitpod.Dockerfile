FROM gitpod/workspace-full
RUN sudo apt update && sudo apt install -y protobuf-compiler
RUN sudo -E env "PATH=$PATH" go install github.com/golang/protobuf/protoc-gen-go@latest
RUN sudo -E env "PATH=$PATH" go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest