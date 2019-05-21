FROM golang:1.12.5-alpine3.9

RUN go env GOPATH
WORKDIR /go/src/github.com
RUN mkdir -p /brieefly/server
WORKDIR /brieefly/server
ADD . .

RUN apk update && \
    apk upgrade && \
    apk add git bash

RUN go get github.com/google/uuid
RUN go get github.com/go-sql-driver/mysql
RUN go get github.com/go-chi/chi
RUN go get github.com/dgrijalva/jwt-go

RUN go build
RUN go run app.go