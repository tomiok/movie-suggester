package logs

import (
	"log"
	"os"
	"path"
	"runtime"
)

var (
	infoLogger  = log.New(os.Stdout, "", 0)
	errorLogger = log.New(os.Stdout, "", log.Lshortfile|log.Ltime)
)

func Info(s interface{}) {
	pc, file, line, ok := runtime.Caller(1)
	if ok {
		infoLogger.Print(s)
		infoLogger.Print(runtime.FuncForPC(pc).Name())
		infoLogger.Print(path.Base(file))
		infoLogger.Print(line)
	}
}

func Error(s interface{}) {
	errorLogger.Print(s)
}
