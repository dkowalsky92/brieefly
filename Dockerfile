FROM golang:alpine AS builder
# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git
COPY . $GOPATH/src/github.com/dkowalsky/brieefly
WORKDIR $GOPATH/src/github.com/dkowalsky/brieefly
# Fetch dependencies.
# Using go get.
RUN go get -d -v
# Build the binary.
RUN go build -o /go/bin/brieefly
############################
# STEP 2 build a small image
############################
FROM scratch
# Copy our static executable.
COPY --from=builder /go/bin/brieefly /go/bin/brieefly

EXPOSE 5000
# Run the hello binary.
ENTRYPOINT ["/go/bin/brieefly"]
