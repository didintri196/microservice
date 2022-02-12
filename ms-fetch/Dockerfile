FROM golang:1.16

ENV GO111MODULE=on

ENV PKG_NAME=ms-fetch
ENV PKG_PATH=$GOPATH/src/$PKG_NAME

WORKDIR $PKG_PATH/
COPY . $PKG_PATH/

RUN go mod vendor
WORKDIR $PKG_PATH/server/http

RUN go build main.go
CMD ["sh", "-c", "./main"]