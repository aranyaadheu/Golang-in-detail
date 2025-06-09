package main

import "fmt"

func add(num1 int, num2 int) int { //define output type
	sum := num1 + num2

	return sum
}

func getNums(num1 int, num2 int) (int, int) {
	sum := num1 + num2 // 30

	mul := num1 * num2 // 200

	return sum, mul
}

func printSomething() {
	fmt.Println("aranyaadheu is a madlad in solitude!")
}

func sayHello(name string) {
	fmt.Println("Wellcome to the world,", name)
}

func main() {
	printSomething()
	sayHello("aranyaadheu")
}
