package conteroller 


import (
	"go-api/entity"
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


func (p *productController) GetProducts(ctx *gin.Context){
	products := []entity.Product{
		{
			ID: 1,
			Name: "pizza",
			Price: 25.00,
		},
	}

	ctx.JSON(http.StatusOK, products)	
}
