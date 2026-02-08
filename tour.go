// HELLO, 世界!
package main // An executable Go program must use the 'main' package. This package must contain a function named 'main', which serves as the program's entry point.

// A package is a collection of one or more Go source files that are compiled together. Multiple files can belong to the same package.

// When an executable program is run, execution begins with the 'main' function in the 'main' package.

import "fmt" // 'fmt' is part of Go's standard library and provides functions for formatted input and output, such as 'Println' and 'Printf'.

func main() {
	fmt.Println("Hello, 世界!")
}
