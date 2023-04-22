package main

import (
    "strconv"
    "github.com/gin-gonic/gin"
    "net/http"
)

type User struct{
    Name string
    Age int
}

func setUpRouter() *gin.Engine{
    r:=gin.Default()
    r.LoadHTMLGlob("resources/views/*")

    r.GET("/user/:name/:age",func(c *gin.Context){
        name:=c.Param("name")
        age_str:=c.Param("age")
        age,_:=strconv.Atoi(age_str)

        user:=User{name,age}

        c.HTML(http.StatusOK, "user.tmpl", gin.H{
            "title": "Gin HTML Example",
            "user":user,
        })
    })

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
