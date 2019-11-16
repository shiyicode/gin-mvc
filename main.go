package main

import (
	"flag"

	"github.com/chuxinplan/gin-mvc/common/config"
	_ "github.com/chuxinplan/gin-mvc/common/validator"
	"github.com/chuxinplan/gin-mvc/router"
)

func main() {
	configFile := flag.String("c", "", "set config file")
	flag.Parse()

	config.Load(*configFile)

	router.Init()
	router.Run()
}
