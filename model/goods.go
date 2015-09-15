package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

type Goods struct {
	gorm.Model
	Name        string
	Price       float32
	Description string
	Photo       string
	Cate        Category
}

func (g *Goods) InsertToDB() bool {
	if T.DB.NewRecord(g) {
		T.DB.Create(&g)
		return true
	} else {
		fmt.Println("FALSE!")
		return false
	}
}

func (g *Goods) GetAll() []Goods {
	goods_list := []Goods{}
	T.DB.Debug().Find(&goods_list)
	return goods_list
}
