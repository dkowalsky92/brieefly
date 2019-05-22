FROM golang:1.12.5-alpine3.9

RUN apk update && \
    apk upgrade && \
    apk add git bash

RUN mkdir /brieefly
WORKDIR /brieefly
ADD . .

RUN export GO111MODULE=on
RUN chmod +x ./file-exist.sh ; ./file-exist.sh "go.mod" "go mod init github.com/dkowalsky/brieefly"

RUN go build
RUN go run app.go