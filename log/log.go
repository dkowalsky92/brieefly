package log

import (
	"io"
	"log"
	"os"
)

// Logger - wraps standard log.Logger and extends it with logging level capabilities
type Logger struct {
	stdlogger *log.Logger
	logLevel  Level
}

// Level - defines verbosity of logs
type Level int8

const (
	debug     Level = 0
	info      Level = 1
	warning   Level = 2
	err       Level = 3
	undefined Level = 4

	logLevelKey string = "logLevel"
)

var logger = newLogger(os.Stdout, debug, log.Ldate|log.Ltime)

func newLogger(dest io.Writer, lvl Level, flags int) *Logger {
	prefix := lvl.String()
	stdlogger := log.New(dest, prefix, flags)
	logger := &Logger{stdlogger: stdlogger, logLevel: lvl}

	return logger
}

func (l Level) String() string {
	var val string
	switch l {
	case debug:
		val = "DEBUG: "
	case info:
		val = "INFO: "
	case warning:
		val = "WARNING: "
	case err:
		val = "ERROR: "
	default:
		val = "UNDEFINED: "
	}
	return val
}

// SetLogLevel - sets log level as an environment variable, returns an error if any occurred
func SetLogLevel(lvl Level) {
	logger.stdlogger.SetPrefix(lvl.String())
	logger.logLevel = lvl
}

// Debug - most verbose logging level
func Debug(args ...interface{}) {
	if logger.logLevel > debug {
		return
	}

	logger.stdlogger.Println(args)
}

// Info - slightly less verbose logging level
func Info(args ...interface{}) {
	if logger.logLevel > info {
		return
	}

	logger.stdlogger.Println(args)
}

// Warning - logging verbosity level displaying only crucial warnings and errors
func Warning(args ...interface{}) {
	if logger.logLevel > warning {
		return
	}

	logger.stdlogger.Println(args)
}

// Error - logging verbosity level displaying only errors
func Error(args ...interface{}) {
	if logger.logLevel > err {
		return
	}

	logger.stdlogger.Println(args)
}
