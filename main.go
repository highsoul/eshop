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
	r.Use(middleware.Cookie())
	r.LoadHTMLGlob("templates/*")
	r.Static("/assets", "./assets")
	r.GET("/", ShowIndex)
	r.GET("/user/register", UserRegisterHandler)
	r.POST("/user/register", UserRegisterHandler)
	r.Run(":8080")
}

func ShowIndex(c *gin.Context) {
	cv, err := c.Request.Cookie("name")
	if err != nil {
		fmt.Println(err.Error())
	} else {

		if cv != nil {
			fmt.Println("Cookie is " + cv.Value)
		} else {
			fmt.Println("Cookie is nil :(")
		}

	}
	c.HTML(http.StatusOK, "index.html", nil)
}

func UserRegisterHandler(c *gin.Context) {
	if c.Request.Method == "GET" {
		c.HTML(http.StatusOK, "register.html", nil)
	} else if c.Request.Method == "POST" {
		u := model.User{Email: c.PostForm("email"), Name: c.PostForm("name"), Password: c.PostForm("password")}
		fmt.Println(u)
		if u.InsertToDB() {
			cookie := c.MustGet("CM")
			cookie.Add("name", u.Name)
		}
		c.Redirect(http.StatusMovedPermanently, "/")
	}
}
