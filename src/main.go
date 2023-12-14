package main

import (
	"W-chat/config"
	"W-chat/src/httpserver"
	"fmt"
)

func main() {
	fmt.Println("Server Start")
	conf := config.New()
	httpserver.Run(HttpServerInjector(conf))
	fmt.Println("Server End")
}
