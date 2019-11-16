package logger

import (
	"io"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"
	"time"

	"github.com/chuxinplan/gin-mvc/common/config"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	log "github.com/sirupsen/logrus"
)

func init() {

}

func Init() {
	conf := config.Get()
	if !conf.Log.Enable {
		log.Info("log to std err")
		log.SetOutput(os.Stdout)
		log.SetLevel(log.DebugLevel)
		return
	}

	err := os.MkdirAll(conf.Log.Path, 0777)
	if err != nil {
		log.Fatalf("create directory %s failure\n", conf.Log.Path)
	}

	log.SetLevel(log.InfoLevel)
	log.SetFormatter(&log.TextFormatter{ForceColors: true})

	logPath := conf.Log.Path
	maxAge := time.Duration(conf.Log.MaxAge)
	rotationTime := time.Duration(conf.Log.RotatTime)
	hook := lfshook.NewHook(lfshook.WriterMap{
		log.DebugLevel: GetLogWriter(logPath, "debug", maxAge, rotationTime),
		log.InfoLevel:  GetLogWriter(logPath, "info", maxAge, rotationTime),
		log.WarnLevel:  GetLogWriter(logPath, "warn", maxAge, rotationTime),
		log.ErrorLevel: GetLogWriter(logPath, "error", maxAge, rotationTime),
		log.FatalLevel: GetLogWriter(logPath, "fatal", maxAge, rotationTime),
		log.PanicLevel: GetLogWriter(logPath, "panic", maxAge, rotationTime),
	}, &log.TextFormatter{ForceColors: true})
	log.AddHook(hook)
}

func GetLogWriter(logPath string, logType string, maxAge time.Duration, rotationTime time.Duration) io.Writer {
	logPath = path.Join(getCurrPath(), logPath, logType, logType)

	writer, err := rotatelogs.New(
		logPath+".%Y%m%d.log",
		rotatelogs.WithLinkName(logPath),
		rotatelogs.WithMaxAge(maxAge*24*time.Hour),
		rotatelogs.WithRotationTime(rotationTime*24*time.Hour),
	)
	if err != nil {
		return nil
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
