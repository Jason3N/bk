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

// Start up the connection to the database
func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, please insert it")
	}

	CONNECTION_STRING := os.Getenv("CONNECTION_STRING")

	// Parse connection config
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
	userRepo := repository.NewUserRepository(conn)
	userHandler := handlers.NewUserHandler(userRepo)

	// create new router
	r := mux.NewRouter()
	r.HandleFunc("/addUser", userHandler.CreateUser).Methods("POST")
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("working"))
	}).Methods("GET")

	if err := http.ListenAndServe(":8080", r); err != nil {
		fmt.Printf("back it up")
		return
	}
}
