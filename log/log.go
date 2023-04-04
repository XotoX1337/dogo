package log

import (
	"fmt"
	"os"

	"github.com/XotoX1337/dogo/constants"
)

var prefixMap = map[int]string{
	constants.LOG_LEVEL_DEBUG: "DEBUG: ",
	constants.LOG_LEVEL_INFO:  "INFO: ",
	constants.LOG_LEVEL_WARN:  "WARN: ",
	constants.LOG_LEVEL_FATAL: "FATAL: ",
}

func Log(level int, message string) {
	fmt.Printf(
		"%s %s\n",
		prefixMap[level],
		message,
	)

	if level == constants.LOG_LEVEL_FATAL {
		os.Exit(1)
	}
}

func Debug(message string) {
	Log(constants.LOG_LEVEL_DEBUG, message)
}

func Info(message string) {
	Log(constants.LOG_LEVEL_INFO, message)
}

func Warn(message string) {
	Log(constants.LOG_LEVEL_WARN, message)
}

func Fatal(message string) {
	Log(constants.LOG_LEVEL_FATAL, message)
}
