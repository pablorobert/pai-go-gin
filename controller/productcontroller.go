package conteroller

import (
	"go-api/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type productController struct {
	procutUsecase usecase.ProductUsecase
}

func NewProductController(usecase usecase.ProductUsecase) productController {
	return productController{
		procutUsecase: usecase,
	}
}

func (p *productController) GetProducts(ctx *gin.Context) {
	products, err := p.procutUsecase.GetProducts()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}
	ctx.JSON(http.StatusOK, products)
}
