package main

import "fmt"

func input() {
	var num1, num2 float64
	var operator string

	fmt.Print("Enter first number: ")
	fmt.Scanln(&num1)

	fmt.Print("Enter operator (+, -, *, /): ")
	fmt.Scanln(&operator)

	fmt.Print("Enter second number: ")
	fmt.Scanln(&num2)

	switch operator {
	case "+":
		fmt.Println(num1 + num2)
	case "-":
		fmt.Println(num1 - num2)
	case "*":
		fmt.Println(num1 * num2)
	case "/":
		if num2 != 0 {
			fmt.Println(num1 / num2)
		} else {
			fmt.Println("Division by zero error")
		}
	default:
		fmt.Println("Invalid operator")
	}
}

func main() {
	input()
}
