package main

import (
	"root/api/handler"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
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
	handler.Routes(router)

	router.Run()
}
