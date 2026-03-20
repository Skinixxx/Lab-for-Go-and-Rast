package main

import (
	"math"
	"testing"
)

func TestTrapezoidRule_Sin(t *testing.T) {
	result := TrapezoidRule("sin", 0, math.Pi, 100000)
	expected := 2.0

	if math.Abs(result.Result-expected) > 0.001 {
		t.Errorf("Expected ~%f, got %f", expected, result.Result)
	}
}

func TestTrapezoidRule_Cos(t *testing.T) {
	result := TrapezoidRule("cos", 0, math.Pi/2, 100000)
	expected := 1.0

	if math.Abs(result.Result-expected) > 0.001 {
		t.Errorf("Expected ~%f, got %f", expected, result.Result)
	}
}

func TestTrapezoidRule_X(t *testing.T) {
	result := TrapezoidRule("x", 0, 1, 100000)
	expected := 0.5

	if math.Abs(result.Result-expected) > 0.001 {
		t.Errorf("Expected ~%f, got %f", expected, result.Result)
	}
}

func TestTrapezoidRule_XSquared(t *testing.T) {
	result := TrapezoidRule("x^2", 0, 1, 100000)
	expected := 1.0 / 3.0

	if math.Abs(result.Result-expected) > 0.001 {
		t.Errorf("Expected ~%f, got %f", expected, result.Result)
	}
}

func TestTrapezoidRule_Exp(t *testing.T) {
	result := TrapezoidRule("exp", 0, 1, 100000)
	expected := math.E - 1

	if math.Abs(result.Result-expected) > 0.001 {
		t.Errorf("Expected ~%f, got %f", expected, result.Result)
	}
}

func TestTrapezoidRule_DefaultPartitions(t *testing.T) {
	result := TrapezoidRule("sin", 0, math.Pi, 0)
	if result.Partitions != 10000 {
		t.Errorf("Expected default 10000 partitions, got %d", result.Partitions)
	}
}

func TestTrapezoidRule_ErrorEstimate(t *testing.T) {
	result := TrapezoidRule("sin", 0, math.Pi, 10000)
	if result.ErrorEstimate < 0 {
		t.Errorf("Error estimate should be positive, got %f", result.ErrorEstimate)
	}
}

func TestEvaluateFunction_Sin(t *testing.T) {
	result := evaluateFunction("sin", 0)
	if result != 0 {
		t.Errorf("sin(0) should be 0, got %f", result)
	}

	result = evaluateFunction("sin", math.Pi/2)
	if math.Abs(result-1) > 0.0001 {
		t.Errorf("sin(pi/2) should be 1, got %f", result)
	}
}

func TestEvaluateFunction_Cos(t *testing.T) {
	result := evaluateFunction("cos", 0)
	if result != 1 {
		t.Errorf("cos(0) should be 1, got %f", result)
	}

	result = evaluateFunction("cos", math.Pi)
	if math.Abs(result+1) > 0.0001 {
		t.Errorf("cos(pi) should be -1, got %f", result)
	}
}

func TestEvaluateFunction_X(t *testing.T) {
	result := evaluateFunction("x", 5)
	if result != 5 {
		t.Errorf("x at 5 should be 5, got %f", result)
	}
}

func TestEvaluateFunction_XSquared(t *testing.T) {
	result := evaluateFunction("x^2", 3)
	if result != 9 {
		t.Errorf("x^2 at 3 should be 9, got %f", result)
	}
}

func TestEvaluateFunction_Unknown(t *testing.T) {
	result := evaluateFunction("unknown", 1)
	if result != 0 {
		t.Errorf("Unknown function should return 0, got %f", result)
	}
}
