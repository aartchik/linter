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




	logger.Info("lower first letter", "test", "test", "test", "test", "test", "test")
	logger.InfoContext(context.WithValue(context.Background(), "UPPER", "42"), "warn context lower first letter")
	logger.ErrorContext(context.WithValue(context.Background(), "UPPER", "42"), "warn context lower first letter")
	logger.WarnContext(context.WithValue(context.Background(), "UPPER", "42"), "warn context lower first letter")
	logger.InfoContext(context.WithValue(context.Background(), "UPPER", "42"), "message", "field1", "value1", slog.String("field2", "value2"))
	logger.InfoContext(context.WithValue(context.Background(), "UPPER", "TEST"), "info context lower first letter")

	logger.Info("server " + "started")
	logger.InfoContext(context.Background(), "request " + "started")

	logger.Info("request completed", "request_id", "req-1")
	logger.Info("user authenticated", "user_id", "42")
	logger.Info("worker started", "worker-id", "w-1")
	logger.Info("cache updated", "cache_key", "user-42")

	logger.Info("request completed", slog.String("request_id", "req-1"))
	logger.Info("user authenticated", slog.String("user_id", "42"))
	logger.Info("worker started", slog.String("worker-id", "w-1"))
	logger.Info("cache updated", slog.String("cache_key", "user-42"))

	logger.Info("job finished",
		slog.String("request_id", "req-1"),
		slog.Int("status_code", 200),
		slog.Bool("cached", true),
	)

	logger.InfoContext(context.Background(),
		"request completed",
		slog.String("trace_id", "abc-123"),
		slog.Int("duration_ms", 15),
	)

	logger.Info("user authenticated successfully")
	logger.Debug("api request completed")
	logger.Info("token validated")
	logger.Info("password updated successfully")
	logger.Info("secret rotation completed")

	slog.Info("request completed", "request_id", "req-2")
	slog.Info("worker started", slog.String("worker_id", "worker-1"))
	slog.InfoContext(context.Background(), "request completed", "status_code", 200)

	logger.InfoContext(context.WithValue(context.Background(), "UPPER", "42"), "request completed", "status_code", 200)
	logger.InfoContext(context.WithValue(context.Background(), "UPPER", "42"), "job finished", slog.String("worker_id", "worker-2"))

	logger.Info("request completed",
		"user_id", "42",
		"request_id", "req-3",
		slog.String("trace_id", "trace-1"),
	)

	logger.Info("request 200 completed")
	logger.InfoContext(context.Background(), "worker 2 started")
}