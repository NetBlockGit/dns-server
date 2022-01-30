FROM gitpod/workspace-full
RUN sudo apt update && sudo apt install -y protobuf-compiler
RUN brew install okteto envoy grpc
RUN curl -LO "https://dl.k8s.io/release/$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/linux/amd64/kubectl" && \
    sudo install -o root -g root -m 0755 kubectl /usr/local/bin/kubectl
RUN curl -Lo skaffold https://storage.googleapis.com/skaffold/releases/latest/skaffold-linux-amd64 && \
    sudo install skaffold /usr/local/bin/