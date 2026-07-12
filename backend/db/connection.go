package db

import (
	"context"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

func Connect() (*pgx.Conn, error) {
	err := godotenv.Load(".env")

	if err != nil {
		return nil, err
	}

	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	url := "postgres://" + user + ":" + password + "@" + host + ":" + port + "/" + dbName

	conn, err := pgx.Connect(context.Background(), url)

	if err != nil {
		return nil, err
	}

	return conn, nil
}
