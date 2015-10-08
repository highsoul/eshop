/**
	Template Data
**/
package tool

import (
	"eshop/model"
	"github.com/gin-gonic/gin"
)

type TD map[string]interface{}

func GetTD(c *gin.Context) TD {
	td := make(TD)
	me, exist := c.Get("me")
	if exist {
		td["me"] = me
	}
	cate := model.Category{Top: 0}
	cates := cate.GetTopList()
	td["top_cates"] = cates
	return td

}
