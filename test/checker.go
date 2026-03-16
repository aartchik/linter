 package test

import (
  "log/slog"
)

func main() {
	slog.Error("ошибка подключения") 
	slog.Info("server started 🚀") 
	slog.Warn("warning: something went wrong")
	slog.Error("connection failed...") 
}
