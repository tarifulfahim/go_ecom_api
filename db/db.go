package db

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
)

func NewPgConnect() (*pgx.Conn, error) {
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	defer conn.Close(context.Background())

	return conn, err
}
