package server

import (
	"github.com/gin-gonic/gin"
	"golang-basic/config"
	"golang-basic/usecase"
	"strconv"
)

//httpServer structure as Server implementation for http api
type httpServer struct {
	conf *config.Config
}

func InitHTTPServer(conf *config.Config) httpServer {
	return httpServer{
		conf: conf,
	}
}

func (s httpServer) Run() {
	usecase, err := usecase.InitUsecase(s.conf)
	if err != nil {
		return
	}

	router := gin.Default()
	order := router.Group("/orders")
	{
		order.GET("/:id", usecase.GetOrder)
		order.POST("/", usecase.CreateOrder)
	}
	router.Run(":"+strconv.Itoa(s.conf.Main.Server.Port))
}
