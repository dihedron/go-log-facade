package hcl

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/dihedron/go-log-facade/logging"
	"github.com/hashicorp/go-hclog"
)

// Logger is the tpe warring an HCL logger.
type Logger struct {
	logger hclog.Logger
	level  *logging.Level
}

// NewLogger returns an instance of HCL logger wrapper
// that complies with the logging.Logger interface.
func NewLogger(logger hclog.Logger) *Logger {
	return &Logger{
		logger: logger,
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
		message := l.format(args...)
		l.logger.Trace(message)
	}
}

// Tracef logs a message at LevelTrace level.
func (l *Logger) Tracef(msg string, args ...interface{}) {
	if *l.GetLevel() <= logging.LevelTrace {
		message := l.formatf(msg, args...)
		l.logger.Trace(message)
	}
}

// Debug logs a message at LevelDebug level.
func (l *Logger) Debug(args ...interface{}) {
	if *l.GetLevel() <= logging.LevelDebug {
		message := l.format(args...)
		l.logger.Debug(message)
	}
}

// Debugf logs a message at LevelDebug level.
func (l *Logger) Debugf(msg string, args ...interface{}) {
	if *l.GetLevel() <= logging.LevelDebug {
		message := l.formatf(msg, args...)
		l.logger.Debug(message)
	}
}

// Info logs a message at LevelInfo level.
func (l *Logger) Info(args ...interface{}) {
	if *l.GetLevel() <= logging.LevelInfo {
		message := l.format(args...)
		l.logger.Info(message)
	}
}

// Infof logs a message at LevelInfo level.
func (l *Logger) Infof(msg string, args ...interface{}) {
	if *l.GetLevel() <= logging.LevelInfo {
		message := l.formatf(msg, args...)
		l.logger.Info(message)
	}
}

// Warn logs a message at LevelWarn level.
func (l *Logger) Warn(args ...interface{}) {
	if *l.GetLevel() <= logging.LevelWarn {
		message := l.format(args...)
		l.logger.Warn(message)
	}
}

// Warnf logs a message at LevelWarn level.
func (l *Logger) Warnf(msg string, args ...interface{}) {
	if *l.GetLevel() <= logging.LevelWarn {
		message := l.formatf(msg, args...)
		l.logger.Warn(message)
	}
}

// Error logs a message at LevelError level.
func (l *Logger) Error(args ...interface{}) {
	if *l.GetLevel() <= logging.LevelError {
		message := l.format(args...)
		l.logger.Error(message)
	}
}

// Errorf logs a message at LevelError level.
func (l *Logger) Errorf(msg string, args ...interface{}) {
	if *l.GetLevel() <= logging.LevelError {
		message := l.formatf(msg, args...)
		l.logger.Error(message)
	}
}

func (l *Logger) format(args ...interface{}) string {
	var buffer bytes.Buffer
	for argNum, arg := range args {
		if argNum > 0 {
			buffer.WriteString(" ")
		}
		buffer.WriteString(fmt.Sprintf("%v", arg))
	}
	return strings.TrimRight(buffer.String(), "\n\r")
}

func (l *Logger) formatf(msg string, args ...interface{}) string {
	message := fmt.Sprintf(msg, args...)
	return strings.TrimRight(message, "\n\r")
}
