FROM golang:1.15.12-alpine3.13

ENV GO111MODULE on

WORKDIR /go/src/app

RUN apk update \
    && apk add git \
    gcc \
    musl-dev \
    && go get -u github.com/cosmtrek/air@latest

RUN go get github.com/jinzhu/gorm \
    && go get -u github.com/jinzhu/gorm/dialects/mysql \
    && go get -u github.com/gin-gonic/gin \
    && go get -u github.com/joho/godotenv \
    && go get -u github.com/gin-contrib/cors \
    && go get bitbucket.org/liamstask/goose/cmd/goose