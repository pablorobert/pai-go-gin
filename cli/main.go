package main

import (
	"go-api/controller"
	"go-api/usecase"
	"go-api/connectdb"
	"github.com/gin-gonic/gin"
)

func main(){


	server := gin.Default()

	dbConnection, err := connectdb.ConnectDb() 
	if(err != nil){
		panic(err)
	}

	ProductUsecase := usecase.NewProductUsecase()


	ProductController := conteroller.NewProductController(ProductUsecase)



	server.GET("/ping", func(ctx *gin.Context){
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})


	server.GET("/products", ProductController.GetProducts)

	server.Run(":9999")
}
