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

func log(level int, message string, format ...any) {
	fmt.Printf(
		"%s %s\n",
		prefixMap[level],
		fmt.Sprintf(message, format...),
	)
	if level == constants.LOG_LEVEL_FATAL {
		os.Exit(1)
	}
}

func Debug(message string, format ...any) {
	log(constants.LOG_LEVEL_DEBUG, message, format...)
}

func Info(message string, format ...any) {
	log(constants.LOG_LEVEL_INFO, message, format...)
}

func Warn(message string, format ...any) {
	log(constants.LOG_LEVEL_WARN, message, format...)
}

func Fatal(message string, format ...any) {
	log(constants.LOG_LEVEL_FATAL, message, format...)
}
