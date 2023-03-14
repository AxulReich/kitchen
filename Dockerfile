FROM golang:1.19.2-alpine3.16 AS build

RUN apk --no-cache add gcc g++ make
RUN apk add git
WORKDIR /go/src/app
COPY . .
RUN go mod download
RUN GOOS=linux go build -ldflags="-s -w" -o ./bin/kitchen ./cmd/main.go


FROM alpine:latest
WORKDIR /usr/bin
COPY --from=build /go/src/app/bin /go/bin

