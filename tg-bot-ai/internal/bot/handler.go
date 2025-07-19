package bot

import (
	"tgbotai/internal/openrouterai"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"context"
)

func HandlerQWEN(ctx context.Context, bot *tgbotapi.BotAPI, prompt *tgbotapi.Message) error {
	result, err := openrouterai.CallQWEN(ctx, prompt.Text, prompt.From.UserName)

	if err != nil{
		return err
	}

	msg := tgbotapi.NewMessage(prompt.Chat.ID, result)
	if _, err := bot.Send(msg); err != nil{
		return err
	}

	return nil
}