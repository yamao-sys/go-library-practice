# syntax=docker/dockerfile:1

FROM golang:1.23.1-alpine3.20

# RUN apk add --no-cache git sqlite-dev gcc musl-dev
RUN apk add --no-cache git sqlite gcc musl-dev

WORKDIR /app

COPY . .

RUN go install golang.org/x/tools/cmd/goimports@latest
