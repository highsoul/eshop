package model

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
)

var T Tool

type Tool struct {
	DB gorm.DB
}

func InitTool() {
	fmt.Println("hahahahhaha")
	T = Tool{}
	var err error
	T.DB, err = gorm.Open("mysql", "root:123456@/eshop?charset=utf8&parseTime=True")
	if err != nil {
		log.Fatalf("链接数据库出错哦 '%v'", err)
	}
	T.DB.AutoMigrate(&User{})
}
