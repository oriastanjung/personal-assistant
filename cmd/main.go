package main

import (
	"log"
	"net"

	"github.com/oriastanjung/personal-va/internal/config"
	"github.com/oriastanjung/personal-va/internal/services"
	"github.com/oriastanjung/personal-va/internal/usecase"
	"google.golang.org/grpc"

	pb "github.com/oriastanjung/personal-va/proto/chat"
)

func main() {
	cfg := config.LoadEnv()

	// Initialize usecase
	usecase := usecase.NewChatUsecase(cfg.OPENAI_API_KEY)

	// Initialize gRPC server
	server := grpc.NewServer()
	chatService := services.NewChatService(usecase)
	pb.RegisterChatServiceServer(server, chatService)

	// Start listening
	listener, err := net.Listen("tcp", ":"+cfg.PORT)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	log.Println("Chat server is running on  0.0.0.0:" + cfg.PORT)

	if err := server.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
