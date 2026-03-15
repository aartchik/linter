# Go Logging Linter

Статический анализатор для Go, проверяющий корректность использования логирования в `slog` и `zap`.


Поддерживаемые логгеры:
- `log/slog`
- `go.uber.org/zap`

---

## Запуск

Собрать кастомный `golangci-lint` в директории linter:

```bash
golangci-lint custom
```

Запустить линтер:
```bash
./custom-golangci-lint run ./...
```

При запуске линтера внутри директории linter должно вывести:
```bash
        slog.Error("ошибка подключения") // want "log message should contain only English letters"
                   ^
test/checker.go:11:12: logrecords: log message contains special symbols or emoji (mylinter)
        slog.Info("server started 🚀") // want "log message contains special symbols or emoji"
                  ^
test/checker.go:12:12: logrecords: log message contains special symbols or emoji (mylinter)
        slog.Warn("warning: something went wrong") // want "log message contains special symbols or emoji"
                  ^
test/checker.go:13:13: logrecords: log message contains special symbols or emoji (mylinter)
        slog.Error("connection failed...") // want "log message contains special symbols or emoji"
                   ^
5 issues:
* mylinter: 5
```

Линтер поддерживает автоматическое исправление некоторых ошибок:

```bash
./custom-golangci-lint run --fix ./...
```

При запуске в той же директории:
```bash
test/checker.go:8:12: logrecords: log message should contain only English letters (mylinter)
        slog.Info("запуск сервера") // want "log message should contain only English letters"
                  ^
test/checker.go:9:13: logrecords: log message should contain only English letters (mylinter)
        slog.Error("ошибка подключения") // want "log message should contain only English letters"
                   ^
2 issues:
* mylinter: 2
```

Тесты
```bash

go test ./...
```

Чтобы проверить линтер на собственном проекте, нужно перейти в директорию вашего проект и запустить линтер через глобальный путь к нему 
```bash
/path/to/custom-golangci-lint run ./...
```