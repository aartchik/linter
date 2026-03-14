package test

import (
	"go.uber.org/zap"
	"log/slog"
	"os"
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	slog.Info("message")                     // want "log message should start with lowercase"
	logger.Info("message", "field", "value") // want "log message should start with lowercase" "log field key should start with lowercase"

	slog.Info("message", "field", "value")   // want "log field key should start with lowercase"
	logger.Info("message", "field", "value") // want "log field key should start with lowercase"

	slog.Info("message", slog.String("field", "value"))   // want "log field key should start with lowercase"
	logger.Info("message", slog.String("field", "value")) // want "log field key should start with lowercase"

	loggernew, _ := zap.NewDevelopment()

	loggernew.Info("message")                               // want "log message should start with lowercase"
	loggernew.Info("message", zap.String("field", "value")) // want "log field key should start with lowercase"

	slog.Warn("warning")                      // want "log message should start with lowercase"
	slog.Error("error", "fieldName", "value") // want "log message should start with lowercase" "log field key should start with lowercase"

	logger.Warn("warning")                      // want "log message should start with lowercase"
	logger.Error("error", "fieldName", "value") // want "log message should start with lowercase" "log field key should start with lowercase"

	slog.Info("message", "userID", "42")      // want "log field key should start with lowercase"
	logger.Info("message", "requestID", "42") // want "log field key should start with lowercase"

	slog.Info("message", slog.String("requestID", "42")) // want "log field key should start with lowercase"
	logger.Info("message", slog.String("userID", "42"))  // want "log field key should start with lowercase"

	loggernew.Warn("warning")                               // want "log message should start with lowercase"
	loggernew.Error("error", zap.String("requestID", "42")) // want "log message should start with lowercase" "log field key should start with lowercase"

	slog.Info("upper first letter", "upperKey", "value")                // want "log message should start with lowercase" "log field key should start with lowercase"
	logger.Info("upper first letter", slog.String("upperKey", "value")) // want "log message should start with lowercase" "log field key should start with lowercase"

	loggernew.Info("upper first letter", zap.String("upperKey", "value")) // want "log message should start with lowercase" "log field key should start with lowercase"

	slog.Info("message", "hTTPStatus", "200")                  // want "log field key should start with lowercase"
	logger.Info("message", slog.String("hTTPStatus", "200"))   // want "log field key should start with lowercase"
	loggernew.Info("message", zap.String("hTTPStatus", "200")) // want "log field key should start with lowercase"
}
