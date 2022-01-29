FROM golang:buster as builder
WORKDIR /app
COPY go.mod .
RUN go mod download
COPY . .
RUN go build -o dnsserver
FROM debian:buster
RUN  apt update
RUN  apt install debian-keyring debian-archive-keyring apt-transport-https curl lsb-release -y
RUN curl -sL 'https://deb.dl.getenvoy.io/public/gpg.8115BA8E629CC074.key' |  gpg --dearmor -o /usr/share/keyrings/getenvoy-keyring.gpg
RUN echo a077cb587a1b622e03aa4bf2f3689de14658a9497a9af2c427bba5f4cc3c4723 /usr/share/keyrings/getenvoy-keyring.gpg | sha256sum --check
RUN echo "deb [arch=amd64 signed-by=/usr/share/keyrings/getenvoy-keyring.gpg] https://deb.dl.getenvoy.io/public/deb/debian $(lsb_release -cs) main" |  tee /etc/apt/sources.list.d/getenvoy.list
RUN  apt update
RUN  apt install getenvoy-envoy -y
WORKDIR /app
COPY --from=builder /app/dnsserver .
COPY envoy.yaml .
COPY start.sh .
CMD ./start.sh