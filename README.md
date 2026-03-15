# Go Logging Linter

Статический анализатор для Go, проверяющий корректность использования логирования в `slog` и `zap`.


Поддерживаемые логгеры:
- `log/slog`
- `go.uber.org/zap`

---

## Запуск

Собрать кастомный `golangci-lint`:

```bash
golangci-lint custom
```

Запустить линтер:
```bash
./custom-golangci-lint run ./...
```

Линтер поддерживает автоматическое исправление некоторых ошибок:

```bash
./custom-golangci-lint run --fix ./...
```
Тесты
```bash

go test ./...
```