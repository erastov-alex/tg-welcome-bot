package db

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Order struct {
	ID        int
	ItemID    int
	UserID    int64
	CreatedAt time.Time
}

func SaveOrder(itemID int, userID int64) error {
	_, err := Pool.Exec(
		context.Background(),
		"INSERT INTO orders (item_id, user_id) VALUES ($1, $2)",
		itemID, userID,
	)
	return err
}

var Pool *pgxpool.Pool

func InitDB() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	dbUrl := os.Getenv("DATABASE_URL")
	if dbUrl == "" {
		log.Fatal("DATABASE_URL не задан")
	}

	var err error
	Pool, err = pgxpool.New(ctx, dbUrl)
	if err != nil {
		log.Fatalf("Ошибка подключения к PostgreSQL: %v", err)
	}

	createTable := `
	CREATE TABLE IF NOT EXISTS orders (
		id SERIAL PRIMARY KEY,
		item_id INTEGER NOT NULL,
		user_id BIGINT NOT NULL,
		created_at TIMESTAMPTZ DEFAULT now()
	);`
	_, err = Pool.Exec(ctx, createTable)
	if err != nil {
		log.Fatalf("Ошибка создания таблицы: %v", err)
	}
}
