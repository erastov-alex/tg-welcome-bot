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
		File: tb.FromURL("https://placekitten.com/600/400"),
		Caption: fmt.Sprintf(
			"🛍️ <b>%s</b>\n👟 <i>%s</i>\n\n💵 Цена: <b>%d ₽</b>\n📏 Размер: US %s / EU %s\n\nНажмите кнопку ниже, чтобы подтвердить заказ.",
			product.Name, product.Brand, product.Price, product.SizeUS, product.SizeEU,
		),
	}

	markup := &tb.ReplyMarkup{}
	btn := markup.Data("✅ Подтвердить заказ", "confirm_order", strconv.Itoa(product.ID))
	markup.Inline(markup.Row(btn))

	return c.Send(photo, &tb.SendOptions{
		ParseMode:   tb.ModeHTML,
		ReplyMarkup: markup,
	})
}
