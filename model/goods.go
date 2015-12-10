package model

import (
	"database/sql"
	"fmt"
	"github.com/jinzhu/gorm"
)

type Goods struct {
	gorm.Model
	Name        string
	Price       float64
	Description string
	Photo       string
	Cate        Category
	CateID      sql.NullInt64
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

func (g *Goods) Get(id int) {
	T.DB.First(g, id)
}

func (g *Goods) GetAll() []Goods {
	goods_list := []Goods{}
	T.DB.Find(&goods_list)
	return goods_list
}

func (g Goods) GetCate() Category {
	cate := Category{}
	T.DB.Model(&g).Related(&cate, "CateID")
	return cate
}

func (g *Goods) Update() {

	T.DB.Debug().Save(g)
}

func (g *Goods) Delete() {
	T.DB.Unscoped().Delete(&g)
}
