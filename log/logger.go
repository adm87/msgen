package log

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"os"
)

var slogger = slog.New(&SimpleHandler{
	writer: os.Stdout,
	level:  slog.LevelInfo,
})

type SimpleHandler struct {
	writer io.Writer
	level  slog.Level
}

func (h *SimpleHandler) Enabled(_ context.Context, level slog.Level) bool {
	return level >= h.level
}

func (h *SimpleHandler) Handle(_ context.Context, rec slog.Record) error {
	fmt.Fprintf(h.writer, "%s %s", rec.Level.String(), rec.Message)

	rec.Attrs(func(attr slog.Attr) bool {
		fmt.Fprintf(h.writer, " %s=%v", attr.Key, attr.Value.Any())
		return true
	})

	fmt.Fprintln(h.writer)
	return nil
}

func (h *SimpleHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return h
}

func (h *SimpleHandler) WithGroup(name string) slog.Handler {
	return h
}

func Info(msg string, args ...any) {
	slogger.Info(msg, args...)
}

func Warn(msg string, args ...any) {
	slogger.Warn(msg, args...)
}

func Error(msg string, args ...any) {
	slogger.Error(msg, args...)
}
