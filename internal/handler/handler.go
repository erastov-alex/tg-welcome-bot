package handler

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"strings"

	"tg-welcome-bot/internal/db"

	tb "gopkg.in/telebot.v3"
)

var BtnConfirmOrder = tb.Btn{Unique: "confirm_order"}

// Структура для хранения контекста товара, чтобы использовать в callback
type productCallback struct {
	ProductID int
}

func StartHandler(c tb.Context) error {
	ctx := context.TODO()
	args := c.Args()
	if len(args) == 0 {
		return c.Send("Пожалуйста, укажите ID товара. Пример: /start 123")
	}

	itemID, err := strconv.Atoi(strings.TrimSpace(args[0]))
	if err != nil {
		return c.Send("Неверный формат ID товара.")
	}

	product, err := db.GetProduct(ctx, itemID)
	if err != nil {
		log.Printf("Товар не найден: %v", err)
		return c.Send("😔 Товар с таким ID не найден.")
	}

	// Отображаем товар с кнопкой
	photo := &tb.Photo{
		File: tb.FromURL("https://static.insales-cdn.com/files/1/6197/40482869/original/%D0%B0%D1%82%D0%BB%D0%B5%D1%82%D0%B8%D0%B7%D0%BC%D0%BE_600%D1%85600_78a0fd89473c72c9c5401dd95a8e9acd.png"),
		Caption: fmt.Sprintf(
			"🛍️ <b>%s</b>\n👟 <i>%s</i>\n\n💵 Цена: <b>%d ₽</b>\n📏 Размер: US %s / EU %s\n\nНажмите кнопку ниже, чтобы подтвердить заказ.",
			product.Name, product.Brand, product.Price, product.SizeUS, product.SizeEU,
		),
	}

	markup := &tb.ReplyMarkup{}
	btn := markup.Data("✅ Подтвердить заказ", BtnConfirmOrder.Unique, strconv.Itoa(product.ID))
	markup.Inline(markup.Row(btn))

	return c.Send(photo, &tb.SendOptions{
		ParseMode:   tb.ModeHTML,
		ReplyMarkup: markup,
	})
}
