package ecs

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/dihedron/go-log-facade/logging"
	"go.elastic.co/ecszap"
	"go.uber.org/zap"
)

// Logger is an adapter that allows to log using Uber's Zap
// wherever a Logger interface is expected.
type Logger struct {
	logger           *zap.Logger
	level            *logging.Level
	path             string
	enableName       bool
	enableCaller     bool
	enableStackTrace bool
}

var (
	configuration zap.Config
	Restore       func()
)

// Option is the type for functional options.
type Option func(*Logger)

// WithPath allows to specify the path to the output file.
func WithPath(path string) Option {
	return func(l *Logger) {
		if path != "" {
			l.path = path
		}
	}
}

// WithNameEnabled allows to specify whether the logger name should be logged.
func WithNameEnabled(enabled bool) Option {
	return func(l *Logger) {
		l.enableName = enabled
	}
}

// WithCallerEnabled allows to specify whether the caller should be logged.
func WithCallerEnabled(enabled bool) Option {
	return func(l *Logger) {
		l.enableCaller = enabled
	}
}

// WithCallerEnabled allows to specify whether the stack trace should be logged.
func WithStackTraceEnabled(enabled bool) Option {
	return func(l *Logger) {
		l.enableStackTrace = enabled
	}
}

// WithLevel allows to specify the logger level.
func WithLevel(level string) Option {
	return func(l *Logger) {
		switch strings.ToLower(level) {
		case "trace", "trc", "t":
			tmp := logging.LevelTrace
			l.level = &tmp
		case "debug", "dbg", "d":
			tmp := logging.LevelDebug
			l.level = &tmp
		case "info", "inf", "i":
			tmp := logging.LevelInfo
			l.level = &tmp
		case "warning", "warn", "wrn", "w":
			tmp := logging.LevelWarn
			l.level = &tmp
		case "error", "err", "e", "fatal", "ftl", "f":
			tmp := logging.LevelError
			l.level = &tmp
		case "off":
			tmp := logging.LevelOff
			l.level = &tmp
		}
	}
}

// TODO: add more options here....

// NewLogger initialises an ECS compatible Zap logger for sending events to Elastic;
// use the functional options to fine tune how the logger should behave.
func NewLogger(options ...Option) (*Logger, error) {

	result := &Logger{
		enableName:       true,
		enableCaller:     true,
		enableStackTrace: true,
		path:             strings.Replace(filepath.Base(os.Args[0]), ".exe", fmt.Sprintf("-%d.log", os.Getpid()), 1), // FIXME: fix on linux where there is no ".exe"
	}

	for _, option := range options {
		option(result)
	}

	configuration := ecszap.NewDefaultEncoderConfig()
	configuration.EnableCaller = result.enableCaller
	configuration.EnableStackTrace = result.enableStackTrace
	configuration.EnableName = result.enableName
	core := ecszap.NewCore(configuration, os.Stdout, zap.DebugLevel)
	logger := zap.New(core, zap.AddCaller())

	// encoderConfig := ecszap.EncoderConfig{
	// 	EncodeName:     customNameEncoder,
	// 	EncodeLevel:    zapcore.CapitalLevelEncoder,
	// 	EncodeDuration: zapcore.MillisDurationEncoder,
	// 	EncodeCaller:   ecszap.FullCallerEncoder,
	// }
	// core := ecszap.NewCore(encoderConfig, os.Stdout, zap.DebugLevel)
	// logger := zap.New(core, zap.AddCaller())

	// // check if there's a file called brokerd-log.json aside the
	// // application excutable; if so, load it as it contains the
	// // logger configuration; if not, assume default for production
	// app := strings.Replace(filepath.Base(os.Args[0]), ".exe", "", 1)
	// content, err := ioutil.ReadFile(app + "-log.json")
	// if err == nil { // the file exists
	// 	if err := json.Unmarshal(content, &configuration); err != nil {
	// 		return nil, fmt.Errorf("error unmarshalling log configuration from '%s': %w", app+"-log.json", err)
	// 	}
	// 	// update the field tags to make Elastic happy
	// 	fillForElastic(&configuration)
	// 	logger, err := configuration.Build()
	// 	if err != nil {
	// 		return nil, fmt.Errorf("error bulding logging configuration: %w", err)
	// 	}
	// 	Restore = zap.ReplaceGlobals(logger)
	// 	logger.Info("application starting with custom log configuration")
	// 	return &Logger{
	// 		logger: logger.WithOptions(zap.AddCallerSkip(1)),
	// 		// logger: logger,
	// 	}, nil
	// }
	// // configuration does not exist, use default
	// configuration = zap.NewProductionConfig()
	// configuration.Encoding = "json" // or "console"
	// // update the field tags to make elastic happy
	// fillForElastic(&configuration)
	// configuration.OutputPaths = []string{fmt.Sprintf("%s-%d.log", app, os.Getpid())}
	// configuration.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	// logger, err := configuration.Build()
	// if err != nil {
	// 	return nil, fmt.Errorf("error initialising logger: %w", err)
	// }
	// Restore = zap.ReplaceGlobals(logger)
	// logger.Info("application starting with default log configuration")

	return &Logger{
		logger: logger.WithOptions(zap.AddCallerSkip(1)),
		//logger: logger,
	}, nil
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
		l.logger.Sugar().Debug(args...)
	}
}

// Tracef logs a message at LevelTrace level.
func (l *Logger) Tracef(format string, args ...interface{}) {
	if *l.GetLevel() <= logging.LevelTrace {
		l.logger.Sugar().Debugf(format, args...)
	}
}

// Debug logs a message at LevelDebug level.
func (l *Logger) Debug(args ...interface{}) {
	if *l.GetLevel() <= logging.LevelDebug {
		l.logger.Sugar().Debug(args...)
	}
}

// Debugf logs a message at LevelDebug level.
func (l *Logger) Debugf(format string, args ...interface{}) {
	if *l.GetLevel() <= logging.LevelDebug {
		l.logger.Sugar().Debugf(format, args...)
	}
}

// Info logs a message at LevelInfo level.
func (l *Logger) Info(args ...interface{}) {
	if *l.GetLevel() <= logging.LevelInfo {
		l.logger.Sugar().Info(args...)
	}
}

// Infof logs a message at LevelInfo level.
func (l *Logger) Infof(format string, args ...interface{}) {
	if *l.GetLevel() <= logging.LevelInfo {
		l.logger.Sugar().Infof(format, args...)
	}
}

// Warn logs a message at LevelWarn level.
func (l *Logger) Warn(args ...interface{}) {
	if *l.GetLevel() <= logging.LevelWarn {
		l.logger.Sugar().Warn(args...)
	}
}

// Warnf logs a message at LevelWarn level.
func (l *Logger) Warnf(format string, args ...interface{}) {
	if *l.GetLevel() <= logging.LevelWarn {
		l.logger.Sugar().Warnf(format, args...)
	}
}

// Error logs a message at LevelError level.
func (l *Logger) Error(args ...interface{}) {
	if *l.GetLevel() <= logging.LevelError {
		l.logger.Sugar().Error(args...)
	}
}

// Errorf logs a message at LevelError level.
func (l *Logger) Errorf(format string, args ...interface{}) {
	if *l.GetLevel() <= logging.LevelError {
		l.logger.Sugar().Errorf(format, args...)
	}
}

func fillForElastic(configuration *zap.Config) {
	// configuration.EncoderConfig.MessageKey = "message"
	// configuration.EncoderConfig.LevelKey = "log.level"
	// configuration.EncoderConfig.TimeKey = "@timestamp"
	// configuration.EncoderConfig.NameKey = "log.logger"
	// configuration.EncoderConfig.CallerKey = "log.origin.file.name"
	// configuration.EncoderConfig.StacktraceKey = "error.stack_trace"
	// configuration.EncoderConfig.EncodeTime = zapcore.RFC3339NanoTimeEncoder
	// configuration.InitialFields = map[string]interface{}{
	// 	"service.name":        appinfo.ServiceName,
	// 	"service.version":     fmt.Sprintf("v%s@%s", appinfo.GitTag, appinfo.GitCommit),
	// 	"service.environment": os.Getenv("BROKERD_STAGE"),
	// }
	// if configuration.InitialFields["service.environment"] == "" {
	// 	configuration.InitialFields["service.environment"] = "development"
	// }
}
