package chat

import (
	"context"
	"log"
)

type Server struct {
	UnimplementedChatServiceServer
}

func (s *Server) Response(ctx context.Context, message *Chat) (*Chat, error) {
	log.Printf("Message received from client: %s", message.Body)

	return &Chat{
		Body: "Server is responding...",
	}, nil
}