package golang

import (
	"bytes"
	"fmt"
	golang "log"
	"os"
	"strings"

	"github.com/dihedron/go-log-facade/logging"
)

// Logger is te type wrapping the default Golang logger.
type Logger struct {
	logger *golang.Logger
	level  *logging.Level
}

// NewLogger returns a new Golang Logger.
func NewLogger(prefix string) *Logger {
	return &Logger{
		logger: golang.New(os.Stderr, prefix, golang.Ltime|golang.Ldate|golang.Lmicroseconds),
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

func (l *Logger) Trace(args ...interface{}) {
	if *l.GetLevel() <= logging.LevelTrace {
		var buffer bytes.Buffer
		for argNum, arg := range args {
			if argNum > 0 {
				buffer.WriteString(" ")
			}
			buffer.WriteString(fmt.Sprintf("%v", arg))
		}
		message := fmt.Sprintf("[TRC] %s", buffer.String())
		message = strings.TrimRight(message, "\n\r")
		golang.Print(message)
	}
}

func (l *Logger) Tracef(msg string, args ...interface{}) {
	if *l.GetLevel() <= logging.LevelTrace {
		message := fmt.Sprintf("[TRC] "+msg, args...)
		message = strings.TrimRight(message, "\n\r")
		golang.Print(message)
	}
}

func (l *Logger) Debug(args ...interface{}) {
	if *l.GetLevel() <= logging.LevelDebug {
		var buffer bytes.Buffer
		for argNum, arg := range args {
			if argNum > 0 {
				buffer.WriteString(" ")
			}
			buffer.WriteString(fmt.Sprintf("%v", arg))
		}
		message := fmt.Sprintf("[DBG] %s", buffer.String())
		message = strings.TrimRight(message, "\n\r")
		golang.Print(message)
	}
}

func (l *Logger) Debugf(msg string, args ...interface{}) {
	if *l.GetLevel() <= logging.LevelDebug {
		message := fmt.Sprintf("[DBG] "+msg, args...)
		message = strings.TrimRight(message, "\n\r")
		golang.Print(message)
	}
}

func (l *Logger) Info(args ...interface{}) {
	if *l.GetLevel() <= logging.LevelInfo {
		var buffer bytes.Buffer
		for argNum, arg := range args {
			if argNum > 0 {
				buffer.WriteString(" ")
			}
			buffer.WriteString(fmt.Sprintf("%v", arg))
		}
		message := fmt.Sprintf("[INF] %s", buffer.String())
		message = strings.TrimRight(message, "\n\r")
		golang.Print(message)
	}
}

func (l *Logger) Infof(msg string, args ...interface{}) {
	if *l.GetLevel() <= logging.LevelInfo {
		message := fmt.Sprintf("[INF] "+msg, args...)
		message = strings.TrimRight(message, "\n\r")
		golang.Print(message)
	}
}

func (l *Logger) Warn(args ...interface{}) {
	if *l.GetLevel() <= logging.LevelWarn {
		var buffer bytes.Buffer
		for argNum, arg := range args {
			if argNum > 0 {
				buffer.WriteString(" ")
			}
			buffer.WriteString(fmt.Sprintf("%v", arg))
		}
		message := fmt.Sprintf("[WRN] %s", buffer.String())
		message = strings.TrimRight(message, "\n\r")
		golang.Print(message)
	}
}

func (l *Logger) Warnf(msg string, args ...interface{}) {
	if *l.GetLevel() <= logging.LevelWarn {
		message := fmt.Sprintf("[WRN] "+msg, args...)
		message = strings.TrimRight(message, "\n\r")
		golang.Print(message)
	}
}

func (l *Logger) Error(args ...interface{}) {
	if *l.GetLevel() <= logging.LevelError {
		var buffer bytes.Buffer
		for argNum, arg := range args {
			if argNum > 0 {
				buffer.WriteString(" ")
			}
			buffer.WriteString(fmt.Sprintf("%v", arg))
		}
		message := fmt.Sprintf("[ERR] %s", buffer.String())
		message = strings.TrimRight(message, "\n\r")
		golang.Print(message)
	}
}

func (l *Logger) Errorf(msg string, args ...interface{}) {
	if *l.GetLevel() <= logging.LevelError {
		message := fmt.Sprintf("[ERR] "+msg, args...)
		message = strings.TrimRight(message, "\n\r")
		golang.Print(message)
	}
}
