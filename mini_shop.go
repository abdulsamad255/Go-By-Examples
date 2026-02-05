package main

import (
	"fmt"
)

func main() {
	// Products and their prices
	products := map[string]float64{
		"Apple":  2,
		"Banana": 3,
		"Milk":   4,
		"Bread":  1,
	}

	fmt.Println("Welcome to Simple Shop!")
	fmt.Println("Products available:")
	for name, price := range products {
		fmt.Printf("- %s: $%.2f\n", name, price)
	}

	var total float64
	for {
		var item string
		var quantity int

		// Ask user for product
		fmt.Print("\nEnter product name (or 'done' to finish): ")
		fmt.Scan(&item)
		if item == "done" {
			break
		}

		// Check if product exists
		price, ok := products[item]
		if !ok {
			fmt.Println("Sorry, product not found!")
			continue
		}

		// Ask for quantity
		fmt.Print("Enter quantity: ")
		fmt.Scan(&quantity)

		// Add to total
		total += price * float64(quantity)
		fmt.Printf("Added %d x %s. Subtotal: $%.2f\n", quantity, item, total)
	}

	fmt.Printf("\nTotal bill: $%.2f\n", total)
	fmt.Println("Thank you for shopping!")
}
