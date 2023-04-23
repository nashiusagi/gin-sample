package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "gin_sample/controllers"
)

func setUpRouter() *gin.Engine{
    r:=gin.Default()
    r.LoadHTMLGlob("resources/views/*")

    //r.GET("/posts", controllers.ShowAllPost)

    r.GET("/",func(c *gin.Context){
        c.HTML(http.StatusOK, "index.tmpl", gin.H{
            "title": "Gin HTML Example",
        })
    })

    r.GET("/show",controllers.ShowAllPost)
        
    return r
}

func main(){
    r:=setUpRouter()
    r.Run(":8080")
    //http.HandleFunc("/",models.Get)
    //if err := http.ListenAndServe(":8080", nil); err != nil {
    //    log.Fatal("ListenAndServe:", err)
    //}
}
