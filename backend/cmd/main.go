package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Jason3n/bk/internal/handlers"
	"github.com/Jason3n/bk/internal/repository"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx"
	"github.com/joho/godotenv"
)

// start up the connection to the db
func main() {

	if err := godotenv.Load(); err != nil {
		log.Println("no .env file found, please insert it")
	}
	// spin up server
	/* API_URL := os.Getenv("PROJECT_URL")
	API_KEY := os.Getenv("API_KEY") */
	CONNECTION_STRING := os.Getenv("CONNECTION_STRING")

	config, err := pgx.ParseConnectionString(CONNECTION_STRING)
	if err != nil {
		fmt.Printf("not able to connect to DB\n")
		os.Exit(1)
	}
	// CONNECT TO PSQL DB ON SUPABASE
	conn, err := pgx.Connect(config)
	if err != nil {
		log.Fatalf("Not able to connect to DB: %v", err)
	}

	// wait to close server until ctrl + c
	defer conn.Close()

	// intialize handler/repo
	userRepo := repository.UserRepository(conn)
	userHandler := handlers.UserHandler(userRepo)

	// create new router
	r := mux.NewRouter()
	r.HandleFunc("/addUser", userHandler.CreateUser).Methods("POST")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("back it up")
		return
	}
}
