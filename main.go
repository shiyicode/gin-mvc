package main

import (
	"flag"

	"github.com/chuxinplan/gin-mvc/common/db"

	"github.com/chuxinplan/gin-mvc/common/config"
	"github.com/chuxinplan/gin-mvc/common/logger"
	_ "github.com/chuxinplan/gin-mvc/common/validator"
	"github.com/chuxinplan/gin-mvc/router"
)

func main() {
	configFile := flag.String("c", "conf/dev/app.conf.toml", "set config file")
	flag.Parse()

	config.Load(*configFile)
	logger.Init()

	db.Init()
	defer db.Close()

	router.Init()
	router.Run()
}
