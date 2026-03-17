import unittest
from unittest.mock import patch

import run_hello


class TestRunHello(unittest.TestCase):
    @patch("run_hello.subprocess.run")
    def test_build_go_binary_invokes_go_build(self, mock_run):
        run_hello.build_go_binary()

        mock_run.assert_called_once()
        args, kwargs = mock_run.call_args

        self.assertEqual(args[0][0:3], ["go", "build", "-o"])
        self.assertEqual(args[0][3], str(run_hello.HELLO_BIN))
        self.assertEqual(args[0][4], ".")

        self.assertEqual(kwargs["cwd"], str(run_hello.ROOT))
        self.assertTrue(kwargs["check"])
        self.assertTrue(kwargs["text"])

    @patch("run_hello.subprocess.run")
    def test_run_hello_invokes_binary(self, mock_run):
        mock_run.return_value = unittest.mock.Mock(
            returncode=0, stdout="hello, Alice\n", stderr=""
        )

        result = run_hello.run_hello("Alice")

        mock_run.assert_called_once()
        args, kwargs = mock_run.call_args

        self.assertEqual(args[0][0], str(run_hello.HELLO_BIN))
        self.assertEqual(args[0][1:], ["--name", "Alice"])
        self.assertEqual(kwargs["cwd"], str(run_hello.ROOT))
        self.assertTrue(kwargs["capture_output"])
        self.assertTrue(kwargs["text"])
        self.assertFalse(kwargs["check"])
        self.assertEqual(result.stdout, "hello, Alice\n")


if __name__ == "__main__":
    unittest.main()

