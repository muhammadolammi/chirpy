package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type apiConfig struct {
	FileserverHits int
	JWT_SECRET     string
	POLKA_API_KEY  string
}

func main() {
	// Define and parse the "debug" flag
	dbg := flag.Bool("debug", false, "Enable debug mode")
	flag.Parse()

	// Access the value through the pointer
	if *dbg {
		fmt.Println("Debug mode is enabled")
	} else {
		fmt.Println("Debug mode is disabled")
	}

	// by default, godotenv will look for a file named .env in the current directory
	godotenv.Load()
	jwtSecret := os.Getenv("JWT_SECRET")
	polkaApiKey := os.Getenv("POLKA_API_KEY")

	cfg := apiConfig{
		JWT_SECRET:    jwtSecret,
		POLKA_API_KEY: polkaApiKey,
	}
	serverEnrty(&cfg)
}
