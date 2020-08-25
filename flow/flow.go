package flow

import (
	"fmt"
)

// Loops is a function to see how loops work in golang
func Loops() {
	x := 10 // shorthand variable declaration

	// Default sintax of for loop
	for i := 0; i < 10; i++ {
		fmt.Println("i")
	}

	// Infinity for loop
	// Equivalent of while true
	/*
		for {
		}
	*/

	// For statement without iterator is like a while loop
	for x > 0 {
		if x%2 == 0 {
			fmt.Printf("%v is even\n", x)
		}
		x--
	}
}

// Conditionals is a function to see how flow control statements work in golang
func Conditionals() {
	x := 5

	// defer is a statement that runs only when the function returns (end)
	// deferred function calls are pushed in a stack, so they run in a LIFO order
	defer fmt.Printf("\nEnd of flow control\n")

	// Default if statement
	if x%2 != 0 {
		fmt.Println("x is odd")
	} else {
		fmt.Println("x is even")
	}

	// If with scoped variable
	// `n` and `err` are only available inside if and else scope
	if n, err := fmt.Printf("test"); err == nil {
		fmt.Printf("%v bytes\n", n)
	} else {
		fmt.Printf("error: %v\n", err)
	}

	// Default Switch case statement
	switch x {
	case 1:
		fmt.Println("x is one")
	default:
		fmt.Println("x is five")
	}

	// Swith without condition is like `if then else` statements
	switch {
	case x > 5:
		fmt.Println("x is more than five")
	case x < 5:
		fmt.Println("x is less than five")
	default:
		fmt.Println("x is five")
	}
}
