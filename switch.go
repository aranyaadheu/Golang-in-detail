/*
Basic Syntax:

switch expression {
case value1:
    // code
case value2:
    // code
default:
    // code
}
*/

package main

import "fmt"

func main() {
	day := "Tuesday"

	switch day {
	case "Monday":
		fmt.Println("Start of the work week")
	case "Tuesday", "Wednesday", "Thursday":
		fmt.Println("Midweek days")
	case "Friday":
		fmt.Println("Last working day of the week")
	case "Saturday", "Sunday":
		fmt.Println("Weekend!")
	default:
		fmt.Println("Unknown day.")
	}
}
