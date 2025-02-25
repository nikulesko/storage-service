FROM golang:1.23.6 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

WORKDIR /app/cmd/client

RUN CGO_ENABLED=0 GOOS=linux go build -o /app/storage-service

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/storage-service /app/storage-service

CMD ["./storage-service"]