// Package calculator provides a library for simple calculations in Go.
package calculator

import (
  "errors"
  "fmt"
  "math"
)

// Add takes two numbers and returns the result of adding them together.
func Add(a, b float64) float64 {
	return a + b
}

// Subtract takes two numbers and returns the result of subtracting the second
// from the first.
func Subtract(a, b float64) float64 {
	return a - b
}

// Multiply takes two numbers as parameters, and returns a single number representing the result.
func Multiply(a, b float64) float64 {
	return a * b
}

// Divide takes two numbers as parameters, and returns first divide by second representing the result.
func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero is undefined")
	}
	return a / b, nil
}

// Sqrt returns the square root of a.
func Sqrt(a float64) (float64, error) {
	if a < 0 {
		return 0, fmt.Errorf("%f: Sqrt of negative number is undefined", a)
	}
	return math.Sqrt(a), nil
}
