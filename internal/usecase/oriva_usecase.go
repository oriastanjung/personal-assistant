package usecase

import (
	"context"
	"errors"
	"io"
	"strings"

	"github.com/sashabaranov/go-openai"
)

type ChatUsecase struct {
	Client *openai.Client
}

func NewChatUsecase(apiKey string) *ChatUsecase {
	client := openai.NewClient(apiKey)
	return &ChatUsecase{Client: client}
}

func (u *ChatUsecase) GenerateResponse(ctx context.Context, chatHistory []openai.ChatCompletionMessage) (string, error) {
	stream, err := u.Client.CreateChatCompletionStream(ctx, openai.ChatCompletionRequest{
		Model:    openai.GPT4,
		Messages: chatHistory,
		Stream:   true,
	})
	if err != nil {
		return "", err
	}
	defer stream.Close()

	var responseBuilder strings.Builder
	for {
		resp, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			return "", err
		}
		content := resp.Choices[0].Delta.Content
		responseBuilder.WriteString(content)
	}
	return responseBuilder.String(), nil
}
