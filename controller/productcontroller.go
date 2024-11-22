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

func (p *productController) CreateProduct(ctx *gin.Context) {
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

	if id_product == "" {
		responde := entity.Response{
			Message: "Id do produto nao pode ser vazio!",
		}
		ctx.JSON(http.StatusBadRequest, responde)
		return
	}

	productId, err := strconv.Atoi(id_product)

	if err != nil {
		responde := entity.Response{
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
		reponse := entity.Response{
			Message: "Produto nao foi encontrado na base de dados",
		}
		ctx.JSON(http.StatusNotFound, reponse)
		return
	}

	ctx.JSON(http.StatusOK, product)
}

func (p *productController) UpdateProduct(ctx *gin.Context) {
	// Obtém o ID do produto a partir dos parâmetros da URL
	id_product := ctx.Param("idProduct")

	// Verifica se o ID é vazio ou contém apenas espaços em branco
	if strings.TrimSpace(id_product) == "" {
		response := entity.Response{
			MessageUpdate: "Erro: o ID nao foi encontrado",
		}

		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	// Converte o ID de string para inteiro
	Idproduct, err := strconv.Atoi(id_product)
	if err != nil {
		response := entity.Response{
			MessageUpdate: "Erro: o ID do produto precisa ser um numero",
		}

		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	// Chama o caso de uso para atualizar o produto
	product, err := p.productUsecase.GetProductById(Idproduct)
	// Verifica se o produto não foi encontrado
	if err != nil {
		response := entity.Response{
			MessageUpdate: "Produto nao encontrado na base de dados",
		}
		ctx.JSON(http.StatusNotFound, response)
		return
	}

	var body entity.Product
	err = ctx.ShouldBindJSON(&body)
	if err != nil {
		response := entity.Response{
			MessageUpdate: "Erro ao atualizar o produto",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	product.Name = body.Name
	product.Price = body.Price

	p.productUsecase.UpdateProduct(*product)

	// Retorna o produto atualizado com sucesso
	ctx.JSON(http.StatusOK, product)
}

func (p *productController) DeleteProduct(ctx *gin.Context) {
	id_product := ctx.Param("idProduct")

	// Verifica se o ID é vazio ou contém apenas espaços em branco
	if strings.TrimSpace(id_product) == "" {
		response := entity.Response{
			Message: "Erro: o ID nao foi encontrado",
		}

		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	// Converte o ID de string para inteiro
	Idproduct, err := strconv.Atoi(id_product)
	if err != nil {
		response := entity.Response{
			Message: "Erro: o ID do produto precisa ser um numero",
		}

		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	_, err = p.productUsecase.GetProductById(Idproduct)
	// Verifica se o produto não foi encontrado
	if err != nil {
		response := entity.Response{
			Message: "Produto nao encontrado na base de dados",
		}
		ctx.JSON(http.StatusNotFound, response)
		return
	}

	// Chama o caso de uso para atualizar o produto
	err = p.productUsecase.DeleteProductById(Idproduct)
	// Verifica se o produto não foi encontrado
	if err != nil {
		response := entity.Response{
			Message: "Erro ao deletar produto de Id " + id_product,
		}
		ctx.JSON(http.StatusNotFound, response)
		return
	}

}
