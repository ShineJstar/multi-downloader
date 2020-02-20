package app

import (
	"io"
	"log"
	"os"
)

// LogLevel set
const (
	NOTSET  LogLevel = 0
	DEBUG   LogLevel = 10
	INFO    LogLevel = 20
	WARNING LogLevel = 30
	ERROR   LogLevel = 40
)

type (
	// Logger extend logging type
	Logger struct {
		level LogLevel

		*log.Logger
	}

	// LogLevel .
	LogLevel int
)

// NewLogger 返回 Logger 指针对象
func NewLogger(level, path string) (lg *Logger, err error) {
	var (
		logLevel    = NOTSET
		logLevelMap = map[string]LogLevel{"DEBUG": DEBUG, "INFO": INFO, "WARN": WARNING, "WARNING": WARNING, "ERROR": ERROR}
	)

	if v, ok := logLevelMap[level]; ok {
		logLevel = v
	}

	logFile, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0664)
	if err != nil {
		return
	}

	mw := io.MultiWriter(logFile, os.Stdout) // 日志同时写入文件和终端
	logger := log.New(mw, "", log.LstdFlags|log.Lshortfile)
	lg = &Logger{logLevel, logger}

	return
}

// Debug log a message with severity "DEBUG"
func (lg *Logger) Debug(format string, v ...interface{}) {
	if lg.level <= DEBUG {
		lg.Printf("DEBUG: "+format, v...)
	}
}

// Info log a message with severity "INFO"
func (lg *Logger) Info(format string, v ...interface{}) {
	if lg.level <= INFO {
		lg.Printf("INFO: "+format, v...)
	}
}

// Warning log a message with severity "WARNING"
func (lg *Logger) Warning(format string, v ...interface{}) {
	if lg.level <= WARNING {
		lg.Printf("WARNING: "+format, v...)
	}
}

// Error log a message with severity "ERROR"
func (lg *Logger) Error(format string, v ...interface{}) {
	if lg.level <= ERROR {
		lg.Printf("ERROR: "+format, v...)
	}
}
