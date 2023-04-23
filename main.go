package main

import (
	"gin_sample/controllers"
	"github.com/gin-gonic/gin"
)

func setUpRouter() *gin.Engine {
	r := gin.Default()
	r.LoadHTMLGlob("resources/views/*")

	r.GET("/", controllers.ShowAllPost)
	r.GET("/posts/:id", controllers.ShowOnePost)

	return r
}

func main() {
	r := setUpRouter()
	r.Run(":8080")
}
