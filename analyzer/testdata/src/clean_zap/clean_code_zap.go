package test

import "go.uber.org/zap"

func main() {
	logger, _ := zap.NewDevelopment()
	defer logger.Sync()

	logger.Info("server started")
	logger.Warn("connection timeout")
	logger.Error("request failed")
	logger.Debug("request completed")

	zap.ReplaceGlobals(logger)
	zap.L().Info("global logger started")
	zap.L().Warn("cache miss detected")
	zap.L().Error("request failed")
	zap.L().Debug("response encoded")

	logger.Info("request completed", zap.String("request_id", "req-1"))
	logger.Info("user authenticated", zap.String("user_id", "42"))
	logger.Info("job finished", zap.Int("status_code", 200))
	logger.Info("cache updated", zap.Bool("cached", true))
	logger.Info("worker started", zap.String("worker-id", "w-1"))
	logger.Info("request completed", zap.String("cache_key", "user-42"))

	logger.Info(
		"request completed",
		zap.String("request_id", "req-2"),
		zap.Int("status_code", 200),
		zap.Bool("cached", true),
	)

	logger.Info(
		"user authenticated",
		zap.String("user_id", "42"),
		zap.String("trace_id", "trace-1"),
	)

	logger.Debug(
		"job finished",
		zap.Int("duration_ms", 15),
		zap.String("worker_id", "worker-2"),
	)

	logger.Info("user authenticated successfully")
	logger.Debug("api request completed")
	logger.Info("token validated")
	logger.Info("password updated successfully")
	logger.Info("secret rotation completed")

	zap.L().Info("request completed", zap.String("request_id", "req-3"))
	zap.L().Info("worker started", zap.String("worker_id", "worker-3"))
	zap.L().Info("cache updated", zap.Bool("cached", true))

	logger.Info("request completed", zap.String("request_id", "req-4"))
	logger.Info("worker started", zap.String("worker-id", "worker-4"))

	logger.Info("request 200 completed")
	logger.Debug("worker 2 started")
}