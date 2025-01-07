package main

import (
	"fmt"
)



func main () {
	var symbol string
	var first float64
	var second float64
	var result float64
	var answer string

	fmt.Println("Welcome to the calculator!")
	for {
		fmt.Println("Enter the first number:")
		fmt.Scan(&first)
		fmt.Println("Enter the second number:")
		fmt.Scan(&second)
		fmt.Println("Choose an operation (+, -, *, /):")
		fmt.Scan(&symbol)
		switch symbol {
			case "+":
				result = first + second
			case "-":
				result = first - second
			case "*":
				result = first * second
			case "/":
				result = first / second
			default:
				fmt.Println("Unknown operations")
		}
		fmt.Println("Result:", first, symbol, second, "=", result)
		fmt.Println("Would you like to perform another operation? (yes to continue, no to exit):")
		fmt.Scan(&answer)
		if answer == "yes" {
			continue
		} else {
			break
		}
	}

	
}