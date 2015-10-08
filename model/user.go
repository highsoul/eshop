package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name     string `sql:"not null"`
	Email    string `sql:"not null;unique"`
	Password string `sql:"not null"`
	Address  string `sql:"not null"`
	Phone    string `sql:"not null"`
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

func (u *User) GetByID(id int) {
	T.DB.Debug().First(u, id)
}
