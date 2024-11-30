package db

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"os"

	_ "github.com/jackc/pgx/v5/pgxpool"
)

var DB *pgxpool.Pool

func InitDB() {
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("DATABASE_URL environment variable not set")
	}
	config, err := pgxpool.ParseConfig(dbURL)
	if err != nil {
		log.Fatalf("Unable to parse DATABASE_URL: %v\n", err)
	}

	DB, err = pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		log.Fatalf("Unable to create connection pool: %v\n", err)
	}

	err = DB.Ping(context.Background())
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}

	createTable()
}

func createTable() {
	query := `
    CREATE TABLE IF NOT EXISTS pastes (
        id TEXT PRIMARY KEY,
        content TEXT NOT NULL,
        created_at TIMESTAMP NOT NULL
    );`

	_, err := DB.Exec(context.Background(), query)

	if err != nil {
		log.Fatalf("Unable to create table: %v\n", err)
	}
}

func CloseDB() {
	DB.Close()
}
