package main

import (
	"eshop/middleware"
	"eshop/model"
	"eshop/tool"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"
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
	r.GET("/user/logout", UserLogout)
	r.POST("/user/register", UserRegisterHandler)
	r.GET("/user/login", UserLogin)
	r.POST("/user/login", UserLogin)

	/* Admin */
	authorized.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "admin_index.html", nil)
	})
	authorized.GET("/goods", func(c *gin.Context) {
		cate := model.Category{Top: 0}
		cates := cate.GetTopListNot()
		goods := model.Goods{}
		goods_list := goods.GetAll()
		obj := make(map[string]interface{})
		obj["cates"] = cates
		obj["goods_list"] = goods_list
		c.HTML(http.StatusOK, "admin_goods.html", obj)
	})
	authorized.POST("/goods/add", func(c *gin.Context) {
		file, _, err := c.Request.FormFile("photo")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()
		timestamp := fmt.Sprintf("%d", time.Now().Unix())
		saveFilePath := "./assets/data/goods/" + timestamp + ".png"
		photo, err := os.OpenFile(saveFilePath, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer photo.Close()
		io.Copy(photo, file)
		name := c.PostForm("name")
		price, _ := strconv.ParseFloat(c.PostForm("price"), 64)
		cate_id, _ := strconv.Atoi(c.PostForm("cate"))
		desc := c.PostForm("description")
		cate := model.Category{}
		cate.Get(cate_id)
		fmt.Println(cate)
		goods := model.Goods{Name: name, Price: price, Description: desc, Cate: cate, Photo: timestamp + ".png"}
		fmt.Println(goods)
		if goods.InsertToDB() {
			c.Redirect(http.StatusMovedPermanently, "/admin/goods")
		}
	})
	authorized.POST("/goods/modify", func(c *gin.Context) {

		m_id, _ := strconv.Atoi(c.PostForm("m-id"))
		g := model.Goods{}
		g.Get(m_id)
		name := c.PostForm("m-name")
		price, _ := strconv.ParseFloat(c.PostForm("m-price"), 64)
		cate_id, _ := strconv.Atoi(c.PostForm("m-cate"))
		desc := c.PostForm("m-description")
		cate := model.Category{}
		cate.Get(cate_id)
		g.Name = name
		g.Price = price
		g.Cate = cate
		g.Description = desc

		file, _, err := c.Request.FormFile("m-photo")
		if file != nil {
			defer file.Close()
			rm_err := os.Remove("./assets/data/goods/" + g.Photo)
			if rm_err != nil {
				fmt.Println(err)
			}
			timestamp := fmt.Sprintf("%d", time.Now().Unix())
			saveFilePath := "./assets/data/goods/" + timestamp + ".png"
			photo, err := os.OpenFile(saveFilePath, os.O_WRONLY|os.O_CREATE, 0666)
			if err != nil {
				fmt.Println(err)
				return
			}
			defer photo.Close()
			io.Copy(photo, file)
			g.Photo = timestamp + ".png"
		}

		g.Update()

		c.Redirect(http.StatusMovedPermanently, "/admin/goods")

	})

	authorized.POST("/goods/delete/", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.PostForm("id"))
		goods := model.Goods{}
		goods.Get(id)
		goods.Delete()
		c.JSON(http.StatusOK, gin.H{"message": "ok"})
	})

	authorized.GET("/goods/:id", func(c *gin.Context) {
		goods_id, _ := strconv.Atoi(c.Param("id"))
		goods := model.Goods{}
		goods.Get(goods_id)
		c.JSON(http.StatusOK, goods)
	})

	authorized.GET("/order", func(c *gin.Context) {
		c.HTML(http.StatusOK, "admin_order.html", nil)
	})

	authorized.GET("/category", func(c *gin.Context) {
		cate := model.Category{Top: 0}
		top_cates := cate.GetTopList()
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
	data := tool.GetTD(c)
	goods := model.Goods{}
	goods_list := goods.GetAll()
	data["goods_list"] = goods_list
	c.HTML(http.StatusOK, "index.html", data)
}

func UserRegisterHandler(c *gin.Context) {
	var cookie middleware.CookieManager
	if c.Request.Method == "GET" {
		c.HTML(http.StatusOK, "register.html", nil)
	} else if c.Request.Method == "POST" {
		u := model.User{Email: c.PostForm("email"), Name: c.PostForm("name"), Password: c.PostForm("password"), Phone: c.PostForm("phone"), Address: c.PostForm("address")}
		if u.InsertToDB() {
			cookie = c.MustGet("CM").(middleware.CookieManager)
			cookie.Add("user_id", u.ID)
			cookie.WriteCookies()
		}

		c.Redirect(http.StatusMovedPermanently, "/")
	}
}

func UserLogout(c *gin.Context) {
	var cookie middleware.CookieManager
	cookie = c.MustGet("CM").(middleware.CookieManager)
	cookie.Delete("user_id")
	cookie.WriteCookies()
	c.Redirect(http.StatusMovedPermanently, "/")
}

func UserLogin(c *gin.Context) {
	var cookie middleware.CookieManager
	cookie = c.MustGet("CM").(middleware.CookieManager)
	if c.Request.Method == "GET" {
		data := tool.GetTD(c)
		data["fail_msg"] = cookie.GetFlash("fail_msg")
		fmt.Println(data)
		c.HTML(http.StatusOK, "login.html", data)
	} else if c.Request.Method == "POST" {

		email := c.PostForm("email")
		password := c.PostForm("password")
		user := model.User{}
		model.T.DB.Debug().Where("email = ? and password = ?", email, password).First(&user)
		if user.Name != "" {
			cookie.Add("user_id", user.ID)
			cookie.WriteCookies()
			c.Redirect(http.StatusMovedPermanently, "/")
		} else {
			cookie.Flash("fail_msg", "login failed :(")
			c.Redirect(http.StatusMovedPermanently, "/user/login")
		}

	}
}
