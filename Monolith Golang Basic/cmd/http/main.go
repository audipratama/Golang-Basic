package main

import (
	"golang-basic/config"
	"golang-basic/server"
	"os"
)

func init() {
	os.Setenv("ENV","development")
}

func main() {
	server:= server.InitHTTPServer(config.NewModuleConfig())
	server.Run()
}
