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

func main() {
	a := 10
	b := 20

	p, q := getNums(a, b)

	fmt.Println(p)
	fmt.Println(q)
}
