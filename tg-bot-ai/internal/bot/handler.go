package bot

import (
	"tgbotai/internal/openrouterai"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func HandlerQWEN(bot *tgbotapi.BotAPI, prompt *tgbotapi.Message) error {
	result, err := openrouterai.CallQWEN(prompt.Text)

	if err != nil{
		return err
	}

	msg := tgbotapi.NewMessage(prompt.Chat.ID, result)
	if _, err := bot.Send(msg); err != nil{
		return err
	}

	return nil
}