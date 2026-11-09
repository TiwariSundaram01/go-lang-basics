package main

import "fmt"

func main() {
	// Explicit type declaration
	var name string = "golang"
	var age int = 22
	var isAdult bool = true
	var price float32 = 15.6

	fmt.Println(name, age, isAdult, price)

	// Type inference (Go detects the type)
	var course = "Go Basics"
	fmt.Println(course)

	// Short-hand syntax (most common, only works inside functions)
	framework := "Gin"
	version := 1.9
	fmt.Println(framework, version)

	// Default / Zero values (when variable is declared without value)
	// string -> "", int -> 0, float -> 0, bool -> false
	var emptyStr string
	var emptyInt int
	var emptyBool bool
	fmt.Printf("%q %d %t\n", emptyStr, emptyInt, emptyBool)

	// Assign value later
	emptyStr = "hello"
	fmt.Println(emptyStr)

	// Multiple variables in one line
	var a, b int = 1, 2
	c, d := "foo", "bar"
	fmt.Println(a, b, c, d)

	// Grouped declaration
	var (
		host = "localhost"
		port = 8080
	)
	fmt.Println(host, port)

	// Quick notes:
	// - Unused variables give compile error: "declared and not used"
	// - ':=' is for new declaration, '=' is for reassigning
	// - Cannot change variable type after declaration
}
