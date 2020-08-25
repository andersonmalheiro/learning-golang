package basics

import "fmt"

// declaring a package scope variable
var x int

// Pi is a declaration of a const
const Pi = 3.14

// Functions with lowercase names are not exported by the package
func basics() {
	fmt.Println("I'm a simple function")
}

// Sum returns the sum of two integers
// Functions with uppercase names are exported by the package
// the return type is before function braces
// if the params share the same type we can declare only the last param type
func Sum(x, y int) int {
	return x + y
}

// Exported functions must have a comment explaining about their

// QuotientWithRemainder returns the value of the division and the reminder
func QuotientWithRemainder(dividend int, divisor int) (result int, remainder int) { // Function with multiple results
	result = int(dividend / divisor) // this is a integer cast
	remainder = int(dividend % divisor)

	return // function with named returns doesn't need to specify the params returned
}
