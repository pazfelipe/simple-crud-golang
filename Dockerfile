FROM golang:1.24.1-alpine as builder

WORKDIR /app

COPY . .

EXPOSE 8000

RUN go build -o main ./cmd/main.go

CMD ["./main"]