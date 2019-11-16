package logger

import (
	"fmt"
	"io"
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
	log "github.com/sirupsen/logrus"
)

func init() {

}

var _levelMap = map[string]log.Level{
	"debug": log.DebugLevel,
	"info":  log.InfoLevel,
	"warn":  log.WarnLevel,
	"error": log.ErrorLevel,
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
		log.Fatalf("create directory %s failure", conf.Log.Path)
	}

	if level, ok := _levelMap[conf.Log.Level]; !ok {
		log.SetLevel(log.InfoLevel)
	} else {
		log.SetLevel(level)
	}
	//log.SetReportCaller(true)

	log.SetFormatter(&log.TextFormatter{ForceColors: true})

	filenameHook := NewHook()
	filenameHook.Field = "line"
	log.AddHook(filenameHook)

	logPath := conf.Log.Path
	maxAge := time.Duration(conf.Log.MaxAge)
	rotateTime := time.Duration(conf.Log.RotateTime)
	fileName := "web.log"
	fmt.Printf("test [%s]", fileName)

	allHook := lfshook.NewHook(lfshook.WriterMap{
		log.DebugLevel: GetLogWriter(logPath, fileName, maxAge, rotateTime),
		log.InfoLevel:  GetLogWriter(logPath, fileName, maxAge, rotateTime),
		log.WarnLevel:  GetLogWriter(logPath, fileName, maxAge, rotateTime),
		log.ErrorLevel: GetLogWriter(logPath, fileName, maxAge, rotateTime),
		log.FatalLevel: GetLogWriter(logPath, fileName, maxAge, rotateTime),
		log.PanicLevel: GetLogWriter(logPath, fileName, maxAge, rotateTime),
	}, &log.TextFormatter{ForceColors: true})
	log.AddHook(allHook)

	wfHook := lfshook.NewHook(lfshook.WriterMap{
		log.WarnLevel:  GetLogWriter(logPath, fileName+".wf", maxAge, rotateTime),
		log.ErrorLevel: GetLogWriter(logPath, fileName+".wf", maxAge, rotateTime),
		log.FatalLevel: GetLogWriter(logPath, fileName+".wf", maxAge, rotateTime),
		log.PanicLevel: GetLogWriter(logPath, fileName+".wf", maxAge, rotateTime),
	}, &log.TextFormatter{ForceColors: true})
	log.AddHook(wfHook)
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
