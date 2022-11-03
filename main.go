package main

import (
	"root/api"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"

	_ "root/docs/ginsimple"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title LBC Test API
// @version 1.0
// @description A custom FizzBuzz API in Go using Gin Framework
// @termsOfService http://swagger.io/terms/

// @contact.name Maxime Marcuccilli
// @contact.url https://www.linkedin.com/in/maxime-marcuccilli/
// @contact.email maxime.marcuccilli@gmail.com

// @host localhost:8000
// @BasePath /
// @schemes http
func main() {

	router := gin.Default()

	router.GET("/fizzbuzz", api.FizzBuzzHandler)
	router.GET("/queries", api.QueriesHandler)

	url := ginSwagger.URL("http://localhost:8000/swagger/doc.json") // The url pointing to API definition
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	router.Run()
}
