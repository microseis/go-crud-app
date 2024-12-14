package main

import (
	"aleksei/go/db"
	"aleksei/go/router"
	"aleksei/go/utils"
	"fmt"

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
	
	fmt.Printf("%+v", &cfg)

	db.InitPostgresDB(&cfg)

	router.InitRouter().Run(fmt.Sprintf("%s:%s", cfg.App.Host, cfg.App.Port))
}
