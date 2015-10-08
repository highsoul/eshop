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

func (c *Category) GetTopList() []Category {
	cates := []Category{}
	T.DB.Debug().Where("top = ?", c.Top).Find(&cates)
	return cates
}

func (c *Category) GetTopListNot() []Category {
	cates := []Category{}
	T.DB.Debug().Not("top", c.Top).Find(&cates)
	return cates
}

func (c *Category) GetSubList() []Category {
	cates := []Category{}
	T.DB.Debug().Where("top = ?", c.ID).Find(&cates)
	return cates
}

func (c Category) GetTop() Category {
	cate := Category{}
	T.DB.Debug().First(&cate, c.Top)
	return cate
}

func (c *Category) Get(id int) {
	T.DB.First(c, id)
}

func (c *Category) Update(id int) {
	top := c.Top
	name := c.Name
	T.DB.First(c, id).Updates(Category{Top: top, Name: name})
}

func (c *Category) Delete(id int) {
	T.DB.First(c, id)
	T.DB.Unscoped().Delete(c)
}
