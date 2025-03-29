package main

import (
	"go-api/controller"
	"go-api/db"
	"go-api/repository"
	"go-api/usecase"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	dbConnection, err := db.ConnectDB()
	repository := repository.NewProductRepository(dbConnection)
	productUseCase := usecase.NewProductUseCase(repository)

	if err != nil {
		panic(err)
	}

	productController := controller.NewProductController(productUseCase)

	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	server.GET("/products", productController.GetProducts)
	server.POST("/products", productController.CreateProduct)
	server.GET("/products/:id", productController.GetProductById)
	server.DELETE("/products/:id", productController.DeleteProductById)

	server.Run(":8000")
}
