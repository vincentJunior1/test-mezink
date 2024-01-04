#!/bin/bash
# Build stage
FROM golang:1.19-alpine3.16 AS builder
WORKDIR /app
RUN apk add alpine-sdk
COPY . .
# RUN go mod download
RUN GOOS=linux GOARCH=amd64 go build -o skeleton-svc -tags musl

# Run stage
FROM alpine:3.16
WORKDIR /app

RUN apk update && apk add --no-cache git
RUN apk update && apk add --no-cache tzdata
ENV TZ="Asia/Jakarta"

COPY --from=builder /app/skeleton-svc .
# COPY .env .
# COPY start.sh .
COPY errorcodes.json .

EXPOSE 7878
ENTRYPOINT [ "/app/skeleton-svc" ]
# ENTRYPOINT [ "/app/start.sh" ]