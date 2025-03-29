package controller

import (
	"go-api/model"
	"go-api/usecase"
	"net/http"
	"strconv"

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
	products, err := c.productUseCase.GetProducts()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, products)
}

func (c *controller) CreateProduct(ctx *gin.Context) {
	var product model.Product
	err := ctx.BindJSON(&product)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	productCreated, err := c.productUseCase.CreateProduct(product)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, productCreated)
}

func (c *controller) GetProductById(ctx *gin.Context) {
	id := ctx.Param("id")

	if id == "" {
		response := model.Response{
			Message: "Product ID is required",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	validID, err := strconv.Atoi(id)

	if err != nil {
		response := model.Response{
			Message: "Invalid Product ID",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	product, err := c.productUseCase.GetProductById(validID)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, product)
}
