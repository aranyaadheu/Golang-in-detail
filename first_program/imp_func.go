package main

import "fmt"

func main() {
	fmt.Println("Welcome to the application.")

	var name string
	fmt.Println("Enter your name - ")
	fmt.Scanln(&name)

	var num1 int
	var num2 int
	fmt.Println("Enter first number - ")
	fmt.Scanln(&num1)

	fmt.Println("Enter second number - ")
	fmt.Scanln(&num2)

	sum := num1 + num2

	fmt.Println("Hello,", name)
	fmt.Println("Summation = ", sum)

	fmt.Println("Thank you for using the application")
	fmt.Println("Goodbye!")
}
