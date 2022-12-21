package logging

import (
	"context"
)

type KeyType string

const key KeyType = "__logging_key__"

// Ctx returns a new context that holds a reference to the given
// logger.
func Ctx(ctx context.Context, logger Logger) context.Context {
	return context.WithValue(ctx, key, logger)
}

// FromContext returns the logger that was recorded in the context,
// if any.
func FromContext(ctx context.Context) Logger {
	if logger, ok := ctx.Value(key).(Logger); ok {
		return logger
	}
	return nil
}
