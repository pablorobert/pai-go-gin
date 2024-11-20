package conteroller

import (
	"go-api/entity"
	"go-api/usecase"
	"net/http"
	"strconv"
	"strings"

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


func (p *productController) UpdateById (ctx *gin.Context) {
	// Obtém o ID do produto a partir dos parâmetros da URL
	id_product := ctx.Param("idProduct")
	
	// Verifica se o ID é vazio ou contém apenas espaços em branco
	if strings.TrimSpace(id_product) == ""{
		response := entity.Responde {
			MessageUpdate: "Erro: o ID nao foi encontrado",
		}

		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	// Converte o ID de string para inteiro 
	Idproduct, err := strconv.Atoi(id_product)
	if err != nil {
		response := entity.Responde {
			MessageUpdate: "Erro: o ID do produto precisa ser um numero",
		}

		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	// Chama o caso de uso para atualizar o produto 
	product, err := p.productUsecase.UpdateById(Idproduct)
	if err != nil {
		// Retorna erro genérico em caso de falha no caso de uso
		response := entity.Responde {
			MessageUpdate: "Erro ao atualizar o produto",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	// Verifica se o produto não foi encontrado 
	if product == nil {
		response := entity.Responde {
			MessageUpdate: "Produto nao encontrado na base de dados",
		}
		ctx.JSON(http.StatusNotFound, response)
		return
	}

	// Retorna o produto atualizado com sucesso
	ctx.JSON(http.StatusOK, product)
}
