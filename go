package main

import "fmt"

func main() {
	var num1, num2 float64
	var op string

	// Add an input prompt
	fmt.Print("Enter an expression (e.g., 3 + 4): ")
	fmt.Scan(&num1, &op, &num2)

	switch op {
	case "+":
		fmt.Println("Result:", num1+num2)
	case "-":
		fmt.Println("Result:", num1-num2)
	case "*":
		fmt.Println("Result:", num1*num2)
	case "/":
		if num2 != 0 {
			fmt.Println("Result:", num1/num2)
		} else {
			fmt.Println("Error: Division by zero")
		}
	default:
		fmt.Println("Invalid operator")
	}
}

