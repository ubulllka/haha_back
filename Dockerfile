FROM golang:1.21

WORKDIR /app

COPY . .

RUN go mod tidy

EXPOSE 8080

CMD ["go", "run", "./cmd/haha/main.go"]