package logs

import (
	"log"
	"os"
)

var (
	infoLogger  = log.New(os.Stdout, "", 0)
	errorLogger = log.New(os.Stdout, "", log.Lshortfile|log.Ltime)
)

func Info(s interface{}) {
	infoLogger.Print(s)
}

func Error(s interface{}) {
	errorLogger.Print(s)
}
