package main

import (
	"golang-basic/config"
	"golang-basic/server"
)

func main() {
	server:= server.InitHTTPServer(config.NewModuleConfig())
	server.Run()
}
