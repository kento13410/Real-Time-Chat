package model

import (
	"Real-Time-Chat/ent"
	"entgo.io/ent/dialect"
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

func EntOpen() *ent.Client {
	if err := godotenv.Load(".env"); err != nil {
		panic(err)
	}
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)
	client, err := ent.Open(dialect.Postgres, dsn)
	if err != nil {
		panic(err)
	}
	return client
}

func EntClose(client *ent.Client) {
	if err := client.Close(); err != nil {
		panic(err)
	}
}

