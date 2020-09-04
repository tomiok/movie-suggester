package logs

import (
	"log"
	"os"
)

var (
	infoLogger  = log.New(os.Stdout, "INFO:", log.Llongfile|log.LstdFlags)
	errorLogger = log.New(os.Stdout, "ERROR:", log.Llongfile|log.LstdFlags)
)

func Info(s interface{}) {
	infoLogger.Print(s)
}

func Error(s interface{}) {
	errorLogger.Print(s)
}
