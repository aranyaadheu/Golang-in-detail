package main

import "fmt"

func main() {
	num := 2

	switch num {
	case 1:
		fmt.Println("One")
		fallthrough
	case 2:
		fmt.Println("Two") // this also runs if num == 1
	}
}
