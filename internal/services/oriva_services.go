package services

import (
	"context"
	"errors"
	"io"
	"log"
	"sync"

	"github.com/sashabaranov/go-openai"

	"github.com/oriastanjung/personal-va/internal/usecase"
	pb "github.com/oriastanjung/personal-va/proto/chat"
)

type ChatService struct {
	pb.UnimplementedChatServiceServer
	Usecase *usecase.ChatUsecase
}

func NewChatService(usecase *usecase.ChatUsecase) *ChatService {
	return &ChatService{Usecase: usecase}
}

func (s *ChatService) Chat(stream pb.ChatService_ChatServer) error {
	var chatHistory []openai.ChatCompletionMessage
	mu := sync.Mutex{}

	chatHistory = append(chatHistory, openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleSystem,
		Content: "Anda adalah ORIVA, asisten pribadi yang ramah dan selalu menggunakan bahasa Indonesia santai.",
	})

	for {
		userMsg, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			log.Println("Chat session ended by client.")
			return nil
		}
		if err != nil {
			log.Printf("Error receiving stream: %v", err)
			return err
		}

		log.Printf("Received message from user: %s", userMsg.Content)

		// Append user message to history
		mu.Lock()
		chatHistory = append(chatHistory, openai.ChatCompletionMessage{
			Role:    openai.ChatMessageRoleUser,
			Content: userMsg.Content,
		})
		mu.Unlock()

		// Generate assistant response
		response, err := s.Usecase.GenerateResponse(context.Background(), chatHistory)
		if err != nil {
			log.Printf("Error generating response: %v", err)
			return err
		}

		// Append assistant response to history
		mu.Lock()
		chatHistory = append(chatHistory, openai.ChatCompletionMessage{
			Role:    openai.ChatMessageRoleAssistant,
			Content: response,
		})
		mu.Unlock()

		// Send assistant response to client
		err = stream.Send(&pb.ChatMessage{
			Role:    "assistant",
			Content: response,
		})
		if err != nil {
			log.Printf("Error sending response: %v", err)
			return err
		}
	}
}
