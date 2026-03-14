package test

import (
	"log/slog"
	"context"
	"os"

)


func main() {

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	slog.Info("запуск сервера") // want "log message should contain only English letters"
	slog.Error("ошибка подключения") // want "log message should contain only English letters"

	logger.Info("запуск сервера") // want "log message should contain only English letters"
	logger.Error("ошибка подключения") // want "log message should contain only English letters"

	slog.Info("server started 🚀") // want "log message contains special symbols or emoji"
	slog.Warn("warning: something went wrong") // want "log message contains special symbols or emoji"
	slog.Error("connection failed...") // want "log message contains special symbols or emoji"

	logger.Info("server started 🚀") // want "log message contains special symbols or emoji"
	logger.Warn("warning: something went wrong") // want "log message contains special symbols or emoji"
	logger.Error("connection failed...") // want "log message contains special symbols or emoji"

	slog.WarnContext(context.Background(), "Warning context lower first letter") // want "log message should start with lowercase"
	logger.WarnContext(context.Background(), "Warning context lower first letter") // want "log message should start with lowercase"

	
	slog.Info("message", "Field1", "value1") // want "log field key should start with lowercase"
	slog.Info("message", "field!!!", "value1") // want "log field key contains special symbols or emoji"
	slog.Info("message", "ключ", "value1") // want "log field key should contain only English letters"

	logger.Info("message", "Field1", "value1") // want "log field key should start with lowercase"
	logger.Info("message", "field!!!", "value1") // want "log field key contains special symbols or emoji"
	logger.Info("message", "ключ", "value1") // want "log field key should contain only English letters"

	
	slog.Info("message", slog.String("Field1", "value1")) // want "log field key should start with lowercase"
	slog.Info("message", slog.String("field!!!", "value1")) // want "log field key contains special symbols or emoji"
	slog.Info("message", slog.String("ключ", "value1")) // want "log field key should contain only English letters"

	logger.Info("message", slog.String("Field1", "value1")) // want "log field key should start with lowercase"
	logger.Info("message", slog.String("field!!!", "value1")) // want "log field key contains special symbols or emoji"
	logger.Info("message", slog.String("ключ", "value1")) // want "log field key should contain only English letters"


	slog.Info("password", "token")
	password := ""
	api_key := ""
	slog.Info("message", "password", password) // want "log message contains potential sensitive word"
	slog.Info("message", "api_key", api_key) // want "log message contains potential sensitive word"

	logger.Info("message", slog.String("password", password)) // want "log message contains potential sensitive word"

	slog.Info("Upper first letter!!!") // want "log message should start with lowercase" "log message contains special symbols or emoji"
	logger.Info("Пароль!!!") // want "log message should start with lowercase" "log message should contain only English letters" "log message contains special symbols or emoji"

	slog.InfoContext(context.Background(), "message", "Field1", "value1") // want "log field key should start with lowercase"
	slog.InfoContext(context.Background(), "message", slog.String("password", password)) // want "log message contains potential sensitive word"

	logger.InfoContext(context.Background(), "message", "Field1", "value1") // want "log field key should start with lowercase"
	logger.InfoContext(context.Background(), "message", slog.String("password", password)) // want "log message contains potential sensitive word"
}