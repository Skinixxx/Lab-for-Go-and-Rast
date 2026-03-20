package main

import (
	"math"
	"strings"
)

func evaluateFunction(funcStr string, x float64) float64 {
	switch {
	case strings.HasPrefix(funcStr, "sin"):
		return math.Sin(x)
	case strings.HasPrefix(funcStr, "cos"):
		return math.Cos(x)
	case strings.HasPrefix(funcStr, "exp"):
		return math.Exp(x)
	case strings.HasPrefix(funcStr, "x^2"):
		return x * x
	case strings.HasPrefix(funcStr, "x^3"):
		return x * x * x
	case funcStr == "x":
		return x
	default:
		return 0
	}
}
