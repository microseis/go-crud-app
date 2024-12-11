package main

import (
	"aleksei/go/db"
	"aleksei/go/router"
	"log"
	"os"

	"github.com/joho/godotenv"

)
// @title           Swagger API
// @version         1.0
// @description     This is a sample Open API


// @host      localhost:5000
// @BasePath  /api/v1

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  OpenAPI
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
