package test

import (
	"log/slog"
	"os"
	"go.uber.org/zap"
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	slog.Info("Message") // want "log message should start with lowercase"
	logger.Info("Message", "Field", "value") // want "log message should start with lowercase" "log field key should start with lowercase"

	slog.Info("message", "Field", "value") // want "log field key should start with lowercase"
	logger.Info("message", "Field", "value") // want "log field key should start with lowercase"

	slog.Info("message", slog.String("Field", "value")) // want "log field key should start with lowercase"
	logger.Info("message", slog.String("Field", "value")) // want "log field key should start with lowercase"

	loggernew, _ := zap.NewDevelopment()

	loggernew.Info("Message") // want "log message should start with lowercase"
	loggernew.Info("message", zap.String("Field", "value")) // want "log field key should start with lowercase"

	slog.Warn("Warning") // want "log message should start with lowercase"
	slog.Error("Error", "FieldName", "value") // want "log message should start with lowercase" "log field key should start with lowercase"

	logger.Warn("Warning") // want "log message should start with lowercase"
	logger.Error("Error", "FieldName", "value") // want "log message should start with lowercase" "log field key should start with lowercase"

	slog.Info("message", "UserID", "42") // want "log field key should start with lowercase"
	logger.Info("message", "RequestID", "42") // want "log field key should start with lowercase"

	slog.Info("message", slog.String("RequestID", "42")) // want "log field key should start with lowercase"
	logger.Info("message", slog.String("UserID", "42")) // want "log field key should start with lowercase"

	loggernew.Warn("Warning") // want "log message should start with lowercase"
	loggernew.Error("Error", zap.String("RequestID", "42")) // want "log message should start with lowercase" "log field key should start with lowercase"

	slog.Info("Upper first letter", "UpperKey", "value") // want "log message should start with lowercase" "log field key should start with lowercase"
	logger.Info("Upper first letter", slog.String("UpperKey", "value")) // want "log message should start with lowercase" "log field key should start with lowercase"

	loggernew.Info("Upper first letter", zap.String("UpperKey", "value")) // want "log message should start with lowercase" "log field key should start with lowercase"

}