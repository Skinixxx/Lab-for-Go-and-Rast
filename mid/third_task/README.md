## Лабораторная: Go бинарь + вызов из Python (subprocess)

### Структура
- `mid/third_task/main.go`: Go-программа (CLI)
- `bin/hello`: собранный бинарь (создаётся командой build или Python-скриптом)
- `mid/third_task/run_hello.py`: Python-скрипт, который при необходимости собирает Go-бинарь и запускает его через `subprocess`

### Сборка Go в бинарь

```bash
mkdir -p bin
go build -o mid/third_task/bin/hello mid/third_task/main.go
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

