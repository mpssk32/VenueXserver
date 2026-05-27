package config

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

var DB *pgxpool.Pool

func ConnectDB() {

	godotenv.Load()

	databaseURL := os.Getenv("DATABASE_URL")

	if databaseURL == "" {

		databaseURL =
			"postgres://postgres:postgres@localhost:5432/concert_platform?sslmode=disable"
	}

	var err error

	for i := 0; i < 10; i++ {

		DB, err = pgxpool.New(
			context.Background(),
			databaseURL,
		)

		if err == nil {

			err = DB.Ping(context.Background())

			if err == nil {

				log.Println(
					"PostgreSQL connected",
				)

				return
			}
		}

		log.Println(
			"Waiting for PostgreSQL...",
		)

		time.Sleep(3 * time.Second)
	}

	log.Fatal(
		"Database connection error:",
		err,
	)
}