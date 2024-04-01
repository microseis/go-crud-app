package router

import (
	db "aleksei/go/db"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	r.POST("/product", postProduct)
	r.GET("/product/:id", getProduct)
	r.GET("/products", getProducts)
	r.PUT("/product/:id", putProduct)
	r.DELETE("/product/:id", deleteProduct)
	return r
}

func postProduct(ctx *gin.Context) {
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

func getProduct(ctx *gin.Context) {
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

func getProducts(ctx *gin.Context) {
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

func deleteProduct(ctx *gin.Context) {
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

func putProduct(ctx *gin.Context) {
	var updatedProduct db.Product
	err := ctx.Bind(&updatedProduct)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	id := ctx.Param("id")
	dbMovie, err := db.GetProduct(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	dbMovie.Price = updatedProduct.Price
	dbMovie.Code = updatedProduct.Code

	res, err := db.UpdateProduct(dbMovie)
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
