package router

import (
	"github.com/gin-gonic/gin"
	"sync"
	"github.com/chuxinplan/go-web-framework/common/middleware"
)

var router *gin.Engine
var once sync.Once

// 获取路由并初始化
func GetInstance() *gin.Engine {
	// 只执行一次
	once.Do(func() {
		initRouter()
	})
	return router
}

// 初始化路由
func initRouter() {
	router = gin.New()

	//gin.SetMode(g.Conf().Run.Mode)


	router.Use(middleware.MaxAllowed(10))

}
