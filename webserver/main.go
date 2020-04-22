package main

import (
	"WebRTCDemo/config"
	"WebRTCDemo/webserver/route"
)

func main() {
	router := route.Router()
	router.Run(config.WebServerHost)
}
