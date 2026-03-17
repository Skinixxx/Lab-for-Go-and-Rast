import unittest

import fastmath


class TestFastmath(unittest.TestCase):
    def test_sum_as_string(self):
        self.assertEqual(fastmath.sum_as_string(1, 1), "2")

    def test_add(self):
        self.assertEqual(fastmath.add(10, 32), 42)

    def test_greet(self):
        self.assertEqual(fastmath.greet("Rust"), "hello, Rust")


if __name__ == "__main__":
    unittest.main()

