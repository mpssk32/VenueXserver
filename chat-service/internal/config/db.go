package config

import (
	"venuex/shared/logger"
	"context"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

var DB *pgxpool.Pool

func ConnectDB() {

	databaseURL := os.Getenv(
		"DATABASE_URL",
	)

	var err error

	DB, err = pgxpool.New(
		context.Background(),
		databaseURL,
	)

	if err != nil {

		logger.Log.Fatalw(
			"failed to create database pool",
			"error",
			err,
		)
	}

	err = DB.Ping(
		context.Background(),
	)

	if err != nil {

		logger.Log.Fatalw(
			"failed to connect to PostgreSQL",
			"error",
			err,
		)
	}

	logger.Log.Info(
		"PostgreSQL connected",
	)
}