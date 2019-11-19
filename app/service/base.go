package service

import (
	"github.com/chuxinplan/gin-mvc/common/logger"
	"github.com/sirupsen/logrus"
)

type baseService struct {
	Username string
	Logger   *logrus.Entry
	DB       interface{} // TODO
}

func newBaseService(username string, requestId string) baseService {
	return baseService{
		Username: username,
		Logger:   logger.GetRequestLogger(requestId),
		DB:       nil,
	}
}
