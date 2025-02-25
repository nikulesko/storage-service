package main

import (
	"context"
	"fmt"
	"github.com/nikulesko/golearn/storage-service/cmd/db"
	"github.com/nikulesko/golearn/storage-service/crypto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"os"
	"time"
)

func main() {
	dataServerHost := os.Getenv("DATA_SERVER_HOST")
	if dataServerHost == "" {
		panic("DATA_SERVER_HOST environment variable not set")
	}

	dataServerPort := os.Getenv("DATA_SERVER_PORT")
	if dataServerPort == "" {
		panic("DATA_SERVER_PORT environment variable not set")
	}

	conn, err := grpc.NewClient(fmt.Sprintf("%s:%s", dataServerHost, dataServerPort), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Cannot connect to server: %v", err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Fatalf("Cannot close connection: %v", err)
		}
	}(conn)

	client := crypto.NewCryptoRateClient(conn)

	req := &crypto.Request{
		BaseCurrencyCode:     "BTC",
		ExchangeCurrencyCode: "USD",
	}

	stream, err := client.GetDataStreaming(context.Background(), req)
	if err != nil {
		log.Fatalf("Failed to get stream: %v", err)
	}

	repo, err := db.NewRepository()
	if err != nil {
		log.Fatalf("Failed to create repository: %v", err)
	}

	for {
		res, err := stream.Recv()
		if err != nil {
			log.Fatalf("Failed to receive data: %v", err)
		}
		log.Printf("Received data: Datetime=%d, Rate=%.2f", res.Datetime, res.Rate)

		err = repo.SaveCurrencyData(&db.CurrencyData{Datetime: res.Datetime, Rate: res.Rate})
		if err != nil {
			log.Fatalf("Failed to insert currency data: %v", err)
		}
		log.Println("currency data saved")

		time.Sleep(500 * time.Millisecond)
	}
}
