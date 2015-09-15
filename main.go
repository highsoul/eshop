package main

import (
	"eshop/middleware"
	"eshop/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func main() {

	model.InitTool()

	r := gin.Default()
	authorized := r.Group("/admin", gin.BasicAuth(gin.Accounts{
		"admin": "123456",
	}))
	r.Use(middleware.Cookie())
	r.LoadHTMLGlob("templates/*")
	r.Static("/assets", "./assets")
	r.GET("/", ShowIndex)
	r.GET("/user/register", UserRegisterHandler)
	r.POST("/user/register", UserRegisterHandler)

	/* Admin */
	authorized.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "admin_index.html", nil)
	})
	authorized.GET("/goods", func(c *gin.Context) {
		cate := model.Category{Top: 0}
		cates := cate.GetListNot()
		goods := model.Goods{}
		goods_list := goods.GetAll()
		obj := make(map[string]interface{})
		obj["cates"] = cates
		obj["goods_list"] = goods_list
		c.HTML(http.StatusOK, "admin_goods.html", obj)
	})
	authorized.GET("/order", func(c *gin.Context) {
		c.HTML(http.StatusOK, "admin_order.html", nil)
	})

	authorized.GET("/category", func(c *gin.Context) {
		cate := model.Category{Top: 0}
		top_cates := cate.GetList()
		cates := cate.GetAll()
		obj := make(map[string]interface{})
		obj["top_cates"] = top_cates
		obj["cates"] = cates
		c.HTML(http.StatusOK, "admin_category.html", obj)
	})
	authorized.POST("/category/add", func(c *gin.Context) {
		top, _ := strconv.Atoi(c.PostForm("top"))
		name := c.PostForm("name")
		cate := model.Category{Name: name, Top: top}
		if cate.InsertToDB() {
			c.Redirect(http.StatusMovedPermanently, "/admin/category")
		}

	})
	authorized.POST("/category/modify", func(c *gin.Context) {
		top, _ := strconv.Atoi(c.PostForm("top"))
		name := c.PostForm("name")
		cate := model.Category{Name: name, Top: top}
		id, _ := strconv.Atoi(c.PostForm("id"))
		cate.Update(id)
		c.Redirect(http.StatusMovedPermanently, "/admin/category")

	})
	authorized.POST("/category/delete", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.PostForm("id"))
		cate := model.Category{}
		cate.Delete(id)
		c.JSON(http.StatusOK, gin.H{"message": "ok"})

	})
	authorized.GET("/category/:id", func(c *gin.Context) {
		cate_id, _ := strconv.Atoi(c.Param("id"))
		cate := model.Category{}
		cate.Get(cate_id)
		c.JSON(http.StatusOK, cate)
	})

	r.Run(":8080")
}

func ShowIndex(c *gin.Context) {

	c.HTML(http.StatusOK, "index.html", nil)

	cookie := c.MustGet("CM").(middleware.CookieManager)
	for k, v := range cookie.CookieMap {
		fmt.Println("Cookie [" + k + "]" + "  =  " + v)
	}
}

func UserRegisterHandler(c *gin.Context) {
	var cookie middleware.CookieManager
	if c.Request.Method == "GET" {
		c.HTML(http.StatusOK, "register.html", nil)
	} else if c.Request.Method == "POST" {
		u := model.User{Email: c.PostForm("email"), Name: c.PostForm("name"), Password: c.PostForm("password")}
		fmt.Println(u)

		if u.InsertToDB() {
			cookie = c.MustGet("CM").(middleware.CookieManager)
			cookie.Add("id", u.ID)
		}

		c.Redirect(http.StatusMovedPermanently, "/")
	}
}
