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

func (c *Category) GetAll() []Category {
	cates := []Category{}
	T.DB.Debug().Find(&cates)
	return cates
}

func (c *Category) GetList() []Category {
	cates := []Category{}
	T.DB.Debug().Where("top = ?", c.Top).Find(&cates)
	return cates
}

func (c Category) GetTop() Category {
	cate := Category{}
	T.DB.First(cate, c.Top)
	return cate
}
