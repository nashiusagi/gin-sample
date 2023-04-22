package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

func setUpRouter() *gin.Engine{
    r:=gin.Default()
    r.LoadHTMLGlob("resources/views/*")

    r.GET("/",func(c *gin.Context){
        c.HTML(http.StatusOK, "index.tmpl", gin.H{
            "title": "Gin HTML Example",
        })
    })

    return r
}

func main(){
    r:=setUpRouter()
    r.Run(":8080")
}
