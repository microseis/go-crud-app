package router

import (
	database "aleksei/go/db"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)
 
func InitRouter() *gin.Engine {
   r := gin.Default()
   r.POST("/product", postProduct)
   r.GET("/product/:id", getProduct)
   return r
}

func postProduct(ctx *gin.Context) {
	var product database.Product
	err := ctx.Bind(&product)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	fmt.Println(product)
	res, err := database.CreateProduct(&product)
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
	res, err := database.GetProduct(id)
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
