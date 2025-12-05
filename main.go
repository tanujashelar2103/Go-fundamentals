package main

import "fmt"

func main() {
	var name string
	var age int

	fmt.Println("Enter your name")
	fmt.Scanln(&name)

	fmt.Println("Enter your age")
	fmt.Scanln(&age) // <-- fixed

	fmt.Println("Hello", name, "your age is", age)
}
