import pytest
import fastmath


def test_sum_as_string():
    assert fastmath.sum_as_string(1, 1) == "2"
