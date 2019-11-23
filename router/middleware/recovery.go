package middleware

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http/httputil"
	"runtime"

	log "github.com/sirupsen/logrus"

	"github.com/chuxinplan/gin-mvc/app/controller"
	"github.com/chuxinplan/gin-mvc/common/errors"
	"github.com/gin-gonic/gin"
)

var (
	_dunno     = []byte("???")
	_centerDot = []byte("·")
	_dot       = []byte(".")
	_slash     = []byte("/")
)

func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				httprequest, _ := httputil.DumpRequest(c.Request, false)
				stack := stack(3)
				if err, ok := err.(*errors.Err);ok{
					httpCode,respData := controller.Failure(err)
					c.JSON(httpCode,respData)
					errInfo := fmt.Sprintf("Err - errno: %d, message: %s, error: %s", err.Errno.Code, err.Errno.Message, err.InnerMsg)
					log.Warningf("[Recovery] panic recovered:\n%s\n%s\n%s\n",httprequest,errInfo,string(stack))
					return
				}
				resErr := errors.New(errors.InternalServerError, "")
				httpCode,respData := controller.Failure(resErr)
				c.JSON(httpCode,respData)
				log.Warningf("[Recovery] panic recovered:\nUnknown Error:\n%s\nError Mess:%s\n%s\n",httprequest,err,string(stack))
			}
		}()
		c.Next()
	}
}

// stack returns a nicely formated stack frame, skipping skip frames
func stack(skip int) []byte {
	buf := new(bytes.Buffer) // the returned data
	// As we loop, we open files and read them. These variables record the currently
	// loaded file.
	var lines [][]byte
	var lastFile string
	for i := skip; ; i++ { // Skip the expected number of frames
		pc, file, line, ok := runtime.Caller(i)
		if !ok {
			break
		}
		// Print this much at least.  If we can't find the source, it won't show.
		fmt.Fprintf(buf, "%s:%d (0x%x)\n", file, line, pc)
		if file != lastFile {
			data, err := ioutil.ReadFile(file)
			if err != nil {
				continue
			}
			lines = bytes.Split(data, []byte{'\n'})
			lastFile = file
		}
		fmt.Fprintf(buf, "\t%s: %s\n", function(pc), source(lines, line))
	}
	return buf.Bytes()
}

// source returns a space-trimmed slice of the n'th line.
func source(lines [][]byte, n int) []byte {
	n-- // in stack trace, lines are 1-indexed but our array is 0-indexed
	if n < 0 || n >= len(lines) {
		return _dunno
	}
	return bytes.TrimSpace(lines[n])
}

// function returns, if possible, the name of the function containing the PC.
func function(pc uintptr) []byte {
	fn := runtime.FuncForPC(pc)
	if fn == nil {
		return _dunno
	}
	name := []byte(fn.Name())
	// The name includes the path name to the package, which is unnecessary
	// since the file name is already included.  Plus, it has center dots.
	// That is, we see
	//	runtime/debug.*T·ptrmethod
	// and want
	//	*T.ptrmethod
	// Also the package path might contains dot (e.util. code.google.com/...),
	// so first eliminate the path prefix
	if lastslash := bytes.LastIndex(name, _slash); lastslash >= 0 {
		name = name[lastslash+1:]
	}
	if period := bytes.Index(name, _dot); period >= 0 {
		name = name[period+1:]
	}
	name = bytes.Replace(name, _centerDot, _dot, -1)
	return name
}
