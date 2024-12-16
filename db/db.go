package db

import (
	"database/sql"
	"errors"
	"fmt"

	"aleksei/go/utils"

	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

var (
	GORM_DB     *gorm.DB
	SQL_DB      *sql.DB
	DB_MIGRATOR gorm.Migrator
)


type Product struct {
	ID    string `json:"id" gorm:"primarykey"`
	Code  string `json:"code"`
	Price int32  `json:"price"`
}

// Инициализация базы данных
func InitPostgresDB(appConfig *utils.Config) error {

	var dbURL = appConfig.Database.Dsn
	
	db, err = gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err == nil {
        GORM_DB = db
        SQL_DB, _ = db.DB()
        DB_MIGRATOR = db.Migrator()
    }
    return err

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
