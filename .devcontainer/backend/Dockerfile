FROM golang:1.16.3-alpine

RUN apk add --update &&  apk add git && apk add bash

RUN mkdir -p /backend/app

WORKDIR /backend/app

# CMD go run main.go