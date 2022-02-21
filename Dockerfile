FROM golang:alpine3.15

WORKDIR /opt

ENV USE_NEW_FEATURE=true

COPY main.go go.mod ./

CMD go run .

