package model

import (
	"github.com/jinzhu/gorm"
)

type Order struct {
	gorm.Model
	UserID    int `sql:"index"`
	GoodsID   int `sql:"index"`
	AddressID int `sql:"index"`
	Count     int
	Status    string
}

func (o Order) Create() {
	if T.DB.NewRecord(o) {
		T.DB.Create(&o)
	}
}
