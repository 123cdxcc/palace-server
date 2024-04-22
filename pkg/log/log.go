package log

import (
	"fmt"
	"log/slog"
)

func Infof(format string, a ...any) {
	slog.Info(fmt.Sprintf(format, a...))
}

func Errorf(format string, a ...any) {
	slog.Error(fmt.Sprintf(format, a...))
}
