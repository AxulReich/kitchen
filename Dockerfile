FROM golang:1.19.2-alpine3.16 AS build

RUN apk --no-cache add gcc g++
WORKDIR /go/src/app
COPY . .
RUN go mod download
RUN GOOS=linux go build -o ./bin/kitchen ./cmd/main.go


FROM alpine:3.16
WORKDIR /usr/bin
EXPOSE 7002 7000 8080
COPY --from=build /go/src/app/bin /usr/bin

CMD ["/usr/bin/kitchen"]