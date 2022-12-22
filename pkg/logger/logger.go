package logger

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"time"
)

const (
	DEBUG = "DEBUG"
	INFO  = "INFO"
	WARN  = "WARN"
	ERROR = "ERROR"
)

type Logger interface {
	Debugf(format string, v ...interface{})
	Infof(format string, v ...interface{})
	Warnf(format string, v ...interface{})
	Errorf(format string, v ...interface{})
}

type BasicLogger struct {
	Logger        *log.Logger
	isStackdriver bool
}

func NewLogger(w io.Writer, isStackdriver bool) Logger {
	return &BasicLogger{
		Logger:        log.New(w, "", log.LstdFlags),
		isStackdriver: isStackdriver,
	}
}

func stackdriverFmtJSON(level string, message string) string {
	entry := map[string]string{
		"severity": level,
		"message":  message,
		"time":     time.Now().Format(time.RFC3339Nano),
	}
	bytes, _ := json.Marshal(entry)
	return string(bytes)
}

func localFmt(level string, message string) string {
	return fmt.Sprintf("%s %s", level, message)
}

func (bl *BasicLogger) Printf(level string, format string, v ...interface{}) {
	m := fmt.Sprintf(format, v...)
	if bl.isStackdriver {
		bl.Logger.Println(stackdriverFmtJSON(level, m))
	} else {
		bl.Logger.Println(localFmt(level, m))
	}
}

func (bl *BasicLogger) Debugf(format string, v ...interface{}) {
	bl.Printf(DEBUG, format, v...)
}

func (bl *BasicLogger) Infof(format string, v ...interface{}) {
	bl.Printf(INFO, format, v...)
}

func (bl *BasicLogger) Warnf(format string, v ...interface{}) {
	bl.Printf(WARN, format, v...)
}

func (bl *BasicLogger) Errorf(format string, v ...interface{}) {
	bl.Printf(ERROR, format, v...)
}
