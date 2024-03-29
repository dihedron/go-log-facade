package test

import (
	"bytes"
	"fmt"
	"runtime"
	"strings"
	"testing"

	"github.com/dihedron/go-log-facade/logging"
)

// Logger wraps the Golang testing framework logger.
type Logger struct {
	t      *testing.T
	caller bool
	level  *logging.Level
}

// NewLogger returns a Logger wrapping a testing logger.
func NewLogger(t *testing.T) *Logger {
	return &Logger{
		t:      t,
		caller: false,
	}
}

// NewLoggerWithStack returns a Logger wrapping a testing logger
// and printing the curernt.
func NewLoggerWithCaller(t *testing.T) *Logger {
	return &Logger{
		t:      t,
		caller: true,
	}
}

func (l *Logger) SetLevel(level logging.Level) {
	l.level = &level
}

func (l *Logger) GetLevel() *logging.Level {
	if l.level != nil {
		// there's a specific logging level for this logger
		return l.level
	}
	// there is no per-instance logging level, return the global level
	level := logging.GetGlobalLevel()
	return &level
}

func (l *Logger) ResetLevel() {
	l.level = nil
}

// Trace logs a message at LevelTrace level.
func (l *Logger) Trace(args ...interface{}) {
	if *l.GetLevel() <= logging.LevelTrace {
		message := l.format("TRC", args...)
		l.t.Log(message)
	}
}

// Tracef logs a message at LevelTrace level.
func (l *Logger) Tracef(msg string, args ...interface{}) {
	if *l.GetLevel() <= logging.LevelTrace {
		message := l.formatf("TRC", msg, args...)
		l.t.Log(message)
	}
}

// Debug logs a message at LevelDebug level.
func (l *Logger) Debug(args ...interface{}) {
	if *l.GetLevel() <= logging.LevelDebug {
		message := l.format("DBG", args...)
		l.t.Log(message)
	}
}

// Debugf logs a message at LevelDebug level.
func (l *Logger) Debugf(msg string, args ...interface{}) {
	if *l.GetLevel() <= logging.LevelDebug {
		message := l.formatf("DBG", msg, args...)
		l.t.Log(message)
	}
}

// Info logs a message at LevelInfo level.
func (l *Logger) Info(args ...interface{}) {
	if *l.GetLevel() <= logging.LevelInfo {
		message := l.format("INF", args...)
		l.t.Log(message)
	}
}

// Infof logs a message at LevelInfo level.
func (l *Logger) Infof(msg string, args ...interface{}) {
	if *l.GetLevel() <= logging.LevelInfo {
		message := l.formatf("INF", msg, args...)
		l.t.Log(message)
	}
}

// Warn logs a message at LevelWarn level.
func (l *Logger) Warn(args ...interface{}) {
	if *l.GetLevel() <= logging.LevelWarn {
		message := l.format("WRN", args...)
		l.t.Log(message)
	}
}

// Warnf logs a message at LevelWarn level.
func (l *Logger) Warnf(msg string, args ...interface{}) {
	if *l.GetLevel() <= logging.LevelWarn {
		message := l.formatf("WRN", msg, args...)
		l.t.Log(message)
	}
}

// Error logs a message at LevelError level.
func (l *Logger) Error(args ...interface{}) {
	if *l.GetLevel() <= logging.LevelError {
		message := l.format("ERR", args...)
		l.t.Log(message)
	}
}

// Errorf logs a message at LevelError level.
func (l *Logger) Errorf(msg string, args ...interface{}) {
	if *l.GetLevel() <= logging.LevelError {
		message := l.formatf("ERR", msg, args...)
		l.t.Log(message)
	}
}

func (l *Logger) format(level string, args ...interface{}) string {
	var buffer bytes.Buffer
	for argNum, arg := range args {
		if argNum > 0 {
			buffer.WriteString(" ")
		}
		buffer.WriteString(fmt.Sprintf("%v", arg))
	}
	extra := ""
	if l.caller {
		pc, _, _, ok := runtime.Caller(2)
		details := runtime.FuncForPC(pc)

		if ok && details != nil {
			line, no := details.FileLine(pc)
			extra = fmt.Sprintf(" (%s:%d)", line, no)
		}
	}
	message := fmt.Sprintf("[%s] %s%s", level, buffer.String(), extra)
	return strings.TrimRight(message, "\n\r")
}

func (l *Logger) formatf(level string, msg string, args ...interface{}) string {
	message := strings.TrimRight(fmt.Sprintf("["+level+"] "+strings.TrimSpace(msg), args...), "\n\r")
	if l.caller {
		pc, _, _, ok := runtime.Caller(2)
		details := runtime.FuncForPC(pc)
		if ok && details != nil {
			line, no := details.FileLine(pc)
			message = fmt.Sprintf("%s (%s:%d)", message, line, no)
		}
	}
	return message
}
