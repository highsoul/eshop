package main

import (
	"eshop/middleware"
	"eshop/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
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
	authorized.GET("/goods", func(c *gin.Context) {
		c.HTML(http.StatusOK, "admin_goods.html", nil)
	})
	authorized.GET("/category", func(c *gin.Context) {
		c.HTML(http.StatusOK, "admin_category.html", nil)
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
