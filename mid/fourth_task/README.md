## Задание 2: Go TCP-сервер + Python TCP-клиент

### Протокол
- Клиент отправляет строку, оканчивающуюся `\n`
- Сервер отвечает одной строкой (тоже с `\n`)
- Поддерживаемые команды:
  - `PING` → `PONG`
  - `ECHO <text>` → `ECHO: <text>`
  - `QUIT` → `BYE` и сервер закрывает соединение

### Запуск сервера (Go)

```bash
cd mid/fourth_task
go run . --addr 127.0.0.1:9090
```

Или собрать бинарь:

```bash
cd mid/fourth_task
go build -o bin/server .
./bin/server --addr 127.0.0.1:9090
```

### Запуск клиента (Python)

```bash
cd mid/fourth_task
python client.py --host 127.0.0.1 --port 9090
```

Отправить свои команды:

```bash
cd mid/fourth_task
python client.py PING "ECHO hi" QUIT
```

### Тесты

Go:

```bash
cd mid/fourth_task
go test ./...
```

Python:

```bash
cd mid/fourth_task
python -m unittest
```

