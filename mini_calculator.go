package main

import (
	"fmt"
)

func main() {
	var num1, num2 float64
	var operator string

	fmt.Println("Welcome to Mini Calculator!")

	// Get first number
	fmt.Print("Enter first number: ")
	fmt.Scan(&num1)

	// Get operator
	fmt.Print("Enter operator (+, -, *, /): ")
	fmt.Scan(&operator)

	// Get second number
	fmt.Print("Enter second number: ")
	fmt.Scan(&num2)

	// Perform calculation
	switch operator {
	case "+":
		fmt.Printf("%.2f + %.2f = %.2f\n", num1, num2, num1+num2)
	case "-":
		fmt.Printf("%.2f - %.2f = %.2f\n", num1, num2, num1-num2)
	case "*":
		fmt.Printf("%.2f * %.2f = %.2f\n", num1, num2, num1*num2)
	case "/":
		if num2 != 0 {
			fmt.Printf("%.2f / %.2f = %.2f\n", num1, num2, num1/num2)
		} else {
			fmt.Println("Error: Division by zero!")
		}
	default:
		fmt.Println("Invalid operator!")
	}
}
