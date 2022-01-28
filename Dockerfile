FROM golang:alpine as builder
WORKDIR /app
RUN apk add build-base
COPY go.mod .
RUN go mod download
COPY . .
RUN go build -o dnsserver
FROM alpine
WORKDIR /app
COPY --from=builder /app/dnsserver .
CMD ./dnsserver