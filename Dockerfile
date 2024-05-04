FROM golang:1.21

WORKDIR /go/src/app

COPY . .

RUN go get -d -v ./...

RUN go build -o /go/bin/app ./cmd/haha

EXPOSE 8080

CMD ["/go/bin/app"]