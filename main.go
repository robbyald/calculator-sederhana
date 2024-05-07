package main

import (
	Calculator "calculator-sederhana/calculator"
	"fmt"
)

func main() {

	// Create an instance of the Calculator struct
	calc := Calculator.Calculator{}

	// Use the methods of the Calculator struct
	add := calc.Add(10, 5)
	fmt.Println("Add:", add) // Output: Add: 15

	subtract := calc.Subtract(10, 5)
	fmt.Println("Subtract:", subtract) // Output: Subtract: 5

	multiply := calc.Multiply(10, 5)
	fmt.Println("Multiply:", multiply) // Output: Multiply: 50

	divide, err := calc.Divide(10, 5)
	if err != nil {
		fmt.Println("Error:", err) // Output: Error: <nil>
	} else {
		fmt.Println("Divide:", divide) // Output: Divide: 2
	}
}
