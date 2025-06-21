package main

import (
	"log"
	"os"
	"time"

	"tg-welcome-bot/internal/db"
	"tg-welcome-bot/internal/handler"

	"github.com/joho/godotenv"
	tb "gopkg.in/telebot.v3"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("Файл .env не найден, использую переменные среды")
	}

	db.InitDB()

	pref := tb.Settings{
		Token:  os.Getenv("BOT_TOKEN"),
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	}

	bot, err := tb.NewBot(pref)
	if err != nil {
		log.Fatalf("Ошибка создания бота: %v", err)
	}

	bot.Handle("/start", handler.StartHandler)

	log.Println("Бот запущен...")
	bot.Start()
}
