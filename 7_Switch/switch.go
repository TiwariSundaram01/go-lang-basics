package main

import (
	"fmt"
	"time"
)

func main() {
	// 1. Basic switch (No need for 'break' - Go breaks automatically)
	day := 3
	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	default:
		fmt.Println("Other day")
	}

	// 2. Multiple values in a single case
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's weekend!")
	default:
		fmt.Println("It's a workday")
	}

	// 3. Switch without expression (clean alternative to long if-else chains)
	hour := 14
	switch {
	case hour < 12:
		fmt.Println("Good morning")
	case hour < 17:
		fmt.Println("Good afternoon")
	default:
		fmt.Println("Good evening")
	}

	// 4. fallthrough (forces execution to continue to next case)
	num := 2
	switch num {
	case 2:
		fmt.Println("Two")
		fallthrough
	case 3:
		fmt.Println("Three (because of fallthrough)")
	}

	// 5. Type switch (checks the type of an interface{} variable)
	whatType := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("It's a bool")
		case int:
			fmt.Println("It's an int")
		case string:
			fmt.Println("It's a string")
		default:
			fmt.Printf("Unknown type: %T\n", t)
		}
	}

	whatType(true)
	whatType(42)
	whatType("golang")

	// Quick notes:
	// - Go cases do NOT fall through by default (no break needed)
	// - Use 'fallthrough' only when you explicitly want to run the next case
}

