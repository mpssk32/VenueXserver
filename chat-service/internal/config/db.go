package config

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

var DB *pgxpool.Pool

func ConnectDB() {

	databaseURL := os.Getenv("DATABASE_URL")

	var err error

	DB, err = pgxpool.New(
		context.Background(),
		databaseURL,
	)

	if err != nil {
		log.Fatal(err)
	}

	err = DB.Ping(context.Background())

	if err != nil {
		log.Fatal(err)
	}

	log.Println("PostgreSQL connected")
}