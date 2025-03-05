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

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("no .env file found, please ask for it from jason it")
	}

	dsn := "postgres://postgres.qgtrgrvxdodhdanparqs:Felichar1927~@aws-0-us-west-1.pooler.supabase.com:6543/postgres"

	config, err := pgx.ParseConnectionString(dsn)
	if err != nil {
		fmt.Printf("not able to connect to DB\n")
		os.Exit(1)
	}

	// connect to server
	conn, err := pgx.Connect(config)
	if err != nil {
		log.Fatalf("not able to connect to DB: %v", err)
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
