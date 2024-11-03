package main

import (
	"go-api/connectdb"
	"go-api/controller"
	"go-api/repository"
	"go-api/usecase"

	"github.com/gin-gonic/gin"
)

func main() {

	server := gin.Default()

	dbConnection, err := connectdb.ConnectDb()
	if err != nil {
		panic(err)
	}

	ProductRepository := repository.NewProductRepository(dbConnection)

	ProductUsecase := usecase.NewProductUsecase(ProductRepository)

	ProductController := conteroller.NewProductController(ProductUsecase)

	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	server.GET("/produtos", ProductController.GetProducts)
	server.POST("/produto", ProductController.CreateProduct)
	server.GET("/produtos/:productId", ProductController.GetProductById)

	server.Run(":9999")
}
