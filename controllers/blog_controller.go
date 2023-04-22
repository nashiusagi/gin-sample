package controllers

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

func ShowAllBlog(c *gin.Context){
    datas:=model.Get()
    c.HTML(200,"show.tmpl",gin.H{"data":data})
