package main

import (
	"WebRTCDemo/config"
	"WebRTCDemo/webserver/route"
)

func main() {
	router := route.Router()
	router.RunTLS(config.WebServerHostTLS, "ssl/server.crt", "ssl/server.key")
}
