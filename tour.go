// HELLO, 世界!
/* package main // An executable Go program must use the 'main' package. This package must contain a function named 'main', which serves as the program's entry point.

// A package is a collection of one or more Go source files that are compiled together. Multiple files can belong to the same package.

// When an executable program is run, execution begins with the 'main' function in the 'main' package.

import "fmt" // 'fmt' is part of Go's standard library and provides functions for formatted input and output, such as 'Println' and 'Printf'.

func main() {
	fmt.Println("Hello, 世界!")
} */

// PRINT
/* package main

import "fmt"

func main() {
	a, b := "One", "Two" // Declares and initializes the variables 'a' and 'b' using 'short variable declaration syntax'. Their types are inferred from the assigned values. This will be explained in detail later.

	print(a, b) // The built-in 'print' function writes its arguments without inserting spaces or a newline between them.

	println(a, b) // The 'println' function separates its arguments with spaces and adds a newline at the end. Like 'print', 'println' is a built-in Go function, so it does not require importing 'fmt'. However, 'fmt.Println' is generally preferred for ordinary output.

	fmt.Println("This is what 'fmt.Println' should be used for:", a, b) // 'Println' prints its arguments, separating them with spaces, and adds a newline at the end.

	fmt.Printf("This is what 'fmt.Printf' should be used for: %v, %v", b, a) // 'Printf' prints text according to a format string. It is somewhat similar to Python's f-strings, although it uses format verbs such as '%v' instead of embedded expressions. '%v' formats a value using its default representation. The supplied arguments are substituted into the format string in order.
} */

// FUNCTIONS
package main

import "fmt"

// A function can take zero or more arguments.

// func add(x int, y int) int { // Notice that the type comes after the variable names.
// 	return x + y
// }

func add(x, y int) int { // When two or more consecutive named function parameters share a type, you can omit the type from all but the last.
	return x + y
}

func main() {
	fmt.Println(add(42, 13))
}
