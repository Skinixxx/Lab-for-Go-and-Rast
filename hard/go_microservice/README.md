# Go Microservice: Numerical Integration

Go-микросервис для численного интегрирования методом трапеций с HTTP и gRPC интерфейсами.

## Быстрый старт

### Локальный запуск

```bash
# Запуск Go сервиса (HTTP: :18080, gRPC: :50051)
cd go-service && go build -o service . && ./service

# В другом терминале — Python клиент
cd python-client && source venv/bin/activate && python client.py
```

### Docker

```bash
docker-compose up --build
```

## API

### HTTP

```bash
curl -X POST "http://localhost:18080/integrate?func=sin&a=0&b=3.14159&n=100000"
```

**Параметры:**
- `func` — функция: `sin`, `cos`, `exp`, `x`, `x^2`, `x^3`
- `a` — нижняя граница
- `b` — верхняя граница
- `n` — количество разбиений (по умолчанию 10000)

**Ответ:**
```json
{"result": 2.000000, "error_estimate": 2.583848e-10, "partitions": 100000}
```

### gRPC

```python
import grpc
import integration_pb2
import integration_pb2_grpc

channel = grpc.insecure_channel('localhost:50051')
stub = integration_pb2_grpc.IntegratorStub(channel)

request = integration_pb2.IntegrationRequest(
    function="sin",
    lower_bound=0,
    upper_bound=3.14159,
    partitions=100000
)
response = stub.Integrate(request)
print(f"Result: {response.result}")
```

## Тесты

```bash
# Go тесты
cd go-service && go test -v .

# Python тесты
cd python-client && source venv/bin/activate && python -m unittest test_client -v
```

## Структура проекта

```
go_microservice/
├── go-service/           # Go микросервис
│   ├── main.go           # Точка входа
│   ├── server.go         # HTTP/gRPC серверы
│   ├── trapezoid.go      # Метод трапеций (параллельный)
│   ├── functions.go      # Парсинг функций
│   ├── integration.go    # Структура результата
│   ├── proto/            # Protobuf определения
│   └── *_test.go        # Тесты
├── python-client/        # Python клиент
│   ├── client.py         # HTTP + gRPC клиент
│   └── test_client.py    # Тесты
├── Dockerfile.go         # Docker для Go
├── Dockerfile.python    # Docker для Python
└── docker-compose.yml   # Оркестрация
```
