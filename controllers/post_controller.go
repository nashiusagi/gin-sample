package controllers

import (
    "github.com/gin-gonic/gin"
    "gin_sample/models"
)

func ShowAllPost(c *gin.Context) {
    posts :=models.Get()
    c.HTML(200,"show.tmpl",gin.H{"posts":posts})
}
