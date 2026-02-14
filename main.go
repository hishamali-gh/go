// VARIABLES
package main

import "fmt"

func main() {
	// ZERO VALUES:
	var name string // ""
	var age int // 0
	var score float64 // 0
	var isAlive bool // false

	fmt.Println(name, age, score, isAlive)

	// DECLARING AND INITIALIZING AT THE SAME TIME:
	var greet string = "Hi!"

	fmt.Println(greet)

	// TYPE INFERENCE:
	// Go can figure out the type from the value.

	var shoot = "POP!" // string
	var n = 1 // int
	var percentage = 9.11 // float64
	var honest = true // bool

	fmt.Printf("%T %T %T %T\n", shoot, n, percentage, honest)

	// SHORT VARIABLE DECLARATION:
	poco := "loco" // Shorter way of 'var poco = "loco"'

	// NB: ":=" can only be used inside a function.

	fmt.Printf("%s - %T\n", poco, poco)	
}
