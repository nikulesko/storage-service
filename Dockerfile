FROM golang:1.23.6

WORKDIR /cmd/client

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o main .

CMD ["./main"]