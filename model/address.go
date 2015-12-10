package model

import (
	"github.com/jinzhu/gorm"
)

type Address struct {
	gorm.Model
	Province string
	City     string
	Detail   string
	Phone    string
	Name     string
	UserID   int `sql:"index"`
}

func (a Address) Create() {
	if T.DB.NewRecord(a) {
		T.DB.Create(&a)
	}
}
func (a *Address) GetAll() []Address {
	address_list := []Address{}
	T.DB.Find(&address_list)
	return address_list
}
func (a *Address) Delete() {
	T.DB.Debug().Unscoped().Delete(a)
}
func (a *Address) Get(id int) {
	T.DB.Debug().First(a, id)
}
