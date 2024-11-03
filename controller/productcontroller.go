package conteroller

import (
	"go-api/entity"
	"go-api/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type productController struct {
	productUsecase usecase.ProductUsecase
}

func NewProductController(usecase usecase.ProductUsecase) productController {
	return productController{
		productUsecase: usecase,
	}
}

func (p *productController) GetProducts(ctx *gin.Context) {
	products, err := p.productUsecase.GetProducts()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}
	ctx.JSON(http.StatusOK, products)
}

func (p *productController) CreateProduct(ctx *gin.Context)  {
	var product entity.Product
	err := ctx.BindJSON(&product)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	insertProduct, err := p.productUsecase.CreateProduct(product)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, insertProduct)

}


func (p *productController) GetProductById(ctx *gin.Context) {
	id_product := ctx.Param("productId")

	if(id_product == ""){
		responde := entity.Responde{
			Message: "Id do produto nao pode ser vazio!",
		}
		ctx.JSON(http.StatusBadRequest, responde)
		return
	}

	productId, err := strconv.Atoi(id_product)

	if(err != nil) {
		responde := entity.Responde {
			Message: "Id do produto tem que ser um valor inteiro",
		}
		ctx.JSON(http.StatusBadRequest, responde)
		return
	}
		
	product, err := p.productUsecase.GetProductById(productId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}

	if product == nil {
		reponse := entity.Responde{
			Message: "Produto nao foi encontrado na base de dados",
		}
		ctx.JSON(http.StatusNotFound, reponse)
		return
	}

	ctx.JSON(http.StatusOK, product)
}
