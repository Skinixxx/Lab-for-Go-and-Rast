## Лабораторная: Go бинарь + вызов из Python (subprocess)

### Структура
- `mid/third_task/main.go`: Go-программа (CLI)
- `mid/third_task/bin/hello`: собранный бинарь (создаётся командой build или Python-скриптом)
- `mid/third_task/run_hello.py`: Python-скрипт, который при необходимости собирает Go-бинарь и запускает его через `subprocess`

### Сборка Go в бинарь

```bash
mkdir -p mid/third_task/bin
cd mid/third_task
go build -o bin/hello .
```

### Запуск бинаря напрямую

```bash
cd mid/third_task
./bin/hello --name Alice
```

### Запуск из Python (subprocess)

```bash
cd mid/third_task
python run_hello.py Alice
```

### Unit-тесты

Go:

```bash
cd mid/third_task
go test ./...
```

Python:

```bash
cd mid/third_task
python -m unittest
```

Вывод:
- результат программы — в `stdout`
- ошибки — в `stderr`
- код возврата — `returncode` (Python-скрипт возвращает его как exit code)

