package controllers

import (
	"gin_sample/models"
	"github.com/gin-gonic/gin"
)

func ShowAllPost(c *gin.Context) {
	posts := models.Get()
	c.HTML(200, "index.tmpl", gin.H{"posts": posts})
}
