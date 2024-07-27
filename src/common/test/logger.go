package common_test

import (
	"io"
	"log/slog"
	"testing"
)

func NewTestLogger(t *testing.T) *slog.Logger {
	t.Helper()
	return slog.New(
		slog.NewJSONHandler(io.Discard, &slog.HandlerOptions{}),
	)
}
