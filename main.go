package main

import (
	"fmt"

	"learning-go/basics"
	"learning-go/flow"
	"learning-go/pointers"
)

func main() {
	fmt.Println("main")

	x, y := basics.QuotientWithRemainder(12, 10)

	fmt.Printf("result %v, rest %v \n", x, y)

	fmt.Printf("sum equals %v\n", basics.Sum(2, 2))

	flow.Loops()

	flow.Conditionals()

	user := pointers.CreateUser("Anderson", "Major")

	fmt.Println(user)

	pointers.UpdateUser(&user, "Anderson Malheiro", "Andim")

	fmt.Printf("Updated user: %+v\n", user)

}
