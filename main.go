package main

import (
	"aleksei/go/db"
	"aleksei/go/router"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	var (
		APP_HOST = os.Getenv("APP_HOST")
		APP_PORT = os.Getenv("APP_PORT")
	)

	db.InitPostgresDB()
	router.InitRouter().Run(APP_HOST + ":" + APP_PORT)
}
