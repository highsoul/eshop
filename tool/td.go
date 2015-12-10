/**
	Template Data
**/
package tool

import (
	"eshop/middleware"
	"eshop/model"
	"github.com/gin-gonic/gin"
)

type TD map[string]interface{}

func GetTD(c *gin.Context) TD {
	td := make(TD)

	var cookie middleware.CookieManager
	cookie = c.MustGet("CM").(middleware.CookieManager)

	me, exist := c.Get("me")
	if exist {
		td["me"] = me
	}
	cate := model.Category{Top: 0}
	cates := cate.GetTopList()
	td["top_cates"] = cates
	td["fail_msg"] = cookie.GetFlash("fail_msg")
	td["success_msg"] = cookie.GetFlash("success_msg")
	return td

}

func GetMe(c *gin.Context) model.User {
	user := model.User{}
	me, exist := c.Get("me")
	if exist {
		user = me.(model.User)
	}
	return user
}
