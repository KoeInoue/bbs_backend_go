FROM golang:1.16.3-alpine

ENV GO111MODULE on

WORKDIR /go/src/app

RUN apk update \
    && apk add git \
    && go install github.com/cosmtrek/air@latest