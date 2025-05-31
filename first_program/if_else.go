package main

import "fmt"

func main() {
	age := 16

	if age > 18 {
		fmt.Println("You're eligible to be married")
	} else if age < 18 {
		fmt.Println("You're not eligible to be married, but you can love someone.")
	} else {
		fmt.Println("You are just a teenager, not eligible to be married.")
	}
}
