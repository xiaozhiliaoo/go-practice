package openai

import (
	"context"
	"fmt"
	"github.com/sashabaranov/go-openai"
	"os"
)

func main() {
	apiKey := os.Getenv("OPENAI_API_KEY")
	client := openai.NewClient(apiKey)

	fmt.Println("你好，我是一个聊天机器人，请你提出你的问题吧?")

	questions := []string{}
	answers := []string{}

	for {
		var userInput string
		fmt.Scanln(&userInput)
		questions = append(questions, userInput)

		if userInput == "bye" || userInput == "goodbye" || userInput == "exit" {
			fmt.Println("Goodbye!")
			break
		}

		prompt := generatePrompt("", questions, answers)

		answer := askGPT3(client, prompt)
		fmt.Println(answer)
		answers = append(answers, answer)
	}
}

func generatePrompt(prompt string, questions, answers []string) string {
	num := len(answers)
	for i := 0; i < num; i++ {
		prompt += "\n Q : " + questions[i]
		prompt += "\n A : " + answers[i]
	}
	prompt += "\n Q : " + questions[num] + "\n A : "
	return prompt
}

func askGPT3(client *openai.Client, prompt string) string {
	response, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model:       "text-davinci-003",
			Prompt:      prompt,
			MaxTokens:   512,
			N:           1,
			Stop:        nil,
			Temperature: 0.5,
		},
	)

	if err != nil {
		fmt.Println("Error:", err)
		return ""
	}

	message := response.Choices[0].Text.String()
	return message
}
