package lib

import (
	openai "github.com/sashabaranov/go-openai"
)

func CreateClient(apikey string) *openai.Client {
	return openai.NewClient(apikey)
}
