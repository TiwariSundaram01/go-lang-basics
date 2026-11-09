package main

import "fmt"

func main() {
	// Value cannot be changed once declared
	const name string = "golang"

	// Type inference works with const too
	const pi = 3.14

	// Grouped constants
	const (
		port = 8080
		host = "localhost"
	)

	// Multiple in one line
	const a, b = 1, 2

	// Cannot reassign
	// name = "python" // error: cannot assign to name

	// Cannot use := with const
	// const x := 10 // error!

	// iota: auto-increments from 0
	const (
		first = iota // 0
		second       // 1
		third        // 2
	)

	fmt.Println(name, pi, host, port, a, b)
	fmt.Println(first, second, third)

	// Quick notes:
	// - Must be known at compile time (cannot assign function return like time.Now())
	// - Unused constants do NOT throw compile errors
}
