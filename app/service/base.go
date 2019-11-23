package service

import (
	"github.com/chuxinplan/gin-mvc/common/logger"
	"github.com/go-xorm/xorm"
	"github.com/sirupsen/logrus"
)

type baseService struct {
	Logger *logrus.Entry
	DB     *xorm.Session
}

func newBaseService(requestId string, db *xorm.Session) baseService {
	return baseService{
		Logger: logger.GetRequestLogger(requestId),
		DB:     db,
	}
}
