package log

import (
	"context"
	"fmt"
)

var ctx = context.Background()

func Debug(format string, args ...any) {
	logger.Debug(fmt.Sprintf(format, args...))
}

func Info(format string, args ...any) {
	logger.Info(fmt.Sprintf(format, args...))
}

func Warn(format string, args ...any) {
	logger.Warn(fmt.Sprintf(format, args...))
}

func Fixme(format string, args ...any) {
	logger.Log(ctx, LevelFixMe, fmt.Sprintf(format, args...))
}

func Error(format string, args ...any) {
	logger.Error(fmt.Sprintf(format, args...))
}

func Fatal(format string, args ...any) {
	logger.Log(ctx, LevelFatal, fmt.Sprintf(format, args...))
}
