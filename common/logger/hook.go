package logger

import (
	"time"

	"github.com/chuxinplan/gin-mvc/common/config"

	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

func getAllHook(conf config.Config, formatter logrus.Formatter) *lfshook.LfsHook {
	logPath := conf.Log.Path
	maxAge := time.Duration(conf.Log.MaxAge)
	rotateTime := time.Duration(conf.Log.RotateTime)
	fileName := "web.log"

	return lfshook.NewHook(lfshook.WriterMap{
		logrus.DebugLevel: GetLogWriter(logPath, fileName, maxAge, rotateTime),
		logrus.InfoLevel:  GetLogWriter(logPath, fileName, maxAge, rotateTime),
		logrus.WarnLevel:  GetLogWriter(logPath, fileName, maxAge, rotateTime),
		logrus.ErrorLevel: GetLogWriter(logPath, fileName, maxAge, rotateTime),
		logrus.FatalLevel: GetLogWriter(logPath, fileName, maxAge, rotateTime),
		logrus.PanicLevel: GetLogWriter(logPath, fileName, maxAge, rotateTime),
	}, formatter)
}

func getWfHook(conf config.Config, formatter logrus.Formatter) *lfshook.LfsHook {
	logPath := conf.Log.Path
	maxAge := time.Duration(conf.Log.MaxAge)
	rotateTime := time.Duration(conf.Log.RotateTime)
	fileName := "web.log"

	return lfshook.NewHook(lfshook.WriterMap{
		logrus.WarnLevel:  GetLogWriter(logPath, fileName+".wf", maxAge, rotateTime),
		logrus.ErrorLevel: GetLogWriter(logPath, fileName+".wf", maxAge, rotateTime),
		logrus.FatalLevel: GetLogWriter(logPath, fileName+".wf", maxAge, rotateTime),
		logrus.PanicLevel: GetLogWriter(logPath, fileName+".wf", maxAge, rotateTime),
	}, formatter)
}
