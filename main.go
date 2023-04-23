package main

import (
    "net/http"
    "log"
    //"gin_sample/controllers"
    "gin_sample/models"
)

//func setUpRouter() *gin.Engine{
//    r:=gin.Default()
//    r.LoadHTMLGlob("resources/views/*")
//
//    //r.GET("/posts", controllers.ShowAllPost)
//
//    //r.GET("/",func(c *gin.Context){
//    r.GET("/",
//        c.HTML(http.StatusOK, "index.tmpl", gin.H{
//            "title": "Gin HTML Example",
//        }))
//    }
//    return 
//}

func main(){
    //r:=setUpRouter()
    //r.Run(":8080")
    http.HandleFunc("/",models.Get)
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal("ListenAndServe:", err)
    }
}
