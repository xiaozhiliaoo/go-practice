package openai

import (
	"context"
	"fmt"
	"github.com/sashabaranov/go-openai"
)

// ChatGPTMsg gpe消息
type ChatGPTMsg struct {
	Role    string `json:"role"`    // 角色 system:系统，user:用户，assistan助手
	Content string `json:"content"` // 内容
}

// Completion 获取Completion
func Completion(ctx context.Context, url, token, model string, messages []ChatGPTMsg) (completionResponse openai.ChatCompletionResponse, err error) {

	var msg []openai.ChatCompletionMessage
	for _, m := range messages {
		msg = append(msg, openai.ChatCompletionMessage{
			Role:    m.Role,
			Content: m.Content,
		})
	}

	model = openai.GPT3Dot5Turbo

	token = ""

	cfg := openai.DefaultConfig(token)
	cfg.BaseURL = url
	client := openai.NewClientWithConfig(cfg)
	resp, err := client.CreateChatCompletion(
		ctx,
		openai.ChatCompletionRequest{
			Model:    model,
			Messages: msg,
		},
	)

	if err != nil {
		fmt.Printf("Completion error: %v,messages:%s,model:%s", err, messages, model)
		return openai.ChatCompletionResponse{}, err
	}

	fmt.Printf("Completion question:%s,messages:%s,resp:%+v\n", messages, model, resp)
	return resp, nil
}
