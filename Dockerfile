FROM golang:1.19 AS builder

COPY . /src
WORKDIR /src

RUN go build -ldflags="-X 'main.Version=v0.7.4' -X 'main.Time=$(LC_TIME=en_US.UTF-8 date)' -X 'main.Commit=$(git rev-parse --short HEAD)'" -o k8s-device-plugin cmd/manager.go

FROM ubuntu:22.04

COPY --from=builder /src/k8s-device-plugin /root/k8s-device-plugin

WORKDIR /root/