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
}
