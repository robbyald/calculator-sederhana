package calculator

import "errors"

// Calculator struct
type Calculator struct{}

// Add method adds two numbers
func (c Calculator) Add(a, b float64) float64 {
	return a + b
}

// Subtract method subtracts two numbers
func (c Calculator) Subtract(a, b float64) float64 {
	return a - b
}

// Multiply method multiplies two numbers
func (c Calculator) Multiply(a, b float64) float64 {
	return a * b
}

// Divide method divides two numbers
func (c Calculator) Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	}
	return a / b, nil
}
