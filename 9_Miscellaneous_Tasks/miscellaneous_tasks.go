package main

import "fmt"

/*
================================================================================
🎯 PRACTICE CHALLENGES: TEST YOUR GO BASICS (Concepts learned today)
================================================================================
Topics covered:
1. Variables & Zero Values
2. Constants
3. Operators & Type Conversion
4. If / Else & Short Statement
5. Switch
6. For Loops (classic, while-style, break, continue)

Write your solution code under each task inside main(),
run the file with: go run miscellaneous_tasks.go
When done, let me know and I will review it!
================================================================================
*/

func main() {
	fmt.Println("=== Go Basics Practice Tasks ===")

	// -------------------------------------------------------------------------
	// TASK 1: Variables, Constants & Type Conversion
	// -------------------------------------------------------------------------
	// 1. Declare a constant 'taxRate' = 0.18 (18% tax).
	// 2. Declare an int variable 'itemPrice' = 250 using ':='.
	// 3. Calculate total price: itemPrice + (itemPrice * taxRate).
	//    (Remember: cannot multiply int with float! Convert itemPrice to float64).
	// 4. Declare a string variable 'orderStatus' without value (test zero value).
	//    Print it using fmt.Printf("%q\n", orderStatus).
	// 5. Print the final total price.

	const taxRate = 0.18

	itemPrice := 250

	totalPrice := float64(itemPrice) + (float64(itemPrice) * taxRate)

	var orderStatus string

	fmt.Printf("%q\n", orderStatus)

	fmt.Printf("Total price: %v\n", totalPrice)

	// -------------------------------------------------------------------------
	// TASK 2: Operators & Logic
	// -------------------------------------------------------------------------
	// 1. Given:
	//    year := 2024
	// 2. A leap year is divisible by 4, but NOT by 100 (unless also divisible by 400).
	// 3. Write a single boolean expression using arithmetic and logical operators
	//    (&&, ||, !=, ==, %) to check if 'year' is a leap year.
	// 4. Print whether it is a leap year (true/false).

	year := 2024

	if year%4 == 0 && year%100 != 0 || year%400 == 0 {
		fmt.Println(true)
	} else {
		fmt.Println(false)
	}

	// -------------------------------------------------------------------------
	// TASK 3: If / Else with Short Statement
	// -------------------------------------------------------------------------
	// 1. Movie ticket check:
	//    - Age must be >= 18
	//    - Must have ticket (hasTicket = true) OR VIP pass (isVIP = true)
	// 2. Use an if statement with inline short statement:
	//    if age := 20; <condition> { ... }
	// 3. Print "c" or "Access Denied".

	hasTicket := false; isVIP := true;
	if age := 18; age >= 18 && (hasTicket || isVIP) {
		fmt.Println("Access Granted")
	} else {
		fmt.Println("Access Denied")
	}
	

	// -------------------------------------------------------------------------
	// TASK 4: Switch Statement (Tagless / Clean Conditions)
	// -------------------------------------------------------------------------
	// Given marks := 78, print the grade using switch (without expression):
	// - 90 to 100 -> "Grade: A"
	// - 75 to 89  -> "Grade: B"
	// - 50 to 74  -> "Grade: C"
	// - Below 50  -> "Grade: Fail"
	// - Less than 0 or greater than 100 -> "Invalid marks"

	switch marks := 32; {
		case marks >= 90 && marks <= 100:
			fmt.Println("Grade: A")
		case marks >= 75 && marks <= 89:
			fmt.Println("Grade: B")
		case marks >= 50 && marks <= 74:
			fmt.Println("Grade: C")
		case marks < 50:
			fmt.Println("Grade: Fail")
		default:
			fmt.Println("Invalid marks")		
	} 

	// -------------------------------------------------------------------------
	// TASK 5: For Loop with break & continue
	// -------------------------------------------------------------------------
	// Using a classic for loop:
	// - Loop i from 1 to 20.
	// - Skip numbers divisible by 3 (using 'continue').
	// - Stop the loop completely when i reaches 16 (using 'break').
	// - Print remaining numbers on a single line.

	for i := 1; i <= 20; i++ {
		if i % 3 == 0 {
			continue
		}
		if i == 16 {
			break
		}
		fmt.Print(i, " ")
	}
	fmt.Println()

	// -------------------------------------------------------------------------
	// TASK 6: While-Style Loop
	// -------------------------------------------------------------------------
	// Using for as a while loop:
	// - Start with n := 64.
	// - Divide n by 2 on every iteration (n = n / 2) while n > 1.
	// - Print n at each step.

	n := 64
	for n > 1 {
		fmt.Println(n)
		n = n / 2
	}
    fmt.Println()

	// -------------------------------------------------------------------------
	// TASK 7: FizzBuzz Challenge
	// -------------------------------------------------------------------------
	// Loop numbers 1 to 15:
	// - If divisible by both 3 and 5, print "FizzBuzz"
	// - If divisible by 3, print "Fizz"
	// - If divisible by 5, print "Buzz"
	// - Otherwise, print the number

	for i := 1; i <= 15; i++ {
		if i % 3 == 0 && i % 5 == 0 {
			fmt.Println("FizzBuzz")
		} else if i % 3 == 0 {
			fmt.Println("Fizz")
		} else if i % 5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}
