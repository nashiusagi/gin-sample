package main

import (
	"gin_sample/controllers"
	"github.com/gin-gonic/gin"
)

func setUpRouter() *gin.Engine {
	r := gin.Default()
	r.LoadHTMLGlob("resources/views/*")

	r.GET("/", controllers.ShowAllPost)
	r.POST("/posts/store", controllers.StorePost)
	r.GET("/posts/:id", controllers.ShowOnePost)
	r.GET("/create", controllers.Create)

	return r
}

func main() {
	r := setUpRouter()
	r.Run(":8080")
}
