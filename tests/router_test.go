package main

import (
	"aleksei/go/db"
	"aleksei/go/router"
	"aleksei/go/utils"
	"bytes"
	"encoding/json"
	"fmt"

	"net/http"
	"net/http/httptest"
	"testing"

	//"github.com/rs/xid"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// Инициализация роутера.
func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}

// Инициализация тестовой бд.
func SetDb() {
	var cfg utils.Config
	utils.ReadFile(&cfg)
	utils.ReadEnv(&cfg)
	fmt.Printf("%+v", cfg)

	db.InitPostgresDB(&cfg)
}

// Тест хендлера создания продукта.
func TestNewProductHandler(t *testing.T) {
	r := SetUpRouter()
	SetDb()
	r.POST("/api/v1/product", router.PostProduct)
	//productId := xid.New().String()
	product := db.Product{
		Price: 111,
		Code:  "1111",
	}
	jsonValue, _ := json.Marshal(product)
	req, _ := http.NewRequest("POST", "/api/v1/product", bytes.NewBuffer(jsonValue))

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)
}

// Тест хендлера получения всех продуктов.
func TestGetProductsHandler(t *testing.T) {
	SetDb()
	r := SetUpRouter()

	r.GET("/api/v1/products", router.GetProducts)
	req, err := http.NewRequest("GET", "/api/v1/products", nil)
	if err != nil {
		return
	}
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, w)
}
