# CI/CD для Rust-модуля с публикацией на PyPI

[![CI](https://github.com/YOUR_USERNAME/YOUR_REPO/actions/workflows/CI.yml/badge.svg)](https://github.com/YOUR_USERNAME/YOUR_REPO/actions/workflows/CI.yml)
[![Lint](https://github.com/YOUR_USERNAME/YOUR_REPO/actions/workflows/lint.yml/badge.svg)](https://github.com/YOUR_USERNAME/YOUR_REPO/actions/workflows/lint.yml)
[![Release](https://github.com/YOUR_USERNAME/YOUR_REPO/actions/workflows/release.yml/badge.svg)](https://github.com/YOUR_USERNAME/YOUR_REPO/actions/workflows/release.yml)
[![PyPI version](https://img.shields.io/pypi/v/fastmath)](https://pypi.org/project/fastmath/)

## Структура

```
hard/cicd_rust_ci_cd/
├── .github/
│   ├── dependabot.yml         # Автообновление зависимостей
│   └── workflows/
│       ├── CI.yml             # Сборка и тесты
│       ├── lint.yml           # Линтинг workflow-файлов
│       └── release.yml        # Публикация на PyPI
├── Cargo.toml                 # Rust зависимости
├── pyproject.toml             # Python метаданные
├── src/lib.rs                 # Rust код
├── python/                    # Python wrapper и тесты
├── .yamllint                  # Конфиг yamllint
└── README.md
```

## Workflows

### CI.yml
- **test** — Arch Linux контейнер: maturin develop + pytest
  - Кэширование: pip packages, cargo registry
- **build-wheels** — собирает wheels для ubuntu/windows/macos (x86_64, aarch64)
  - Кэширование: sccache (результаты компиляции Rust)

### lint.yml
- **actionlint** — проверка синтаксиса GitHub Actions
- **yamllint** — проверка YAML-файлов

### release.yml
- **build** — собирает wheels
- **test-publish** — публикует на TestPyPI
- **publish** — публикация на PyPI (только при пуше тега `v*`)

## Dependabot

Автоматическое обновление зависимостей (проверка раз в неделю):

| Экосистема | Что обновляется | Файл |
|------------|----------------|------|
| GitHub Actions | `actions/*` | `.github/workflows/*.yml` |
| pip | `maturin`, `pytest` | `pyproject.toml` |
| cargo | `pyo3`, крейты | `Cargo.toml` |

## Кэширование

| Тип | Что кэшируется | Где |
|-----|----------------|-----|
| pip | Установленные Python пакеты | job `test` |
| cargo | Registry, git db | job `test` |
| sccache | Результаты компиляции Rust | job `build-wheels` |

## Настройка

### Secrets (GitHub → Settings → Secrets)

| Secret | Описание |
|--------|----------|
| `PYPI_TEST_API_TOKEN` | Токен для TestPyPI |
| `PYPI_API_TOKEN` | Токен для PyPI |

### Получить токены

1. [TestPyPI](https://test.pypi.org/manage/account/#api-tokens) — создать токен с префиксом `pypi-`
2. [PyPI](https://pypi.org/manage/account/#api-tokens) — создать токен с префиксом `pypi-`

## Релиз

```bash
# 1. Обновить версию в pyproject.toml и Cargo.toml
# 2. Закоммитить
git add -A && git commit -m "Release v0.1.0"

# 3. Создать и запушить тег
git tag v0.1.0
git push --tags
```

Поток публикации: `build` → `test-publish` (TestPyPI) → `publish` (PyPI)

## Линтинг

Проверка workflow-файлов перед коммитом:

```bash
# yamllint
pip install yamllint
yamllint -c .yamllint .github/workflows/

# actionlint (скачать бинарник)
bash <(curl -s https://raw.githubusercontent.com/rhysd/actionlint/main/scripts/download-actionlint.bash)
./actionlint
```

## Локальная сборка

```bash
cd hard/cicd_rust_ci_cd

# Arch Linux: требуется ABI3 совместимость
PYO3_USE_ABI3_FORWARD_COMPATIBILITY=1 maturin develop --release

# Другие ОС:
maturin develop --release

# Тесты
pytest python/tests/ -v
```
