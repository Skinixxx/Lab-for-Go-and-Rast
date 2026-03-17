import argparse
import socket
from dataclasses import dataclass
from typing import BinaryIO


@dataclass(frozen=True)
class Reply:
    sent: str
    received: str


def _ensure_newline(line: str) -> str:
    return line if line.endswith("\n") else line + "\n"


def send_line(sock: socket.socket, line: str) -> None:
    sock.sendall(_ensure_newline(line).encode("utf-8"))


def read_line(sock_file: BinaryIO) -> str:
    data = sock_file.readline()
    if data == b"":
        raise ConnectionError("server closed connection")
    return data.decode("utf-8").rstrip("\n")


def run_session(host: str, port: int, lines: list[str], timeout_s: float = 5.0) -> list[Reply]:
    with socket.create_connection((host, port), timeout=timeout_s) as sock:
        sock.settimeout(timeout_s)
        with sock.makefile("rwb") as f:
            replies: list[Reply] = []
            for line in lines:
                send_line(sock, line)
                replies.append(Reply(sent=line, received=read_line(f)))
            return replies


def main() -> int:
    p = argparse.ArgumentParser()
    p.add_argument("--host", default="127.0.0.1")
    p.add_argument("--port", type=int, default=9090)
    p.add_argument(
        "lines",
        nargs="*",
        default=["PING", "ECHO hello from python", "QUIT"],
        help="lines to send to server (without \\n)",
    )
    args = p.parse_args()

    replies = run_session(args.host, args.port, args.lines)
    for r in replies:
        print(f"> {r.sent}")
        print(f"< {r.received}")
    return 0


if __name__ == "__main__":
    raise SystemExit(main())

