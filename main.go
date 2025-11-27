package main

import (
	"log"

	"github.com/joho/godotenv"
	"backenduas_sistemprestasi/database"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("⚠️  .env not found, using system environment variables")
	}

	database.InitPostgres()
	database.InitMongo()
}