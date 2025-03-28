package controller

import (
	"go-api/model"
	"go-api/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type controller struct {
	productUseCase usecase.ProductUseCase
}

func NewProductController(usecase usecase.ProductUseCase) controller {
	return controller{
		productUseCase: usecase,
	}
}

func (c *controller) GetProducts(ctx *gin.Context) {
	products := []model.Product{
		{ID: 1, Name: "Product 1", Price: 100},
		{ID: 2, Name: "Product 2", Price: 200},
		{ID: 3, Name: "Product 3", Price: 300},
	}

	ctx.JSON(http.StatusOK, products)
}
