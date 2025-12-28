package llm

import (
	"context"
	"encoding/json"
	"errors"
	"os"
	"strings"

	"newsApp/internal/models"

	openai "github.com/sashabaranov/go-openai"
)

func AnalyzeQuery(query string) (*models.LLMResult, error) {
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		return nil, errors.New("OPENAI_API_KEY not set")
	}

	client := openai.NewClient(apiKey)

	prompt := `
Extract entities, intent, concepts, and location from this query.
Respond ONLY in valid JSON. Do not add explanations.

{
  "intent": "",
  "entities": [],
  "concepts": [],
  "location": ""
}

Query: ` + query

	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: "openai.GPT-4o-mini",
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: prompt,
				},
			},
		},
	)
	if err != nil {
		return nil, err
	}

	if len(resp.Choices) == 0 {
		return nil, errors.New("empty response from LLM")
	}

	content := strings.TrimSpace(resp.Choices[0].Message.Content)

	var result models.LLMResult
	if err := json.Unmarshal([]byte(content), &result); err != nil {
		return nil, err
	}

	return &result, nil
}
