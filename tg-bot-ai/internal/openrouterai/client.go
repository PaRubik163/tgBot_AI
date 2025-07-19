package openrouterai

//Здесь логика работа с апи ИИ

import (
	"encoding/json"
	"tgbotai/internal/config"
	"tgbotai/internal/models"
	"tgbotai/internal/utils"
	"github.com/go-resty/resty/v2"
)

func CallQWEN(prompt string) (string, error) {
	client := resty.New()

	request := &models.Body{
		Model: "qwen/qwq-32b:free",
		Messages: []models.Message{
			{Role: "user", Content: prompt},
		},
	}

	resp, err := client.R().SetHeader("Authorization", config.QWEN_API_TOKEN).SetHeader("Content-Type", "application/json").SetBody(request).
		Post("https://openrouter.ai/api/v1/chat/completions")

	if err != nil {
		return "", err
	}

	result := &models.Response{}
	err = json.Unmarshal(resp.Body(), result)

	if err != nil {
		return "",err
	}

	res := utils.EditResult(result.Choices[0].Message.Content)
	res = utils.RemoveThinkBlockFromAnswer(res)
	
	return res, nil
}