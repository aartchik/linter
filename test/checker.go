package test

import (
	"log/slog"
)

func main() {
	slog.Info("запуск сервера")      // want "log message should contain only English letters"
	slog.Error("ошибка подключения") // want "log message should contain only English letters"

	slog.Info("server started")               // want "log message contains special symbols or emoji"
	slog.Warn("warning something went wrong") // want "log message contains special symbols or emoji"
	slog.Error("connection failed")           // want "log message contains special symbols or emoji"
}
