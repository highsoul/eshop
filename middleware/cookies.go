package middleware

import (
	"eshop/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"reflect"
	"strconv"
	"time"
)

type CookieManager struct {
	Context   *gin.Context
	CookieMap map[string]string
	Cookies   []http.Cookie
}

func ConvToString(i interface{}) string {
	var s string
	itype := reflect.TypeOf(i).Name()
	fmt.Println(itype)
	switch itype {
	case "int":
		s = strconv.Itoa(i.(int))
	case "uint":
		s = fmt.Sprintf("%d", i.(uint))
	case "string":
		s = fmt.Sprintf("%s", i.(string))

	}
	return s
}

func (cm *CookieManager) InitCM(c *gin.Context) {

	for _, cookie := range c.Request.Cookies() {
		cm.CookieMap[cookie.Name] = cookie.Value
	}
}

func (cm *CookieManager) Add(name string, value interface{}) {
	c_value := ConvToString(value)
	expire := time.Now().AddDate(0, 0, 1)
	cookie := http.Cookie{Name: name, Value: c_value, Path: "/", Expires: expire, MaxAge: 86400}
	newCookies := append(cm.Cookies, cookie)
	cm.Cookies = newCookies
}
func (cm *CookieManager) Delete(name string) {
	cookie := http.Cookie{Name: name, Path: "/", MaxAge: -1}
	newCookies := append(cm.Cookies, cookie)
	cm.Cookies = newCookies
}
func (cm *CookieManager) WriteCookies() {
	for _, c := range cm.Cookies {
		http.SetCookie(cm.Context.Writer, &c)
	}
}
func (cm *CookieManager) GetString(name string) string {
	return cm.CookieMap[name]
}
func (cm *CookieManager) GetInt(name string) int {
	value, _ := strconv.Atoi(cm.CookieMap[name])
	return value
}
func (cm *CookieManager) Get(name string) interface{} {
	return cm.CookieMap[name]
}

func (cm *CookieManager) Flash(name string, value interface{}) {
	c_value := ConvToString(value)
	expire := time.Now().AddDate(0, 0, 1)
	cookie := http.Cookie{Name: name, Value: c_value, Path: "/", Expires: expire, MaxAge: 86400}
	http.SetCookie(cm.Context.Writer, &cookie)
}

func (cm *CookieManager) GetFlash(name string) string {
	flash, err := cm.Context.Request.Cookie(name)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	cookie := http.Cookie{Name: name, Path: "/", MaxAge: -1}
	http.SetCookie(cm.Context.Writer, &cookie)
	return flash.Value
}

func (cm *CookieManager) GetUser() {
	id_str := cm.GetString("user_id")
	if id_str != "" {
		user := model.User{}
		id, _ := strconv.Atoi(id_str)
		user.GetByID(id)
		cm.Context.Set("me", user)
	}

}

func Cookie() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookieManager := CookieManager{Context: c, CookieMap: make(map[string]string), Cookies: []http.Cookie{}}
		cookieManager.InitCM(c)
		c.Set("CM", cookieManager)
		cookieManager.GetUser()
		c.Next()

		fmt.Println("Cookie is Working")

	}
}
