package models

import (
    "fmt"
    "net/http"
)

type Post struct{
    Id uint `gorm:"primaryKey"`
    Title string
    Body string
}

func init(){
    Db.Set("gorm:table_options", "ENGINE = InnoDB").AutoMigrate(Post{})
}

func Get(w http.ResponseWriter, r *http.Request) {
    var allPost []Post
    result:=Db.Find(&allPost)
    if result.Error!=nil{
        panic(result.Error)
    }

    fmt.Println(allPost)
    return
}

func (data Post) Create (){
    result:=Db.Create(data)
    if result.Error!=nil{
        panic(result.Error)
    }
    return
}

