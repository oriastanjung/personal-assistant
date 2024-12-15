package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	pb "github.com/oriastanjung/personal-va/proto/chat"

	"google.golang.org/grpc"
)

func main() {
	// Connect to gRPC server
	conn, err := grpc.Dial("localhost:50027", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to server: %v", err)
	}
	defer conn.Close()

	client := pb.NewChatServiceClient(conn)
	stream, err := client.Chat(context.Background())
	if err != nil {
		log.Fatalf("Failed to create stream: %v", err)
	}

	// Create a goroutine to receive server responses
	go func() {
		for {
			response, err := stream.Recv()
			if err != nil {
				log.Fatalf("Error receiving response: %v", err)
			}
			fmt.Printf("\nORIVA: %s\n", response.Content)
		}
	}()

	// Send user input to the server
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("\nYou: ")
		if !scanner.Scan() {
			break
		}
		userInput := strings.TrimSpace(scanner.Text())
		if userInput == "exit" {
			log.Println("Exiting chat...")
			break
		}

		err := stream.Send(&pb.ChatMessage{
			Role:    "user",
			Content: userInput,
		})
		if err != nil {
			log.Fatalf("Error sending message: %v", err)
		}
	}

	stream.CloseSend()
}
