# Lab-for-Go-and-Rast

Лабораторные работы по интеграции Go/Rust с Python.

## Проекты

### Mid (учебные задания)

| Задание | Описание | Технологии |
|---------|----------|------------|
| [third_task](mid/third_task/) | Go бинарь + вызов из Python (subprocess) | Go, Python |
| [fourth_task](mid/fourth_task/) | Go TCP-сервер + Python TCP-клиент | Go, Python |
| [maturin_task](mid/maturin_task/) | Maturin модуль (Rust) → импорт в Python | Rust, Maturin |

### Hard (продвинутые)

| Задание | Описание | Технологии |
|---------|----------|------------|
| [cicd_rust_ci_cd](hard/cicd_rust_ci_cd/) | CI/CD для Rust-модуля + публикация на PyPI | Rust, GitHub Actions, PyPI |
| [go_microservice](hard/go_microservice/) | Go микросервис (HTTP/gRPC) + Python клиент | Go, gRPC, Python, Docker |

---

## Быстрый старт

### third_task — Go бинарь из Python

```bash
python mid/third_task/run_hello.py Alice
```

### fourth_task — TCP клиент-сервер

```bash
# Терминал 1: сервер
cd mid/fourth_task && go run . --addr 127.0.0.1:9090

# Терминал 2: клиент
cd mid/fourth_task && python client.py --host 127.0.0.1 --port 9090
```

### maturin_task — Rust модуль для Python

```bash
cd mid/maturin_task
maturin develop
python -c "import fastmath; print(fastmath.add(2, 3))"
```

### cicd_rust_ci_cd — CI/CD пайплайн

Документация: [hard/cicd_rust_ci_cd/README.md](hard/cicd_rust_ci_cd/README.md)

### go_microservice — Микросервис с HTTP/gRPC

```bash
# Go сервис
cd hard/go_microservice/go-service && ./service

# Python клиент
cd hard/go_microservice/python-client && python client.py
```

---

## Тесты

| Проект | Команда |
|--------|---------|
| third_task (Go) | `cd mid/third_task && go test ./...` |
| third_task (Python) | `cd mid/third_task && python -m unittest` |
| fourth_task (Go) | `cd mid/fourth_task && go test ./...` |
| fourth_task (Python) | `cd mid/fourth_task && python -m unittest` |
| maturin_task (Rust) | `cd mid/maturin_task && cargo test` |
| maturin_task (Python) | `cd mid/maturin_task && pytest python/tests/ -v` |
| go_microservice (Go) | `cd hard/go_microservice/go-service && go test -v .` |
| go_microservice (Python) | `cd hard/go_microservice/python-client && python -m unittest test_client -v` |

---

## Структура репозитория

```
.
├── mid/                          # Учебные задания
│   ├── third_task/              # Go → Python (subprocess)
│   ├── fourth_task/             # Go TCP-сервер + Python клиент
│   └── maturin_task/            # Rust → Python (maturin)
├── hard/                         # Продвинутые задания
│   ├── cicd_rust_ci_cd/         # CI/CD + PyPI публикация
│   └── go_microservice/         # Go HTTP/gRPC + Python
├── PROMT_LOG.md                  # Лог промптов
└── GO_CHEATSHEET.md             # Шпаргалка по Go
```
