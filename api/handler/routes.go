package handler

import (
	_ "root/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Routes(router *gin.Engine) {
	router.GET("/fizzbuzz", FizzBuzzHandler)
	router.GET("/queries", QueriesHandler)

	url := ginSwagger.URL("http://localhost:8000/swagger/doc.json") // The url pointing to API definition
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
}
