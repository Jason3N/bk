package config

import (
	"os"
)

type Config struct {
	SERVER_ADDRESS string
	API_KEY        string
	PROJECT_URL    string
}

// start up supabase
func LoadSupabase(*Config) {

}

// get the .env from the file directory on computer
func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
