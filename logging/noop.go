package logging

// NoOpLogger is a logger that writes nothing.
type NoOpLogger struct{}

// SetLevel does nothing.
func (l *NoOpLogger) SetLevel(_ Level) {}

// GetLevel does nothing.
func (l *NoOpLogger) GetLevel() *Level { return nil }

// ResetLevel does nothing.
func (l *NoOpLogger) ResetLevel() {}

// Trace logs a message at LevelTrace level.
func (*NoOpLogger) Trace(args ...interface{}) {}

// Tracef logs a message at LevelTrace level.
func (*NoOpLogger) Tracef(format string, args ...interface{}) {}

// Debug logs a message at LevelDebug level.
func (*NoOpLogger) Debug(args ...interface{}) {}

// Debugf logs a message at LevelDebug level.
func (*NoOpLogger) Debugf(format string, args ...interface{}) {}

// Info logs a message at LevelInfo level.
func (*NoOpLogger) Info(args ...interface{}) {}

// Infof logs a message at LevelInfo level.
func (*NoOpLogger) Infof(format string, args ...interface{}) {}

// Warn logs a message at LevelWarn level.
func (*NoOpLogger) Warn(args ...interface{}) {}

// Warnf logs a message at LevelWarn level.
func (*NoOpLogger) Warnf(format string, args ...interface{}) {}

// Error logs a message at LevelError level.
func (*NoOpLogger) Error(args ...interface{}) {}

// Errorf logs a message at LevelError level.
func (*NoOpLogger) Errorf(format string, args ...interface{}) {}
