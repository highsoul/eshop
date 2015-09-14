package model

import (
	//"fmt"
	"github.com/jinzhu/gorm"
)

type Goods struct {
	gorm.Model
	Name        string
	Price       float32
	Description string
	Photo       string
}
