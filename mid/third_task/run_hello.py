import os
import subprocess
import sys
from pathlib import Path


ROOT = Path(__file__).resolve().parent
BIN_DIR = ROOT / "bin"
HELLO_BIN = BIN_DIR / "hello"


def build_go_binary() -> None:
    BIN_DIR.mkdir(exist_ok=True)
    subprocess.run(
        ["go", "build", "-o", str(HELLO_BIN), "."],
        cwd=str(ROOT),
        check=True,
        text=True,
    )


def run_hello(name: str) -> subprocess.CompletedProcess[str]:
    env = os.environ.copy()
    return subprocess.run(
        [str(HELLO_BIN), "--name", name],
        cwd=str(ROOT),
        capture_output=True,
        text=True,
        env=env,
        check=False,
    )


def main(argv: list[str]) -> int:
    name = argv[1] if len(argv) > 1 else "Alice"

    if not HELLO_BIN.exists():
        build_go_binary()

    result = run_hello(name)
    sys.stdout.write(result.stdout)
    if result.stderr:
        sys.stderr.write(result.stderr)

    return result.returncode


if __name__ == "__main__":
    raise SystemExit(main(sys.argv))

