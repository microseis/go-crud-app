package main


import (
	"aleksei/go/db"
	"aleksei/go/router"
)
 

func main() {
    db.InitPostgresDB()
    router.InitRouter().Run("localhost:5000")
}