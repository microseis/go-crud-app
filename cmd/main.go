package main

import (
	"aleksei/go/db"
	"aleksei/go/router"
	"aleksei/go/utils"
	"fmt"
	"log"

	"github.com/pressly/goose/v3"
)

// @title           Swagger API
// @version         1.0
// @description     This is a sample Open API
// @host      localhost:5000
// @BasePath  /api/v1
// @securityDefinitions.basic  BasicAuth
// @externalDocs.description  OpenAPI

func main() {

	var cfg utils.Config
	utils.ReadFile(&cfg)
	utils.ReadEnv(&cfg)
	
	log.Printf("%+v", &cfg)

	db.InitPostgresDB(&cfg)

	if err := goose.SetDialect("postgres"); err != nil {
        panic(err)
    }
	// применение всех миграций goose
    if err := goose.Up(db.SQL_DB, "/app/migrations"); err != nil {
        panic(err)
    }
	log.Println("INFO: all migrations successfully applied to db")

	router.InitRouter().Run(fmt.Sprintf("%s:%s", cfg.App.Host, cfg.App.Port))
}
