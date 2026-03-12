package test

import (
	"context"
	"log/slog"
	"os"
)


func main() {

	slog.Info("lower first letter")
	slog.Error("error lower first letter")
	slog.Warn("warn lower first letter")
	slog.Debug("debug lower first letter")
	slog.InfoContext(context.WithValue(context.Background(), "UPPER", "42"), "warn CONTEXT lower first letter")
	slog.ErrorContext(context.WithValue(context.Background(), "UPPER", "42"), "warn CONTEXT lower first letter")
	slog.WarnContext(context.WithValue(context.Background(), "UPPER", "42"), "warn CONTEXT lower first letter")
	slog.InfoContext(context.WithValue(context.Background(), "UPPER", "42"), "message", "test", "test", slog.String("field2", "value2"))
	slog.InfoContext(context.WithValue(context.Background(), "UPPER", "TEST"), "info CONTEXT lower first letter")


	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	logger.Info("lower first letter")
	logger.Error("error lower first letter")
	logger.Warn("warn lower first letter")
	logger.Debug("debug lower first letter")
	logger.Info("lower first letter", "test", "test", "test", "test", "test", "test")
	logger.InfoContext(context.WithValue(context.Background(), "UPPER", "42"), "warn CONTEXT lower first letter")
	logger.ErrorContext(context.WithValue(context.Background(), "UPPER", "42"), "warn CONTEXT lower first letter")
	logger.WarnContext(context.WithValue(context.Background(), "UPPER", "42"), "warn CONTEXT lower first letter")
	logger.InfoContext(context.WithValue(context.Background(), "UPPER", "42"), "message", "field1", "value1", slog.String("field2", "value2"))
	logger.InfoContext(context.WithValue(context.Background(), "UPPER", "TEST"), "info CONTEXT lower first letter")




	logger.Info("server " + "Started")
	logger.InfoContext(context.Background(), "request " + "started")

}