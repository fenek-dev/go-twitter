package common

import (
	"log/slog"
	"os"
)

func SetupLogger(env string) *slog.Logger {

	log := slog.New(
		slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
	)

	return log
}
