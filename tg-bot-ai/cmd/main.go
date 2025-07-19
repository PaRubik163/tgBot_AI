package main

import (
	"context"
	"log"
	"os"
	"time"

	hndl "tgbotai/internal/bot"
	"tgbotai/internal/config"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func main() {
	
	err := godotenv.Load()
	if err != nil{
		logrus.Fatal(err)
	}

	config.BOT_TOKEN = os.Getenv("BOT_TOKEN")
	config.QWEN_API_TOKEN = os.Getenv("QWEN_API")

	bot, err := tgbotapi.NewBotAPI(config.BOT_TOKEN)
	
	if err != nil{
		logrus.Fatalf("Ошибка в запуске бота")
	}

	//bot.Debug = true
	log.Printf("Авторизован как %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60 
	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue 
		}

		go func ()  {
			ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
			defer cancel()

			err := hndl.HandlerQWEN(ctx, bot, update.Message)

			if err != nil{
				bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Ошибка API"))
				logrus.Warn("Ошибка в работе API")
			}
		}()
	}
}