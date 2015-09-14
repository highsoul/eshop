package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

type Category struct {
	gorm.Model
	Name string
	Top  int
}

func (c *Category) InsertToDB() bool {
	if T.DB.NewRecord(c) {
		T.DB.Create(&c)
		return true
	} else {
		fmt.Println("FALSE!")
		return false
	}
}

func GetTop(top int) []Category {
	cates := []Category{}
	T.DB.Debug().Where(&cates, Category{Top: top}).Find(&cates)
	return cates
}
