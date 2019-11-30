package router

import (
	"net/http"
	"time"

	"github.com/chuxinplan/gin-mvc/common/validator"
	"github.com/gin-gonic/gin/binding"

	"github.com/chuxinplan/gin-mvc/router/middleware"

	"github.com/chuxinplan/gin-mvc/app/controller"
	"github.com/chuxinplan/gin-mvc/common/config"
	"github.com/chuxinplan/gin-mvc/common/logger"
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func Init() {
	binding.Validator = validator.GetValidator()

	router = gin.New()
	gin.SetMode(config.Get().Run.Mode)

	router.Use(middleware.BindKey())
	router.Use(middleware.MaxAllowed(config.Get().Run.MaxAllowed))
	router.Use(middleware.Logger())
	router.Use(middleware.Recovery())

	router.GET("ping", controller.HttpHandlerPing)

	v1Router := router.Group("v1/api").Use(middleware.GetUser())
	{
		v1Router.POST("login", controller.HttpHandlerLogin)
		v1Router.POST("register", controller.HttpHandlerRegister)
	}

	authV1Router := router.Group("v1/auth").Use(middleware.MustGetUser())
	{
		authV1Router.POST("login", controller.HttpHandlerLogin)
		authV1Router.GET("groups", controller.HttpHandlerGetAllByUser)
	}
}

// Run start http server
func Run() {
	conf := config.Get()

	endless.DefaultReadTimeOut = conf.Endless.ReadTimeOut * time.Second
	endless.DefaultWriteTimeOut = conf.Endless.WriteTimeOut * time.Second
	endless.DefaultMaxHeaderBytes = 1 << conf.Endless.MaxHeaderBytes
	endless.DefaultHammerTime = conf.Endless.HammerTime * time.Second

	server := endless.NewServer(conf.Run.HTTPAddr, router)
	err := server.ListenAndServe()
	if err != nil {
		logger.Logger.Errorf("server err: ", err.Error())
	}
}

func ServeHTTP(w http.ResponseWriter, req *http.Request) {
	router.ServeHTTP(w, req)
}
