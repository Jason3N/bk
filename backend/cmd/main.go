package main

import (
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx"
	"github.com/joho/godotenv"
)

// start up the connection to the db
func main() {

	if err := godotenv.Load(); err != nil {
		log.Println("no .env file found, please insert it")
	}
	// API_URL := os.Getenv("PROJECT_URL")
	// API_KEY := os.Getenv("API_KEY")
	CONNECTION_STRING := os.Getenv("CONNECTION_STRING")

	config, err := pgx.ParseConnectionString(CONNECTION_STRING)
	if err != nil {
		fmt.Printf("not able to connect to DB\n")
		os.Exit(1)
	}
	conn, err := pgx.Connect(config)
	if err != nil {
		log.Fatalf("Not able to connect to DB: %v", err)
	}

	defer conn.Close()

	rows, err := conn.Query("SELECT * FROM user")
	if err != nil {
		log.Fatalf("Query failed: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var user string

		if err := rows.Scan(&user); err != nil {
			log.Fatalf("Error scanning row: %v", err)
		}

		fmt.Printf(user + "\n")
	}

	// Check for errors encountered during iteration.
	if err = rows.Err(); err != nil {
		log.Fatalf("Error iterating rows: %v", err)
	}

}
