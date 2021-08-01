FROM golang:1.17rc1-alpine3.14
RUN apk update && apk add git
RUN mkdir /go/src/app
WORKDIR /go/src/app
ENV GO111MODULE=on 
RUN apk add gcc