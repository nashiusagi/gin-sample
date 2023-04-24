package controllers

import (
	"fmt"
	"gin_sample/models"
	"github.com/gin-gonic/gin"
	"strconv"
)

func ShowAllPost(c *gin.Context) {
	posts := models.Get()
	c.HTML(200, "index.tmpl", gin.H{"posts": posts})
}

func ShowOnePost(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	post := models.GetOne(id)
	fmt.Println(post)
	c.HTML(200, "show.tmpl", gin.H{"post": post})
}

func Create(c *gin.Context) {
	c.HTML(200, "create.tmpl", gin.H{})
}

func StorePost(c *gin.Context){
    title:=c.PostForm("title")
    body:=c.PostForm("body")
    post:=models.Post{Title:title,Body:body}
    post.Create()
    c.Redirect(301,"/")
}
