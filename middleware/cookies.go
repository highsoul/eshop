package middleware

import (
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
		s = i.(string)

	}
	return s
}

func (cm CookieManager) InitCM(c *gin.Context) {

	for _, cookie := range c.Request.Cookies() {
		cm.CookieMap[cookie.Name] = cookie.Value
	}
}

func (cm CookieManager) Add(name string, value interface{}) {
	c_value := ConvToString(value)
	expire := time.Now().AddDate(0, 0, 1)
	cookie := http.Cookie{Name: name, Value: c_value, Path: "/", Expires: expire, MaxAge: 86400}
	http.SetCookie(cm.Context.Writer, &cookie)
}

func (cm CookieManager) GetString(name string) string {
	return cm.CookieMap[name]
}

func (cm CookieManager) Get(name string) interface{} {
	return cm.CookieMap[name]
}

func Cookie() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookieManager := CookieManager{Context: c, CookieMap: make(map[string]string)}
		cookieManager.InitCM(c)
		c.Set("CM", cookieManager)
		c.Next()
		fmt.Println("Cookie is Working")
	}
}
