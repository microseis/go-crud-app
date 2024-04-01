package db

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
	Price int32 `json:"price"`
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
	  return nil, fmt.Errorf("product of id %s not found", id)
	}
   return &product, nil
  }

  func GetProducts() ([]*Product, error) {
	var products []*Product
	res := db.Find(&products)
	if res.Error != nil {
		return nil, errors.New("no products found")
	}
	return products, nil
 }

 func UpdateProduct(product *Product) (*Product, error) {
    var productToUpdate Product
    result := db.Model(&productToUpdate).Where("id = ?", product.ID).Updates(product)
    if result.RowsAffected == 0 {
        return &productToUpdate, errors.New("product not updated")
    }
    return product, nil
 }

 func DeleteProduct(id string) error {
    var deletedProduct Product
    result := db.Where("id = ?", id).Delete(&deletedProduct)
    if result.RowsAffected == 0 {
        return errors.New("product not deleted")
    }
    return nil
 }