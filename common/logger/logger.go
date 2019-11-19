package logger

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"
	"time"

	"github.com/pkg/errors"

	"github.com/chuxinplan/gin-mvc/common/config"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

var _levelMap = map[string]logrus.Level{
	"debug": logrus.DebugLevel,
	"info":  logrus.InfoLevel,
	"warn":  logrus.WarnLevel,
	"error": logrus.ErrorLevel,
}

func Init() {
	conf := config.Get()
	if !conf.Log.Enable {
		logrus.Info("log to logger err")
		logrus.SetOutput(os.Stdout)
		logrus.SetLevel(logrus.DebugLevel)
		return
	}

	err := os.MkdirAll(conf.Log.Path, 0777)
	if err != nil {
		log.Fatalf("create directory %s failure", conf.Log.Path)
	}

	if level, ok := _levelMap[conf.Log.Level]; !ok {
		logger.SetLevel(logrus.InfoLevel)
	} else {
		logger.SetLevel(level)
	}
	//log.SetReportCaller(true)

	logger.SetFormatter(&logrus.TextFormatter{ForceColors: true})

	logPath := conf.Log.Path
	maxAge := time.Duration(conf.Log.MaxAge)
	rotateTime := time.Duration(conf.Log.RotateTime)
	fileName := "web.log"
	fmt.Printf("test [%s]", fileName)

	allHook := lfshook.NewHook(lfshook.WriterMap{
		logrus.DebugLevel: GetLogWriter(logPath, fileName, maxAge, rotateTime),
		logrus.InfoLevel:  GetLogWriter(logPath, fileName, maxAge, rotateTime),
		logrus.WarnLevel:  GetLogWriter(logPath, fileName, maxAge, rotateTime),
		logrus.ErrorLevel: GetLogWriter(logPath, fileName, maxAge, rotateTime),
		logrus.FatalLevel: GetLogWriter(logPath, fileName, maxAge, rotateTime),
		logrus.PanicLevel: GetLogWriter(logPath, fileName, maxAge, rotateTime),
	}, &logrus.TextFormatter{ForceColors: true})

	wfHook := lfshook.NewHook(lfshook.WriterMap{
		logrus.WarnLevel:  GetLogWriter(logPath, fileName+".wf", maxAge, rotateTime),
		logrus.ErrorLevel: GetLogWriter(logPath, fileName+".wf", maxAge, rotateTime),
		logrus.FatalLevel: GetLogWriter(logPath, fileName+".wf", maxAge, rotateTime),
		logrus.PanicLevel: GetLogWriter(logPath, fileName+".wf", maxAge, rotateTime),
	}, &logrus.TextFormatter{ForceColors: true})

	filenameHook := NewHook()
	filenameHook.Field = "line"

	logger.AddHook(allHook)
	logger.AddHook(wfHook)
	logger.AddHook(filenameHook)
}

func GetRequestLogger(requestId string) *logrus.Entry {
	return logger.WithField("reqId", requestId)
}

func GetLogWriter(logPath string, logType string, maxAge time.Duration, rotateTime time.Duration) io.Writer {
	logPath = path.Join(getCurrPath(), logPath, logType)

	writer, err := rotatelogs.New(
		logPath+".%Y%m%d",
		rotatelogs.WithLinkName(logPath),
		rotatelogs.WithMaxAge(maxAge*24*time.Hour),
		rotatelogs.WithRotationTime(rotateTime*24*time.Hour),
	)
	if err != nil {
		log.Fatal(errors.Wrap(err, "get log writer failed"))
	}
	return writer
}

func getCurrPath() string {
	file, _ := exec.LookPath(os.Args[0])
	logPath, _ := filepath.Abs(file)
	index := strings.LastIndex(logPath, string(os.PathSeparator))
	ret := logPath[:index]
	return ret
}
