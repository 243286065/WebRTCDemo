package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//DefaultHomePageHandler: 处理默认首页
func DefaultHomePageHandler(c *gin.Context) {
	c.Redirect(http.StatusFound, "/static/index.html")
}
