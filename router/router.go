package router

import (
	"log"
	"time"

	"github.com/chuxinplan/gin-mvc/app/controller"
	"github.com/chuxinplan/gin-mvc/router/middleware"
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func Init() {
	router = gin.New()
	//gin.SetMode(g.Conf().Run.Mode)

	router.Use(middleware.MaxAllowed(10))

	v1Router := router.Group("v1/api")
	{
		v1Router.POST("login", controller.HttpHandlerLogin)
		v1Router.POST("register", controller.HttpHandlerLogin)
	}

	authV1Router := router.Group("v1/auth")
	{
		authV1Router.POST("login", controller.HttpHandlerLogin)
	}
}

// Run start http server
func Run() {
	endless.DefaultReadTimeOut = 10 * time.Second
	endless.DefaultWriteTimeOut = 10 * time.Second
	endless.DefaultMaxHeaderBytes = 1 << 20
	endless.DefaultHammerTime = 10 * time.Second

	server := endless.NewServer(":8080", router)
	err := server.ListenAndServe()
	if err != nil {
		log.Printf("server err: %v", err)
	}
}
