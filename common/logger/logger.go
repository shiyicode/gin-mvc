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

	"github.com/chuxinplan/gin-mvc/common/config"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

var (
	Logger *logrus.Logger

	_levelMap = map[string]logrus.Level{
		"debug": logrus.DebugLevel,
		"info":  logrus.InfoLevel,
		"warn":  logrus.WarnLevel,
		"error": logrus.ErrorLevel,
	}
)

func Init() {
	Logger = logrus.New()

	Logger.SetReportCaller(true)

	formatter := &Formatter{}
	Logger.SetFormatter(formatter)

	conf := config.Get()
	if !conf.Log.Enable {
		logrus.Info("log to stdout")
		logrus.SetOutput(os.Stdout)
		logrus.SetLevel(logrus.DebugLevel)
		return
	}

	if level, ok := _levelMap[conf.Log.Level]; !ok {
		Logger.SetLevel(logrus.InfoLevel)
	} else {
		Logger.SetLevel(level)
	}

	err := os.MkdirAll(conf.Log.Path, 0777)
	if err != nil {
		log.Fatalf("create directory %s failure", conf.Log.Path)
	}

	Logger.AddHook(getAllHook(conf, formatter))
	Logger.AddHook(getWfHook(conf, formatter))
}

func GetRequestLogger(requestId string) *logrus.Entry {
	return Logger.WithField("reqId", requestId)
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
