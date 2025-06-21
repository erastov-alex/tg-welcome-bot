package handler

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"tg-welcome-bot/internal/db"

	tb "gopkg.in/telebot.v3"
)

func ConfirmOrderHandler(c tb.Context) error {
	ctx := context.TODO()
	productID, err := strconv.Atoi(c.Data())
	if err != nil {
		return c.Respond(&tb.CallbackResponse{Text: "Ошибка: неверный ID товара"})
	}

	userID := c.Sender().ID

	err = db.SaveOrder(ctx, productID, userID)
	if err != nil {
		log.Printf("Ошибка сохранения заказа: %v", err)
		return c.Respond(&tb.CallbackResponse{Text: "Не удалось сохранить заказ."})
	}

	paymentURL := os.Getenv("PAYMENT_URL") + strconv.Itoa(productID)

	// Уведомление в чате и отклик на кнопку
	_, _ = c.Bot().Send(c.Sender(), fmt.Sprintf(
		"🎉 Ваш заказ подтверждён!\n\nПерейдите по ссылке для оплаты:\n🔗 %s", paymentURL,
	))

	return c.Respond(&tb.CallbackResponse{Text: "✅ Заказ подтверждён!"})
}
