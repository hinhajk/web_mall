package main

import (
	"web_mall/config"
	"web_mall/router"
)

func main() {
	config.Init()
	r := router.NewRouter()
	r.Run(":3000")
}
