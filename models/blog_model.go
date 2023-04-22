package models

import (
    "gorm.io/gorm"
)

type Blog struct{
    Id uint `gorm:"primaryKey"`
    Title string
    Body string
}

func Get() (datas []Blog){
    result:=Db.Find(&datas)
    if result.Error!=nil{
        panic(result.Error)
    }
    return
}

func (data Blog) Create (){
    result:=Db.Create(datas)
    if result.Error!=nil{
        panic(result.Error)
    }
    return
}

