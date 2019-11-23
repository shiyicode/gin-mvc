package logger

import (
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

	logger.SetReportCaller(true)

	formatter := &Formatter{}
	logger.SetFormatter(formatter)

	logger.AddHook(getAllHook(conf, formatter))
	logger.AddHook(getWfHook(conf, formatter))
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
