package model

import (
	"database/sql"
	"github.com/jinzhu/gorm"
)

type Order struct {
	gorm.Model
	UserID  `sql:"index"`
	GoodsID `sql:"index"`
	Count   int
	Status  string
}

func (o Order) Create() {
	if T.DB.NewRecord(o) {
		T.DB.Create(&o)
	}
}
