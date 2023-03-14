FROM golang:1.20

WORKDIR /usr/src/app

COPY go.mod ./
RUN go mod download

COPY . .
RUN go build -v -o /usr/src/app ./...

CMD ["app"]