package log

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/fatih/color"
)

// Logger - wraps standard log.Logger and extends it with logging level capabilities
type Logger struct {
	Level Level
	opt   options
}

type options struct {
	level Level
	color color.Attribute
	pfx   string
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

var logger = newLogger(os.Stdout, debug)

func newOptions(lvl Level) options {
	switch lvl {
	case debug:
		return options{color: color.FgHiMagenta, level: lvl, pfx: "DEBUG:"}
	case info:
		return options{color: color.FgHiBlue, level: lvl, pfx: "INFO:"}
	case warning:
		return options{color: color.FgHiYellow, level: lvl, pfx: "WARN:"}
	case err:
		return options{color: color.FgHiRed, level: lvl, pfx: "ERROR:"}
	default:
		return options{color: color.FgHiMagenta, level: lvl, pfx: "DEBUG:"}
	}
}

func newLogger(dest io.Writer, lvl Level) *Logger {
	logger := &Logger{Level: lvl, opt: newOptions(lvl)}

	return logger
}

func pfxStr(lvl Level) string {
	switch lvl {
	case debug:
		return "DEBUG:"
	case info:
		return "INFO:"
	case warning:
		return "WARN:"
	case err:
		return "ERROR:"
	default:
		return ""
	}
}

func pfxColor(lvl Level) func(a ...interface{}) string {
	switch lvl {
	case debug:
		return color.New(color.FgHiCyan).SprintFunc()
	case info:
		return color.New(color.FgHiBlue).SprintFunc()
	case warning:
		return color.New(color.FgHiYellow).SprintFunc()
	case err:
		return color.New(color.FgHiRed).SprintFunc()
	default:
		return color.New(color.FgHiCyan).SprintFunc()
	}
}

func date(format string) string {
	return time.Now().Format(format)
}

// SetLogLevel - sets log level as an environment variable, returns an error if any occurred
func SetLogLevel(lvl Level) {
	logger.Level = lvl
	logger.opt = newOptions(lvl)
}

func strFromArgs(args []interface{}) string {
	var str string
	for _, arg := range args {
		str += fmt.Sprintf(" %s", arg)
	}
	return str
}

// Debug - most verbose logging level
func Debug(args ...interface{}) {
	if logger.Level > debug {
		return
	}
	now := date(time.RFC822)
	pfx := pfxStr(debug)
	colored := pfxColor(debug)
	fmt.Printf("%s %s%s \n", colored(now), colored(pfx), strFromArgs(args))
}

// Info - slightly less verbose logging level
func Info(args ...interface{}) {
	if logger.Level > info {
		return
	}
	now := date(time.RFC822)
	pfx := pfxStr(info)
	colored := pfxColor(info)
	fmt.Printf("%s %s%s \n", colored(now), colored(pfx), strFromArgs(args))
}

// Warning - logging verbosity level displaying only crucial warnings and errors
func Warning(args ...interface{}) {
	if logger.Level > warning {
		return
	}
	now := date(time.RFC822)
	pfx := pfxStr(warning)
	colored := pfxColor(warning)
	fmt.Printf("%s %s%s \n", colored(now), colored(pfx), strFromArgs(args))
}

// Error - logging verbosity level displaying only errors
func Error(args ...interface{}) {
	if logger.Level > err {
		return
	}
	now := date(time.RFC822)
	pfx := pfxStr(err)
	colored := pfxColor(err)
	fmt.Printf("%s %s%s \n", colored(now), colored(pfx), strFromArgs(args))
}
