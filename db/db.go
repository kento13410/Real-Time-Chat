package db

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"os"
)

func Open() *sql.DB {
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
	fmt.Println(dsn)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("failed to init database: ", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("failed to connect database: ", err)
	}

	log.Default().Println("success to connect db!!")

	return db
}

func Close(db *sql.DB) {
	if err := db.Close(); err != nil {
		panic(err)
	}
}
