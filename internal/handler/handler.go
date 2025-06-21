package handler

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"tg-welcome-bot/internal/db"

	tb "gopkg.in/telebot.v3"
)

func StartHandler(c tb.Context) error {
	args := c.Args()
	if len(args) == 0 {
		return c.Send("Пожалуйста, передайте ID товара. Пример: /start 123")
	}

	itemID, err := strconv.Atoi(strings.TrimSpace(args[0]))
	if err != nil {
		return c.Send("Неверный формат ID товара.")
	}

	userID := c.Sender().ID
	err = db.SaveOrder(itemID, userID)
	if err != nil {
		log.Printf("Ошибка сохранения заказа: %v", err)
		return c.Send("Ошибка при сохранении заказа.")
	}

	// URL оплаты
	paymentURL := os.Getenv("PAYMENT_URL") + strconv.Itoa(itemID)

	// Картинка (можно заменить на любую: локальную, URL или []byte)
	photo := &tb.Photo{
		File:    tb.FromURL("https://placekitten.com/600/400"), // Замените на свою картинку товара
		Caption: fmt.Sprintf("🎉 Спасибо за выбор товара #%d!\n\n🛒 Переходите по ссылке, чтобы оплатить заказ:\n%s", itemID, paymentURL),
	}

	return c.Send(photo)
}
