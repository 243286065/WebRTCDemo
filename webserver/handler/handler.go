package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//DefaultHomePageHandler: 处理默认首页
func DefaultHomePageHandler(c *gin.Context) {
	c.Redirect(http.StatusFound, "/static/index.html")
}

// func TLSHandler() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		secureMiddleware := secure.New(secure.Options{
// 			SSLRedirect: true,
// 			SSLHost:     "0.0.0.0:8081",
// 		})
// 		err := secureMiddleware.Process(c.Writer, c.Request)

// 		if err != nil {
// 			return
// 		}
// 		c.Next()
// 	}
// }
