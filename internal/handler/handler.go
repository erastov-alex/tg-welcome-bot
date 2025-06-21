package handler

import (
	"fmt"
	"log"
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
	paymentURL := "TEST_PAYMENT_LINK"

	// Картинка (можно заменить на любую: локальную, URL или []byte)
	photo := &tb.Photo{
		File:    tb.FromURL("https://static.insales-cdn.com/files/1/6197/40482869/original/%D0%B0%D1%82%D0%BB%D0%B5%D1%82%D0%B8%D0%B7%D0%BC%D0%BE_600%D1%85600_78a0fd89473c72c9c5401dd95a8e9acd.png"), // Замените на свою картинку товара
		Caption: fmt.Sprintf("🎉 Спасибо за выбор товара #%d!\n\n🛒 Переходите по ссылке, чтобы оплатить заказ:\n%s", itemID, paymentURL),
	}

	return c.Send(photo)
}
