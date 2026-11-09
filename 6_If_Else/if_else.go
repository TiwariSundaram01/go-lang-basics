package main

import "fmt"

func main() {
	// Basic if-else
	age := 18
	if age >= 18 {
		fmt.Println("Person is an adult")
	} else {
		fmt.Println("Person is not an adult")
	}

	// if - else if - else
	score := 85
	if score >= 90 {
		fmt.Println("Grade A")
	} else if score >= 75 {
		fmt.Println("Grade B")
	} else {
		fmt.Println("Grade C")
	}

	// Variable declaration inside if condition (scoped only inside this if-else block)
	if num := 15; num%2 == 0 {
		fmt.Println(num, "is even")
	} else {
		fmt.Println(num, "is odd")
	}
	// fmt.Println(num) // error! 'num' is not accessible outside if block

	// Logical operators in condition
	hasID := true
	isMember := false
	if hasID && isMember {
		fmt.Println("Entry allowed to lounge")
	} else if hasID || isMember {
		fmt.Println("Entry allowed to general area")
	}

	// Quick notes:
	// - No parentheses () needed around conditions: `if x > 0`
	// - Curly braces {} are mandatory even for single-line bodies
	// - Go does NOT have ternary operator: use if-else instead
}

