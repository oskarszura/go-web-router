package logger

import (
	"log"
	"os"
)

var (
	logger	*log.Logger
)

const (
	DEBUG	= "debug"
	INFO	= "info"
	ERROR	= "error"
)


func buildPrefix(prefix string) string {
	return prefix + ": "
}

func Init(prefix string) {
	logger = log.New(os.Stdout, buildPrefix(prefix), log.LstdFlags)
}

func Log(level string, message string, args ...interface{}) {
	var logs []interface{}
	logs = append(logs, message)

	if len(args) > 0 {
		logs = append(logs, args)
	}

	switch level {
	case DEBUG:
		logger.Println(logs...)

	case INFO:
		logger.Println(logs...)

	case ERROR:
		logger.Println(logs...)

	default:
		logger.Println("Not supported log level")
	}
}