package db

import (
	"errors"
	"fmt"

	"aleksei/go/utils"

	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

type Product struct {
	ID    string `json:"id" gorm:"primarykey"`
	Code  string `json:"code"`
	Price int32  `json:"price"`
}

// Инициализация базы данных
func InitPostgresDB(appConfig *utils.Config) {

	var (
		DB_USER     = appConfig.Database.Username
		DB_PASSWORD = appConfig.Database.Password
		DB_NAME     = appConfig.Database.Name
		DB_HOST     = appConfig.Database.Host
		DB_PORT     = appConfig.Database.Port
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
		panic("Failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&Product{})
	fmt.Println("Table Product has been successfully migrated")

}

// Создание продукта.
func CreateProduct(product *Product) (*Product, error) {
	product.ID = uuid.New().String()
	res := db.Create(&product)

	if res.Error != nil {
		return nil, res.Error
	}
	return product, nil
}

// Получение продукта по id.
func GetProduct(id string) (*Product, error) {
	var product Product
	res := db.First(&product, "id = ?", id)

	if res.RowsAffected == 0 {
		return nil, fmt.Errorf("product of id %s not found", id)
	}
	return &product, nil
}

// Получение продуктов.
func GetProducts() ([]*Product, error) {
	var products []*Product
	res := db.Find(&products)

	if res.Error != nil {
		return nil, errors.New("no products found")
	}

	return products, nil
}

// Обновление продукта.
func UpdateProduct(product *Product) (*Product, error) {
	var productToUpdate Product
	result := db.Model(&productToUpdate).Where("id = ?", product.ID).Updates(product)

	if result.RowsAffected == 0 {
		return &productToUpdate, errors.New("product not updated")
	}

	return product, nil
}

// Удаление продукта.
func DeleteProduct(id string) error {
	var deletedProduct Product
	result := db.Where("id = ?", id).Delete(&deletedProduct)

	if result.RowsAffected == 0 {
		return errors.New("product not deleted")
	}

	return nil
}
