# 第一阶段：构建 Go 可执行文件
FROM golang:1.23.2 AS builder

RUN export GO111MODULE="on"
RUN mkdir /go/src/app
RUN export GOPATH="/go/src"
RUN export  GOPROXY="https://goproxy.cn,direct"

ADD . /go/src/app

WORKDIR /go/src/app
RUN pwd
RUN  go mod download

RUN go build -o main .
ENTRYPOINT ["./main"]