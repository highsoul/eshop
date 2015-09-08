package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type CookieManager struct {
	Context   *gin.Context
	CookieMap map[string]string
}

func (cm CookieManager) InitCM(c *gin.Context) {

	for _, cookie := range c.Request.Cookies() {
		cm.CookieMap[cookie.Name] = cookie.Value
	}
}

func (cm CookieManager) Add(name string, value string) {
	expire := time.Now().AddDate(0, 0, 1)
	cookie := http.Cookie{Name: name, Value: value, Path: "/", Expires: expire, MaxAge: 86400}
	http.SetCookie(cm.Context.Writer, &cookie)
}

func (cm CookieManager) GetString(name string) string {
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
