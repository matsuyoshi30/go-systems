FROM golang:1.14.1-alpine3.11

RUN apk add --update openssl-dev pcre-dev git gcc musl-dev linux-headers sudo

WORKDIR /go/src/github.com/matsuyoshi30/go-systems/ch17/container/
ADD main.go ./main.go
ADD rootfs ./rootfs/
RUN go get github.com/opencontainers/runc
RUN go get golang.org/x/sys/unix
RUN go build main.go
CMD echo "test"