package logger

import (
	"log"
	"os"
)

var logger *log.Logger

func NewLogger() *log.Logger {
	if logger == nil {
		logger = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	}
	return logger
}
