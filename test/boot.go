package test

import (
	"path"
	"runtime"

	"github.com/chuxinplan/gin-mvc/common/db"

	"github.com/chuxinplan/gin-mvc/common/config"
	"github.com/chuxinplan/gin-mvc/common/logger"
	"github.com/chuxinplan/gin-mvc/router"
)

func init() {
	configFile := getCurrentPath() + "/../conf/test/app.conf.toml"
	config.Load(configFile)
	logger.Init()

	db.Init()

	router.Init()
}

func getCurrentPath() string {
	_, filename, _, _ := runtime.Caller(1)

	return path.Dir(filename)
}
