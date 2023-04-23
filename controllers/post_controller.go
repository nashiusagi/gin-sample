package controllers

import (
    //"github.com/gin-gonic/gin"
    "gin_sample/models"
)

func ShowAllPost() (datas []models.Post){
    posts:=models.Post.Get()
    return posts
    //fmt.Println(blogs)
    //c.HTML(200,"show.tmpl",gin.H{"blogs":blogs})
}
