# Функциональные типы и заготовка валидаторов

Действующий main демонстрирует тип Add, метод Summ и slices.Delete. Реализация валидаторов паролей закомментирована, а тесты по-прежнему ожидают validator, digits, letters, minlen, and и or.

Учебная папка `basics/validationPass`. Код сохранён в исходном состоянии.

## Структура

- `.gitignore`
- `go.mod`
- `main.go`
- `main_test.go`

## Требования и команды

Проверено на Windows с Go 1.27.1 (windows/amd64); `go.mod` требует Go 1.27.1. Команды выполняются из корня репозитория.

Точка входа: `main.go`.

```sh
go run .
```

```sh
go test ./...
go vet ./...
```

## Результаты проверки

- `go test ./...`: ошибка (код 1).
- `go vet ./...`: ошибка (код 1).
- `go run .`: успешно (код 0).

Диагностика `go test ./...` (фактический результат):

```text
FAIL	validationPass [build failed]
FAIL
# validationPass [validationPass.test]
.\main_test.go:13:19: undefined: validator
.\main_test.go:19:24: undefined: digits
.\main_test.go:25:24: undefined: digits
.\main_test.go:31:24: undefined: digits
.\main_test.go:37:24: undefined: letters
.\main_test.go:43:24: undefined: letters
.\main_test.go:49:24: undefined: letters
.\main_test.go:55:24: undefined: minlen
.\main_test.go:61:24: undefined: and
.\main_test.go:61:28: undefined: digits
.\main_test.go:61:28: too many errors
```

Проверки выполнены 22 ноября 2026 года. Сборка и тесты не исправляют существующие учебные ошибки.
