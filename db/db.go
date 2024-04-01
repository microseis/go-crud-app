package database

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)
var db *gorm.DB
var err error

type Product struct {
	ID          string `json:"id" gorm:"primarykey"`
	Code        string `json:"code"`
	Price string `json:"price"`
 }


func InitPostgresDB()  {

    err := godotenv.Load(".env")
    if err != nil{
     log.Fatalf("Error loading .env file: %s", err)
    }
   
    var (
        DB_USER     = os.Getenv("DB_USER")
        DB_PASSWORD = os.Getenv("DB_PASSWORD")
        DB_NAME     = os.Getenv("DB_NAME")
        DB_HOST     = os.Getenv("DB_HOST")
        DB_PORT     = os.Getenv("DB_PORT")
    )

    dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
        DB_HOST,
        DB_PORT,
        DB_USER,
        DB_NAME,
        DB_PASSWORD,
   )
    db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        panic("failed to connect database")
    }

  // Migrate the schema
  db.AutoMigrate(&Product{})
  fmt.Println("Table Product has been sucessfully migrated")

}

func CreateProduct(product *Product) (*Product, error) {
	fmt.Println(product)
	product.ID = uuid.New().String()
    res := db.Create(&product)
    if res.Error != nil {
        return nil, res.Error
    }
    return product, nil
 }


 func GetProduct(id string) (*Product, error) {
	var product Product
	res := db.First(&product, "id = ?", id)
	if res.RowsAffected == 0 {
	  return nil, errors.New(fmt.Sprintf("product of id %s not found", id))
	}
   return &product, nil
  }