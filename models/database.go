package models

import (
    "fmt"
    "os"
    "time"

    "github.com/joho/godotenv"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
)

var Db *gorm.DB

func init(){
    err:=godotenv.Load(".env")
    if err!=nil{
        panic("cannot load .env")
    }
    user:=os.Getenv("DB_USER")
    pw:=os.Getenv("DB_PASSWORD")
    db_name:=os.Getenv("DB_NAME")
    var path string=fmt.Sprintf("%s:%s@tcp(localhost:3306)/%s?charset=utf8&parseTime=true",user,pw,db_name)
    dialector:=mysql.Open(path)

    if Db, err=gorm.Open(dialector);err!=nil{
        connect(dialector,10)
    }
    fmt.Println("db connected!!")
}

func connect (dialector gorm.Dialector, count int){
    var err error
    if Db, err=gorm.Open(dialector);err!=nil{
        if count>1{
            time.Sleep(time.Second*2)
            count--
            fmt.Printf("retry... count:%v\n",count)
            connect(dialector,count)
            return
        }
        panic(err.Error())
    }
}
