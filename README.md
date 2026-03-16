# Go Logging Linter

Статический анализатор для Go, проверяющий корректность использования логирования в `slog` и `zap`.


Поддерживаемые логгеры:
- `log/slog`
- `go.uber.org/zap`

---

## Запуск

Установить все нужные зависимости в директории linter:

```bash
go mod tidy
```

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
        test/checker.go:8:20: logrecords: log message should contain only English letters (mylinter)
                slog.Error("ошибка подключения") 
                        ^
        test/checker.go:9:19: logrecords: log message contains special symbols or emoji (mylinter)
                slog.Info("server started 🚀") 
                        ^
        test/checker.go:10:19: logrecords: log message contains special symbols or emoji (mylinter)
                slog.Warn("warning: something went wrong")
                        ^
        test/checker.go:11:20: logrecords: log message contains special symbols or emoji (mylinter)
                slog.Error("connection failed...") 
                        ^
4 issues:
* mylinter: 4
```

Линтер поддерживает автоматическое исправление некоторых ошибок:

```bash
./custom-golangci-lint run --fix ./...
```

При запуске в той же директории:
```bash
        test/checker.go:8:20: logrecords: log message should contain only English letters (mylinter)
                slog.Error("ошибка подключения") 
                           ^
1 issues:
* mylinter: 1
```

Тесты
```bash

go test ./...
```

Чтобы проверить линтер на собственном проекте, нужно перейти в директорию вашего проект и запустить линтер через глобальный путь к нему 
```bash
/path/to/custom-golangci-lint run ./...
```