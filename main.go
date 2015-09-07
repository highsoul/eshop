package main

import (
	"eshop/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {

	model.InitTool()

	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.Static("/assets", "./assets")
	r.GET("/", ShowIndex)
	r.GET("/user/register", UserRegisterHandler)
	r.POST("/user/register", UserRegisterHandler)
	r.Run(":8080")
}

func ShowIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func UserRegisterHandler(c *gin.Context) {
	if c.Request.Method == "GET" {
		c.HTML(http.StatusOK, "register.html", nil)
	} else if c.Request.Method == "POST" {
		u := model.User{Email: c.PostForm("email"), Name: c.PostForm("name"), Password: c.PostForm("password")}
		fmt.Println(u)
		u.InsertToDB()
		c.Redirect(http.StatusMovedPermanently, "/")
	}
}
