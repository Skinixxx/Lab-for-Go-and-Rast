package main

import (
	"math"
	"sync"
)

func TrapezoidRule(funcStr string, a, b float64, n int64) IntegrationResult {
	if n <= 0 {
		n = 10000
	}
	h := (b - a) / float64(n)
	sum := evaluateFunction(funcStr, a) + evaluateFunction(funcStr, b)

	var mu sync.Mutex
	var wg sync.WaitGroup
	chunkSize := int64(1000)
	results := make([]float64, (n-1+chunkSize-1)/chunkSize)

	for i := int64(0); i < n-1; i += chunkSize {
		wg.Add(1)
		go func(start, end int64) {
			defer wg.Done()
			localSum := 0.0
			if end > n-1 {
				end = n - 1
			}
			for j := start; j < end; j++ {
				x := a + float64(j+1)*h
				localSum += evaluateFunction(funcStr, x)
			}
			mu.Lock()
			results[start/chunkSize] = localSum
			mu.Unlock()
		}(i, i+chunkSize)
	}
	wg.Wait()

	for _, v := range results {
		sum += v
	}

	result := h * sum
	secondDerivativeMax := estimateSecondDerivativeMax(funcStr, a, b)
	errorEstimate := (math.Pow(b-a, 3) / (12 * float64(n*n))) * secondDerivativeMax

	return IntegrationResult{
		Result:        result,
		ErrorEstimate: math.Abs(errorEstimate),
		Partitions:    n,
	}
}

func estimateSecondDerivativeMax(funcStr string, a, b float64) float64 {
	h := (b - a) / 1000
	maxVal := 0.0

	var mu sync.Mutex
	var wg sync.WaitGroup
	chunkSize := 100

	for i := 0; i < 1000; i += chunkSize {
		wg.Add(1)
		go func(start, end int) {
			defer wg.Done()
			localMax := 0.0
			for j := start; j < end && j < 1000; j++ {
				x := a + float64(j)*h
				secondDeriv := estimateSecondDerivative(funcStr, x, h)
				if math.IsNaN(secondDeriv) || math.IsInf(secondDeriv, 0) {
					secondDeriv = 0
				}
				if math.Abs(secondDeriv) > localMax {
					localMax = math.Abs(secondDeriv)
				}
			}
			mu.Lock()
			if localMax > maxVal {
				maxVal = localMax
			}
			mu.Unlock()
		}(i, i+chunkSize)
	}
	wg.Wait()

	return maxVal
}

func estimateSecondDerivative(funcStr string, x, h float64) float64 {
	return (evaluateFunction(funcStr, x+h) - 2*evaluateFunction(funcStr, x) + evaluateFunction(funcStr, x-h)) / (h * h)
}
