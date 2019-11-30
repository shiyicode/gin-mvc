package db

import (
	"xorm.io/core"

	"github.com/chuxinplan/gin-mvc/common/config"
	"github.com/chuxinplan/gin-mvc/common/logger"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var DB *xorm.Engine

func Init() {
	conf := config.Get()

	var err error
	DB, err = xorm.NewEngine("mysql", conf.Mysql.WebAddr)

	if err != nil {
		logger.Logger.Fatalln("fail to connect mysql", conf.Mysql.WebAddr, err)
		return
	}

	DB.SetMaxIdleConns(conf.Mysql.MaxIdle)
	DB.SetMaxOpenConns(conf.Mysql.MaxOpen)

	if conf.Mysql.Debug {
		DB.ShowSQL(true)
		DB.ShowExecTime(true)

		DB.Logger().SetLevel(core.LOG_DEBUG)
	}
}

func Close() {
	conf := config.Get()

	err := DB.Close()
	if err != nil {
		logger.Logger.Errorf("fail to connect mysql", conf.Mysql.WebAddr, err)
	}
}
