# Prompt Log
## Задание M3: Go → бинарь → вызов из Python (subprocess)
### Промпт 1
**Инструмент:** Cursor AI Agent (GPT-5.2)
**Промпт:** "Скомпилировать Go-программу в бинарь и вызвать её из Python (subprocess). Объясни как сделать, что использовать, и напиши код-примеры."
**Результат:** Подготовлены шаги сборки Go-бинаря, пример Go-программы (stdin/stdout + exit codes), и пример Python-скрипта для вызова через subprocess с обработкой stdout/stderr/кодов возврата.
### Промпт 2
**Инструмент:** Cursor AI Agent (GPT-5.2)
**Промпт:** "Давай теперь напишем юнит-тесты для проверки нашего приложение и учти мои изменения в расположении файлов"
**Результат:** Добавлены Go unit-тесты для логики приветствия (`greet.go`, `greet_test.go`), обновлён `main.go` для использования тестируемой функции, исправлена сборка Go из Python (build target `.`), добавлены Python unit-тесты для `run_hello.py` (моки `subprocess.run`), обновлён `mid/third_task/README.md` с актуальными путями и командами запуска тестов.
### Промпт 3
**Инструмент:** Cursor AI Agent (GPT-5.2)
**Промпт:** "Проведи код-ревью"
**Результат:** Проведено ревью и применены правки: исправлен путь запуска бинаря в `mid/third_task/README.md`; в Go добавлена “сентинельная” ошибка `ErrEmptyName` и тест обновлён на проверку через `errors.Is`, при этом пользовательское сообщение об ошибке CLI сохранено в `main.go`. Тесты `go test ./...` и `python -m unittest` проходят.
### Итого
- Количество промптов: 3
- Что пришлось исправлять вручную: ничего
- Время: ~30-45 мин

---
## Задание M5: Go TCP-сервер + Python TCP-клиент
### Промпт 1
**Инструмент:** Cursor AI Agent (GPT-5.2)
**Промпт:** "Реализовать на Go простой TCP-сервер, к которому подключается Python-клиент. Давай составим сперва план реализации"
**Результат:** Составлен план реализации: выбран построчный протокол (команды `PING`, `ECHO <text>`, `QUIT`), намечена структура `mid/fourth_task/` (Go сервер + Python клиент + README), описаны шаги запуска/проверки и критерии готовности (подключение клиента, корректные ответы, закрытие соединения).
### Промпт 2
**Инструмент:** Cursor AI Agent (GPT-5.2)
**Промпт:** "Мне план понравился, давай его реализуем."
**Результат:** Реализованы `mid/fourth_task/server.go` (TCP-сервер на Go с командами PING/ECHO/QUIT), `mid/fourth_task/client.py` (Python TCP-клиент), `mid/fourth_task/README.md` (инструкция запуска), `mid/fourth_task/go.mod`. Проверено локально: клиент получает ответы `PONG`, `ECHO: ...`, `BYE`.

### Промпт 3
**Инструмент:** Cursor AI Agent (GPT-5.2)
**Промпт:** "Реализуй юнит-тестирование"
**Результат:** Добавлены unit-тесты: `mid/fourth_task/server_test.go` (Go: тесты `processLine` для PING/ECHO/QUIT/unknown) и `mid/fourth_task/test_client.py` (Python: тесты `send_line`, `read_line`, `run_session` с моками сокета). Проверено: `go test ./...` и `python -m unittest` проходят.

### Промпт 4
**Инструмент:** Cursor AI Agent (GPT-5.2)
**Промпт:** "Проведи рефакторинг-кода"
**Результат:** Проведён рефакторинг без изменения протокола: добавлен `mid/fourth_task/protocol.go` (вынесена логика `processLine` и ошибки `ErrEmptyLine`/`ErrUnknownCommand`), упрощён `server.go` (буфер Scanner, таймауты чтения/записи, единый ответ `ERR ...`), обновлены Go-тесты (проверка `errors.Is` и тест на пустую строку). В Python добавлена `_ensure_newline`, уточнены типы, обновлены тесты. Проверено: `go test ./...` и `python -m unittest` проходят.

### Итого
- Количество промптов: 4
- Что пришлось исправлять вручную: ничего
- Время: ~30 мин

---
## Задание M8: Maturin (Rust) модуль → импорт в Python
### Промпт 1
**Инструмент:** Cursor AI Agent (GPT-5.2)
**Промпт:** "Следующее наше задание 8. Собрать модуль с Maturin и импортировать в Python. Давай проведем планирование"
**Результат:** Составлен план: проверить Rust/Python окружение, создать проект через `maturin init` (PyO3), реализовать 1–2 функции, собрать/установить через `maturin develop`, проверить `import` и при необходимости собрать wheel через `maturin build`.
### Промпт 2
**Инструмент:** Cursor AI Agent (GPT-5.2)
**Промпт:** "Меня все устраивает реализовываем"
**Результат:** Установлен `maturin` в venv, создан проект `mid/maturin_task` (PyO3). Реализованы функции `sum_as_string`, `add`, `greet` в Rust (`src/lib.rs`), собран и установлен модуль через `maturin develop`, проверено `import fastmath` и вызов функций из Python. Добавлен `mid/maturin_task/README.md`.
### Промпт 3
**Инструмент:** Cursor AI Agent (GPT-5.2)
**Промпт:** "Давай теперь сделаем юнит-тесты"
**Результат:** Добавлены unit-тесты на Rust и Python: в `src/lib.rs` вынесены чистые функции `*_impl` и добавлены `cargo test`; в `python/tests/` добавлены/обновлены тесты на `unittest` (`test_all.py`, `test_fastmath_unittest.py`). Проверено: `cargo test` и `python -m unittest discover -s mid/maturin_task/python/tests` проходят.
### Промпт 4
**Инструмент:** Cursor AI Agent (GPT-5.2)
**Промпт:** "Проведи рефакторинг кода"
**Результат:** Рефакторинг без изменения API: в `src/lib.rs` сделана более идиоматичная структура PyO3 (`#[pymodule] fn fastmath(...)` + `#[pyfunction]` на верхнем уровне), при этом сохранены функции `sum_as_string`, `add`, `greet`. В Python-тестах убран дубликат `test_all.py`, оставлен единый `unittest`-набор. Проверено: `cargo test` и `python -m unittest discover -s mid/maturin_task/python/tests` проходят.
### Итого
- Количество промптов: 4
- Что пришлось исправлять вручную: ничего
- Время: ~15–25 мин

---

## Задание B3: CI/CD для Rust-модуля + публикация на PyPI
### Промпт 1
**Инструмент:** Cursor AI Agent (GPT-5.2)
**Промпт:** "Настроить сборку Rust-модуля в CI/CD и публикацию на PyPI. Давай спланируем наши действия, накидай идеи и способы реализации."
**Результат:** Составлен план реализации:
- **Workflow 1: `ci.yml`** — на push/PR: matrix build (linux/macos/windows), шаги checkout/setup-python/install-rust/maturin build/upload artifacts, job test: maturin develop + pytest
- **Workflow 2: `release.yml`** — на тег `v*`: job testpypi (maturin publish --repository testpypi), job pypi (maturin publish, только с main)
- **Secrets:** PYPI_TEST_API_TOKEN, PYPI_API_TOKEN
- **Реестры:** TestPyPI + PyPI; артефакты: только wheel

### Промпт 2
**Инструмент:** Cursor AI Agent (GPT-5.2)
**Промпт:** "Давай приступим к реализации"
**Результат:** 
- Восстановлен `mid/maturin_task/.github/workflows/CI.yml` (оригинальный автогенерированный maturin) и `pyproject.toml` (dynamic version)
- Создан `hard/cicd_rust_ci_cd/` с базовой реализацией:
  - `.github/workflows/CI.yml` — job `test` (maturin develop + pytest), job `build-wheels` (ubuntu/windows/macos)
  - `.github/workflows/release.yml` — build → test-publish (TestPyPI) → publish (PyPI)
  - `README.md` с инструкциями по настройке secrets и релизу

### Промпт 3
**Инструмент:** Cursor AI Agent (GPT-5.2)
**Промпт:** "Теперь объясни что ты реализовал"
**Результат:** Добавлены подробные комментарии к каждой секции CI.yml и release.yml (триггеры, permissions, jobs, steps, matrix, условия if)

### Промпт 4
**Инструмент:** Cursor AI Agent (GPT-5.2)
**Промпт:** "можно на арч заменить если да то как?"
**Результат:** Заменён job `test` на Arch Linux через контейнер (`container: archlinux:latest` + pacman для установки python/rust/cargo). Для wheel-сборки Arch оставлен как ubuntu/windows/macos, т.к. GitHub не имеет готовых Arch-раннеров.

### Промпт 5
**Инструмент:** Cursor AI Agent (GPT-5.2)
**Промпт:** "Так давай продолжим нашу разработку, пока меня все устраивает"
**Результат:** Добавлен `lint.yml` — job `actionlint` (валидация GitHub Actions синтаксиса) и `yamllint` (валидация YAML). Добавлен `.yamllint` конфиг. Триггер: только при изменениях в `.github/workflows/`.

### Промпт 6
**Инструмент:** Cursor AI Agent (GPT-5.2)
**Промпт:** "Добавь комментарии в код"
**Результат:** Добавлены подробные комментарии в `lint.yml` (назначение actionlint/yamllint, что проверяют) и `.yamllint` (правила конфигурации yamllint). Объяснено различие между actionlint (специфика GitHub Actions) и yamllint (общий YAML).

### Промпт 7
**Инструмент:** Cursor AI Agent (GPT-5.2)
**Промпт:** "Давай пока проверим наш результат и тесты. Если все удачно: реализуй баджейки"
**Результат:** Добавлены промпты 5-6 в промтлог. yamllint выявил trailing spaces — исправлены. Обновлён `.yamllint` (отключены `document-start` и `truthy` warnings). Добавлены баджейки в README: CI, Lint, Release, PyPI version.

### Промпт 8
**Инструмент:** Cursor AI Agent (GPT-5.2)
**Промпт:** "Кстати а где наш Rust-модуль? давай реализуем новый минимальный модуль"
**Результат:** Создан минимальный Rust-модуль в `hard/cicd_rust_ci_cd/`: `pyproject.toml` (maturin build-system), `Cargo.toml` (pyo3 0.22), `src/lib.rs` (функции `add`, `multiply`, `greet`), `python/fastmath/__init__.py`, `python/tests/test_fastmath.py`. Собран с `PYO3_USE_ABI3_FORWARD_COMPATIBILITY=1` (Arch Python 3.14). pytest проходит: 3 теста.

### Промпт 9
**Инструмент:** Cursor AI Agent (GPT-5.2)
**Промпт:** "Так я вручную проверил все работает давай dependabot и кэширование реализуем"
**Результат:**
- Создан `.github/dependabot.yml` — автообновление GitHub Actions, pip и cargo зависимостей (проверка раз в неделю)
- Обновлён `CI.yml` — добавлено кэширование pip packages, cargo registry (job test) и sccache (job build-wheels)
- Обновлён README — добавлены секции Dependabot и Кэширование
- yamllint проходит

### Итого
- Количество промптов: 9
- Что пришлось исправлять вручную: Cargo.toml, CI.yml,lint.yml, release.yml.
- Время: ~2-3 часа

---

## Задание B1: Go микросервис (HTTP/gRPC) + Python клиент
### Промпт 1
**Инструмент:** opencode
**Промпт:** "Создать микросервис на Go, который выполняет тяжёлые вычисления, и вызывать его из Python-приложения (через HTTP/gRPC)"
**Результат:** Составлен план с пунктами: структура проекта, Go-микросервис, Python-клиент, unit-тесты, Docker-compose. Выбраны численное интегрирование как тяжёлые вычисления.
### Промпт 2
**Инструмент:** opencode
**Промпт:** "Численное интегрирование производим. Реализацию пиши в hard создай там отдельную папочку и пиши"
**Результат:** Записан промпт в prompt_log.md. Создана папка hard/go_microservice/ с реализацией:
- `go-service/` — Go микросервис (HTTP :8080, gRPC :50051) с методом трапеций для численного интегрирования, поддержка функций sin/cos/exp/x/x²/x³
- `python-client/` — Python клиент (HTTP + gRPC вызовы)
- `Dockerfile.go`, `Dockerfile.python`, `docker-compose.yml`
- `go-service/main_test.go` — 12 unit-тестов (все проходят)
### Промпт 3
**Инструмент:** opencode
**Промпт:** "Давай реализуем юнит-тесты. После этого ты снова сделаешь рефакторинг кода и напишешь мне отчет по окончанию результата"
**Результат:** 
- Порты изменены на :18080 (HTTP) и :50051 (gRPC) из-за конфликта с портом 8080
- Добавлены Python unit-тесты: `python-client/test_client.py` — 11 тестов (HTTP + gRPC моки)
- Расширены Go unit-тесты: `go-service/http_test.go` — 9 HTTP-тестов
- Go тесты: 21 тест (все проходят)
- Проведён рефакторинг Go-сервиса: разделение на модули
- Структура после рефакторинга:
  - `main.go` — точка входа
  - `server.go` — gRPC/HTTP серверы
  - `trapezoid.go` — метод трапеций (с параллелизацией)
  - `functions.go` — парсинг и вычисление функций
  - `integration.go` — структура результата
  - `http_test.go` — HTTP тесты
  - `main_test.go` — тесты интегрирования
### Итого
- Количество промптов: 3
- Что пришлось исправлять вручную: решение конфликтов при push.
- Время: ~25 мин


