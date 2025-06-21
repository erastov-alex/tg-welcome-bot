package db

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

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

type Product struct {
	ID     int
	Name   string
	Brand  string
	Price  int
	SizeUS string
	SizeEU string
}

func GetProduct(ctx context.Context, id int) (*Product, error) {
	query := `
	SELECT id, name, brand, price, size_us, size_eu FROM products_male WHERE id = $1
	UNION
	SELECT id, name, brand, price, size_us, size_eu FROM products_female WHERE id = $1
	LIMIT 1;
	`

	row := Pool.QueryRow(ctx, query, id)

	var p Product
	err := row.Scan(&p.ID, &p.Name, &p.Brand, &p.Price, &p.SizeUS, &p.SizeEU)
	if err != nil {
		// if errors.Is(err, pgxpool.ErrNoRows) {
		// 	return nil, fmt.Errorf("товар не найден")
		// }
		return nil, err
	}

	return &p, nil
}

func SaveOrder(ctx context.Context, productID int, userID int64) error {
	_, err := Pool.Exec(ctx, `INSERT INTO orders (item_id, user_id) VALUES ($1, $2)`, productID, userID)
	return err
}
