package main

import (
	"log"
	"os"
	"time"

	"tg-welcome-bot/internal/handler"

	tb "gopkg.in/telebot.v3"
)

func main() {
	pref := tb.Settings{
		Token:  os.Getenv("BOT_TOKEN"),
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	}

	bot, err := tb.NewBot(pref)
	if err != nil {
		log.Fatal(err)
	}

	bot.Handle("/start", handler.StartHandler)
	bot.Handle(&handler.BtnConfirmOrder, handler.ConfirmOrderHandler)

	log.Println("🤖 Бот запущен!")
	bot.Start()
}
