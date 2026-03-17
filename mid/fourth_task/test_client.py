import io
import unittest
from unittest.mock import Mock, patch

import client


class TestClientUnit(unittest.TestCase):
    def test_ensure_newline_appends_when_missing(self):
        self.assertEqual(client._ensure_newline("PING"), "PING\n")

    def test_ensure_newline_keeps_existing(self):
        self.assertEqual(client._ensure_newline("PING\n"), "PING\n")

    def test_send_line_appends_newline_and_encodes_utf8(self):
        sock = Mock()
        client.send_line(sock, "PING")
        sock.sendall.assert_called_once_with(b"PING\n")

    def test_send_line_keeps_existing_newline(self):
        sock = Mock()
        client.send_line(sock, "PING\n")
        sock.sendall.assert_called_once_with(b"PING\n")

    def test_read_line_decodes_and_strips_newline(self):
        f = io.BytesIO(b"PONG\n")
        self.assertEqual(client.read_line(f), "PONG")

    def test_read_line_raises_on_eof(self):
        f = io.BytesIO(b"")
        with self.assertRaises(ConnectionError):
            client.read_line(f)

    @patch("client.socket.create_connection")
    def test_run_session_sends_all_lines_and_reads_replies(self, mock_create_connection):
        sock = Mock()
        mock_create_connection.return_value.__enter__.return_value = sock

        f = io.BytesIO(b"PONG\nECHO: hi\nBYE\n")
        makefile_cm = Mock()
        makefile_cm.__enter__ = Mock(return_value=f)
        makefile_cm.__exit__ = Mock(return_value=False)
        sock.makefile.return_value = makefile_cm

        replies = client.run_session("127.0.0.1", 9090, ["PING", "ECHO hi", "QUIT"])

        self.assertEqual([r.received for r in replies], ["PONG", "ECHO: hi", "BYE"])
        sock.sendall.assert_any_call(b"PING\n")
        sock.sendall.assert_any_call(b"ECHO hi\n")
        sock.sendall.assert_any_call(b"QUIT\n")


if __name__ == "__main__":
    unittest.main()

