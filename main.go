package main

import (
	"flag"

	"github.com/chuxinplan/go-web-framework/common/conf"
	"github.com/chuxinplan/go-web-framework/common/db"
	"github.com/chuxinplan/go-web-framework/common/log"
	"github.com/chuxinplan/go-web-framework/router"
)

func main() {
	confPath := flag.String("conf", "conf/app.conf.toml", "set config file")
	flag.Parse()

	Init(*confPath)

	router := router.GetInstance()
	router.Run()
}

func Init(confPath string) {
	conf.Init(confPath)

	db.Init()
	defer db.Close()

	log.Init()
	defer log.Close()
}
