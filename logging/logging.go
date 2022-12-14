package logging

import (
	"sync"
)

// Level represents the logging level.
type Level uint8

const (
	LevelTrace Level = iota
	LevelDebug
	LevelInfo
	LevelWarn
	LevelError
	LevelOff
)

// Logger is the common interface to all loggers.
type Logger interface {
	// SetLevel sets the logging level for this specific Logger; if present, it overrides the global logging level.
	SetLevel(level Level)
	// GetLevel returns the current logging level for this specific Logger, or nil if no level is set.
	GetLevel() *Level
	// ResetLevel removes the per-logger level from this specific Logger, if present.
	ResetLevel()
	// Trace sends out a debug message with the given arguments to the logger.
	Trace(args ...interface{})
	// Tracef formats a debug message using the given arguments and sends it to the logger.
	Tracef(format string, args ...interface{})
	// Debug sends out a debug message with the given arguments to the logger.
	Debug(args ...interface{})
	// Debugf formats a debug message using the given arguments and sends it to the logger.
	Debugf(format string, args ...interface{})
	// Info sends out an informational message with the given arguments to the logger.
	Info(args ...interface{})
	// Infof formats an informational message using the given arguments and sends it to the logger.
	Infof(format string, args ...interface{})
	// Warn sends out a warning message with the given arguments to the logger.
	Warn(args ...interface{})
	// Warnf formats a warning message using the given arguments and sends it to the logger.
	Warnf(format string, args ...interface{})
	// Error sends out an error message with the given arguments to the logger.
	Error(args ...interface{})
	// Errorf formats an error message using the given arguments and sends it to the logger.
	Errorf(format string, args ...interface{})
}

var (
	lock1 sync.RWMutex
	level Level = LevelDebug
)

// SetLGlobalevel sets the logging level globally.
func SetGlobalLevel(l Level) {
	lock1.Lock()
	defer lock1.Unlock()
	level = l
}

// GetGlobalLevel retrieves the current global logging level.
func GetGlobalLevel() Level {
	lock1.RLock()
	defer lock1.RUnlock()
	return level
}

var (
	lock2  sync.RWMutex
	logger Logger = &NoOpLogger{}
)

// SetLogger sets the logger globally.
func SetLogger(l Logger) Logger {
	lock2.Lock()
	defer lock2.Unlock()
	logger = l
	return l
}

// GetLogger retrieves the current global logger.
func GetLogger() Logger {
	lock2.RLock()
	defer lock2.RUnlock()
	return logger
}
