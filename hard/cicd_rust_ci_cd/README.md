# CI/CD для Rust-модуля с публикацией на PyPI

[![CI](https://github.com/YOUR_USERNAME/YOUR_REPO/actions/workflows/CI.yml/badge.svg)](https://github.com/YOUR_USERNAME/YOUR_REPO/actions/workflows/CI.yml)
[![Lint](https://github.com/YOUR_USERNAME/YOUR_REPO/actions/workflows/lint.yml/badge.svg)](https://github.com/YOUR_USERNAME/YOUR_REPO/actions/workflows/lint.yml)
[![Release](https://github.com/YOUR_USERNAME/YOUR_REPO/actions/workflows/release.yml/badge.svg)](https://github.com/YOUR_USERNAME/YOUR_REPO/actions/workflows/release.yml)
[![PyPI version](https://img.shields.io/pypi/v/fastmath)](https://pypi.org/project/fastmath/)

## Структура

```
hard/cicd_rust_ci_cd/
├── .github/
│   └── workflows/
│       ├── CI.yml          # Сборка и тесты
│       ├── lint.yml        # Линтинг workflow-файлов
│       └── release.yml     # Публикация на PyPI
├── .yamllint               # Конфиг yamllint
└── README.md
```

## Workflows

### CI.yml
- **test** — Arch Linux контейнер: maturin develop + pytest
- **build-wheels** — собирает wheels для ubuntu/windows/macos (x86_64, aarch64)

### lint.yml
- **actionlint** — проверка синтаксиса GitHub Actions
- **yamllint** — проверка YAML-файлов

### release.yml
- **build** — собирает wheels
- **test-publish** — публикует на TestPyPI
- **publish** — публикация на PyPI (только при пуше тега `v*`)

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
# 1. Обновить версию в pyproject.toml
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
