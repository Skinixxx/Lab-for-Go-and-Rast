## Задание 8: Maturin модуль (Rust) → импорт в Python

### Требования
- Rust (`rustc`, `cargo`)
- Python 3 + `pip`
- `maturin`

### Установка maturin (рекомендуется в venv)

```bash
cd /path/to/cimlov-laba-1
python -m venv .venv
. .venv/bin/activate
python -m pip install -U pip maturin
```

### Сборка и установка в текущий Python (dev)

```bash
cd mid/maturin_task
maturin develop
```

### Проверка импорта

```bash
python -c "import fastmath; print(fastmath.sum_as_string(2,3)); print(fastmath.add(10,32)); print(fastmath.greet('Rust'))"
```

### Сборка wheel

```bash
cd mid/maturin_task
maturin build
ls -la target/wheels
```

