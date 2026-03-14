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

Запустить линтер:

./custom-golangci-lint run ./...

Автоматическое исправление

Линтер поддерживает автоматическое исправление некоторых ошибок:

./custom-golangci-lint run --fix ./...

Тесты

go test ./...