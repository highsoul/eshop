package model

import (
	//"fmt"
	"github.com/jinzhu/gorm"
)

type Category struct {
	gorm.Model
	Name string
}
