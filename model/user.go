package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name     string `sql:"not null;unique"`
	Email    string `sql:"not null;unique"`
	Password string `sql:"not null"`
}

func (u *User) InsertToDB() bool {
	if T.DB.NewRecord(u) {
		T.DB.Create(&u)
		return true
	} else {
		fmt.Println("FALSE!")
		return false
	}
}
