package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name     string
	Email    string
	Password string
}

func (u User) CreateTable() {
	T.DB.DropTableIfExists(&u)
	T.DB.CreateTable(&u)

}

func (u User) InsertToDB() {
	if T.DB.NewRecord(u) {
		T.DB.Create(&u)
	} else {
		fmt.Println("FALSE!")
	}
}
