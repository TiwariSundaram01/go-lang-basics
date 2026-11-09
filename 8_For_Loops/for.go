package main

import "fmt"

func main() {
	// 1. Classic for loop (init; condition; post)
	for i := 1; i <= 5; i++ {
		fmt.Print(i, " ")
	}
	fmt.Println()

	// 2. While-style loop (Go has no 'while' keyword, uses 'for' with single condition)
	n := 1
	for n <= 3 {
		fmt.Println("while-style:", n)
		n++
	}

	// 3. Infinite loop (uses 'break' to exit)
	count := 0
	for {
		fmt.Println("infinite loop, count:", count)
		count++
		if count == 3 {
			break // exit loop
		}
	}

	// 4. continue (skip current iteration)
	for i := 1; i <= 6; i++ {
		if i%2 != 0 {
			continue // skip odd numbers
		}
		fmt.Println("even number:", i)
	}

	// 5. range loop over integer (introduced in Go 1.22)
	for i := range 3 {
		fmt.Println("range int:", i) // prints 0, 1, 2
	}

	// Quick notes:
	// - Go has only ONE looping construct: 'for' (no while, do-while)
	// - 'break' terminates the loop
	// - 'continue' jumps directly to the next iteration
}

