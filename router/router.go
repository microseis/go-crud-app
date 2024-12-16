package router

import (
	db "aleksei/go/db"
	docs "aleksei/go/docs"
	"net/http"

	"github.com/gin-gonic/gin"
	swagger_files "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// инициализация роутера
func InitRouter() *gin.Engine {

	r := gin.Default()
	r.SetTrustedProxies([]string{"localhost"})
	// добавление в роутер сваггера
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swagger_files.Handler))

	docs.SwaggerInfo.BasePath = "/api/v1"

	v1 := r.Group("/api/v1")
	{
		v1.POST("/product", PostProduct)
		v1.GET("/product/:id", GetProduct)
		v1.GET("/products", GetProducts)
		v1.PUT("/product/:id", PutProduct)
		v1.DELETE("/product/:id", DeleteProduct)
	}

	return r
}

// @Summary      Create Product
// @Description  create a product
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        Code    query     string  true  "product code"
// @Param        Price    query     string  true  "product price"
// @Success      200  {object}  db.Product
// @Router       /product [post]

// Хендлер создания продукта
func PostProduct(ctx *gin.Context) {
	var product db.Product
	err := ctx.Bind(&product)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	res, err := db.CreateProduct(&product)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"product": res,
	})
}

// @Summary      Get Product
// @Description  get product by ID
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Product ID"
// @Success      200  {object}  db.Product
// @Router       /product/{id} [get]

// Хендлер получения продукта по id
func GetProduct(ctx *gin.Context) {
	id := ctx.Param("id")
	res, err := db.GetProduct(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"product": res,
	})
}

// @Summary      Get Products
// @Description  get all products
// @Tags         products
// @Accept       json
// @Produce      json
// @Success      200  {object}  db.Product
// @Router       /products [get]

func GetProducts(ctx *gin.Context) {
	res, err := db.GetProducts()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"products": res,
	})
}

// @Summary      Delete Product
// @Description  delete product by ID
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Product ID"
// @Success      200  {object}  db.Product
// @Router       /product/{id} [delete]

func DeleteProduct(ctx *gin.Context) {
	id := ctx.Param("id")
	err := db.DeleteProduct(id)

	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "product deleted successfully",
	})
}

// Обновление продукта по id
// @Summary      Update Product
// @Description  update a product by id
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        id    path     string  true  "product id"
// @Param        Price    query     string  true  "product price"
// @Param        Code    query     string  true  "product code"
// @Success      200  {object}  db.Product
// @Router       /product/{id} [put]

func PutProduct(ctx *gin.Context) {
	var updatedProduct db.Product
	err := ctx.Bind(&updatedProduct)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	id := ctx.Param("id")
	product, err := db.GetProduct(id)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	product.Price = updatedProduct.Price
	product.Code = updatedProduct.Code

	res, err := db.UpdateProduct(product)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"product": res,
	})
}
