package openrouterai

//Здесь логика работа с апи ИИ

import (
	"context"
	"encoding/json"
	"tgbotai/internal/config"
	"tgbotai/internal/models"
	"tgbotai/internal/utils"

	"github.com/go-resty/resty/v2"
	"github.com/sirupsen/logrus"
)

func CallQWEN(ctx context.Context, prompt string, username string) (string, error) {
	client := resty.New()

	request := &models.Body{
		Model: "qwen/qwq-32b:free",
		Messages: []models.Message{
			{Role: "user", Content: prompt},
		},
	}
	logrus.Info("Отправлен запрос от @" + username)
	resp, err := client.R().
		SetContext(ctx).
		SetHeader("Authorization", config.QWEN_API_TOKEN).
		SetHeader("Content-Type", "application/json").
		SetBody(request).
		Post("https://openrouter.ai/api/v1/chat/completions")

	if err != nil {
		return "", err
	}
	
	logrus.Info("Ответ QWEN получен")

	result := &models.Response{}
	err = json.Unmarshal(resp.Body(), result)

	if err != nil {
		return "",err
	}

	res := utils.EditResult(result.Choices[0].Message.Content)
	res = utils.RemoveThinkBlockFromAnswer(res)
	logrus.Info("Ответ преобразован")

	return res, nil
}