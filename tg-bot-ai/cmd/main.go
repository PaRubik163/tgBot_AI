package main

import (
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"tgbotai/internal/config"
	hndl "tgbotai/internal/bot"
	"github.com/joho/godotenv"
)

func main() {
	
	err := godotenv.Load()
	if err != nil{
		log.Fatal(err)
	}

	config.BOT_TOKEN = os.Getenv("BOT_TOKEN")
	config.QWEN_API_TOKEN = os.Getenv("QWEN_API")

	bot, err := tgbotapi.NewBotAPI(config.BOT_TOKEN)
	
	if err != nil{
		log.Fatalf("Ошибка в запуске бота")
	}

	bot.Debug = true
	log.Printf("Авторизован как %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60 
	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue 
		}

		go func ()  {
			err := hndl.HandlerQWEN(bot, update.Message)

			if err != nil{
				bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Ошибка API"))
			}
		}()
	}
}