## Лабораторная: Go бинарь + вызов из Python (subprocess)

### Структура
- `cmd/hello/main.go`: Go-программа (CLI)
- `bin/hello`: собранный бинарь (создаётся командой build или Python-скриптом)
- `run_hello.py`: Python-скрипт, который при необходимости собирает Go-бинарь и запускает его через `subprocess`

### Сборка Go в бинарь

```bash
mkdir -p bin
go build -o bin/hello ./cmd/hello
```

### Запуск бинаря напрямую

```bash
./bin/hello --name Alice
```

### Запуск из Python (subprocess)

```bash
python run_hello.py Alice
```

Вывод:
- результат программы — в `stdout`
- ошибки — в `stderr`
- код возврата — `returncode` (Python-скрипт возвращает его как exit code)

