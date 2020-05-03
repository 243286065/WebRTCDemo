package route

import (
	"WebRTCDemo/webserver/handler"

	"github.com/gin-gonic/gin"
)

// Router 返回路由器
func Router() *gin.Engine {
	//初始化
	gin.SetMode(gin.ReleaseMode) //全局设置环境，此为开发环境，线上环境为gin.ReleaseMode
	router := gin.Default()

	// router.Use(handler.TLSHandler())
	//处理静态资源
	router.Static("/static/", "./static")

	//处理默认首页
	router.GET("/", handler.DefaultHomePageHandler)

	//处理socketio请求
	router.GET("/socket.io/", handler.SocketIOServerHandler)
	router.POST("/socket.io/", handler.SocketIOServerHandler)
	router.Handle("WS", "/socket.io", handler.SocketIOServerHandler)
	router.Handle("WSS", "/socket.io", handler.SocketIOServerHandler)

	return router
}
