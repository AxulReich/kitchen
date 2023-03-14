# syntax=docker/dockerfile:1

## Build
FROM golang:1.20 AS builder
WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY . .
RUN go build -o /kitchen

## Deploy
FROM alpine:3.11

WORKDIR /

COPY --from=build /kitchen /kitchen

EXPOSE 8080 7000 7002

USER nonroot:nonroot

ENTRYPOINT ["/kitchen"]