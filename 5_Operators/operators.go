package main

import "fmt"

func main() {
	// Arithmetic operators
	a, b := 10, 3
	fmt.Println(a + b)                   // 13
	fmt.Println(a - b)                   // 7
	fmt.Println(a * b)                   // 30
	fmt.Println(a / b)                   // 3 (integer division drops decimal)
	fmt.Println(float64(a) / float64(b)) // 3.3333333333333335
	fmt.Println(a % b)                   // 1 (remainder)

	// Increment / Decrement (statements only, not expressions)
	x := 5
	x++ // x = 6
	x-- // x = 5
	fmt.Println(x)

	// Comparison operators (returns true/false)
	fmt.Println(a == b) // false
	fmt.Println(a != b) // true
	fmt.Println(a > b)  // true
	fmt.Println(a < b)  // false
	fmt.Println(a >= b) // true
	fmt.Println(a <= b) // false

	// Logical operators
	isLoggedIn := true
	isAdmin := false
	fmt.Println(isLoggedIn && isAdmin) // false (AND)
	fmt.Println(isLoggedIn || isAdmin) // true  (OR)
	fmt.Println(!isLoggedIn)           // false (NOT)

	// Assignment operators
	num := 10
	num += 5 // num = 15
	num -= 3 // num = 12
	num *= 2 // num = 24
	num /= 4 // num = 6
	fmt.Println(num)

	// Quick notes:
	// - No ternary operator in Go (no condition ? a : b)
	// - Cannot mix types: int + float requires explicit conversion (e.g. float64(a) + f)
}
