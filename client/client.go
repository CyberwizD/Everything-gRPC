package main

import (
	"context"
	"log"

	"github.com/CyberwizD/Everything-gRPC/chat"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	var connect *grpc.ClientConn

	// Initialize new client with disabled transport security
	connect, err := grpc.NewClient(":9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}

	defer connect.Close()

	// Initialize new chat client
	chatClient := chat.NewChatServiceClient(connect)

	// Client message
	message := &chat.Chat {
		Body: "Message from the client",
	}

	// Return gRPC client reponse
	response, err := chatClient.Response(context.Background(), message)
	if err != nil {
		log.Fatalf("Failed to send client response, %v", err)
	}

	log.Printf("Response from server: %s", response.Body)
}