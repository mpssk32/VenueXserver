package config

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
)

var DB *pgx.Conn

func ConnectDB() {
	dbUrl := os.Getenv("DATABASE_URL")

	conn, err := pgx.Connect(context.Background(), dbUrl)
	if err != nil {
		log.Fatal("Database connection error:", err)
	}

	DB = conn

	log.Println("PostgreSQL connected")
}