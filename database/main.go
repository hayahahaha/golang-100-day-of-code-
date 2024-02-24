package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

func GetDatabaseUrl() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// urlExample := "postgres://username:password@localhost:5432/database_name"
	databaseName := os.Getenv("DATABASE_NAME")
	databaseUsername := os.Getenv("DATABASE_USERNAME")
	databasePassword := os.Getenv("DATABASE_PASSWORD")
	databaseHost := os.Getenv("DATABASE_HOST")

	return "postgres://" + databaseUsername + ":" + databasePassword + "@" + databaseHost + ":5432/" + databaseName
}

func main() {
	databaseUrl := GetDatabaseUrl()
	fmt.Println("database url : %s", databaseUrl)
	conn, err := pgx.Connect(context.Background(), databaseUrl)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	time := time.Now()
	err = conn.QueryRow(context.Background(), "SELECT  now();").Scan(&time)

	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(time)

}
