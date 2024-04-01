package main


import (
	"aleksei/go/db"
	"aleksei/go/router"
)
 

func main() {
    database.InitPostgresDB()
    router.InitRouter().Run("localhost:5000")
}