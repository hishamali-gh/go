// HELLO, 世界!
/* package main  // An executable Go program must use the 'main' package. This package must contain a function named 'main', which serves as the program's entry point.

// A package is a collection of one or more Go source files that are compiled together. Multiple files can belong to the same package.
// Go considers a directory as a package, and insists every regular '.go' files in it to have the same package name.

// When an executable program runs, execution begins with the 'main' function in the 'main' package.

// Every non-'main' package is considered a 'library'.
// Since it has no starting point, 'package main' and 'func main() {...}', no library can run on its own,
// but is expected to be run in a main function.

import "fmt"  // 'fmt' is part of Go's standard library and provides functions for formatted input and output, such as 'Println' and 'Printf'.

func main() {
	fmt.Println("Hello, 世界!")
} */

// PRINT
/* package main

import "fmt"

func main() {
	a, b := "One", "Two"  // Declares and initializes the variables 'a' and 'b' using 'short variable declaration syntax'. Their types are inferred from the assigned values. This will be explained in detail later.

	print(a, b)  // The built-in 'print' function writes its arguments without inserting spaces between them or a new-line at the end.

	println(a, b)  // The 'println' function separates its arguments with spaces and adds a new-line at the end. Like 'print', 'println' is a built-in Go function, so it does not require importing 'fmt'. However, 'fmt.Println' is generally preferred for ordinary output.

	fmt.Println("This is what 'fmt.Println' should be used for:", a, b)  // 'Println' prints its arguments, separating them with spaces, and adds a new-line at the end.

	fmt.Printf("This is what 'fmt.Printf' should be used for: %v, %v", b, a)  // 'Printf' prints text according to a format string. It is somewhat similar to Python's f-strings, although it uses format verbs such as '%v' instead of embedded expressions. '%v' formats a value using its default representation. The supplied arguments are substituted into the format string in order.
} */

// FUNCTIONS
/* package main

import "fmt"

// A function can take zero or more arguments.

func add(x int, y int) int {  // Notice that the type comes after the variable names.
	return x + y
}

func multiply(x, y int) int {  // When two or more consecutive named function parameters share a type, you can omit the type from all but the last.
	return x * y
}

func main() {
	fmt.Println(add(42, 13))
	fmt.Println(multiply(42, 13))
} */

// RETURNS
// Single return:
/* package main

import "fmt"

func combine(x, y string) string {  // Single return; no parentheses
	return x + y
}

func main() {
	fmt.Println(combine("a", "b"))
} */

// No returns:
/* package main

import "fmt"

func swap(x, y int) {  // No return; no return definition
	fmt.Println(y, x)
}

func main() {
	swap(11, 9)
} */

// Multiple returns:
/* package main

import "fmt"

func swap(x, y int) (int, int) {  // Multiple returns; in parentheses
	return y, x
}

func main() {
	fmt.Println(swap(11, 9))
} */

// NAMED RETURN VALUES
// The ordinary way:
/* package main

import "fmt"

func split(sum int) (int, int) {
	x := sum * 4 / 9
	y := sum - x

	return x, y
}

func main() {
	fmt.Println(split(9))
} */

// The "Named Return Value" way:
/* package main

import "fmt"

func split(sum int) (x, y int) {  // 1. x and y are 'declared' here.
	x = sum * 4 / 9  // Notice '=' (assignment), not ':=' (declaration) (variable x already exists)
	y = sum - x

	return  // Naked Rturn (Returns x and y automatically)
}

func main() {
	fmt.Println(split(9))
} */

// VARIABLES
// Zero Values:
/* package main

import "fmt"

var c, python, java bool

func main() {
	var i int

	fmt.Println(c, python, java, i)  // false, false, false, 0 -> The 'zero values' of each variable type
} */

// With initializers:
/* package main

import "fmt"

var i, j int = 1, 2  // One initalizer per variable

func main() {
	var c, python, java = true, false, "no!"  // If an initializer is provided, the type can be omitted; the variable will take the type of the initializer

	fmt.Println(i, j, c, python, java)
} */

// The short assignment statement, ':=':
/* // Inside a function, the := short assignment statement can be used in place of a var declaration with implicit type.
// Outside a function, every statement begins with a keyword (var, func, and so on) and so the := construct is not available.

package main

import "fmt"

func main() {
	var i, j = 1, 2
	k := 3
	c, python, java := true, false, "no!"

	fmt.Println(i, j, k, c, python, java)
} */

// DATA TYPES
// PRIMITIVE
// Boolean (bool):
/* package main

import "fmt"

var isActive bool = true
var isCompleted bool  // Defaults to its zero value 'false'

func main() {
	fmt.Println(isActive, isCompleted)  // true, false
} */

// Numeric:
// Signed Integers:
// Hold both positive and negative whole numbers.

/* int8: 8-bit (-128 to 127)
int16: 16-bit (-32,768 to 32,767)
int32: 32-bit (-2,147,483,648 to 2,147,483,647)
int64: 64-bit (~ -9.22 * 10**18 to 9.22 * 10**18)
int: Platform-dependent (32 bits on 32-bit systems, 64 bits on 64-bit systems). This is the standard integer type used in most Go code. */

// Unsigned Integers:
// Hold only non-negative whole numbers (0 and positive).

/* uint8: 8-bit (0 to 255)
uint16: 16-bit (0 to 65,535)
uint32: 32-bit (0 to 4,294,967,295)
uint64: 64-bit (0 to ~1.84 * 10^19)  
uint: Platform-dependent (32 or 64 bits)  
uintptr: An unsigned integer large enough to store the uninterpreted bits of a pointer value. */

// Integer Aliases:
/* byte: Alias for uint8. Used to emphasize raw data or individual ASCII/UTF-8 bytes.
rune: Alias for int32. Represents a single Unicode code point (characters in single quotes, e.g., 'a', '世'). */

/* package main

import "fmt"

func main() {
	fmt.Printf("%v, %T", 'a', 'a')  // 97
} */

// Floating-Point Numbers:
// Represent numbers with fractional components (IEEE-754 standard).

/* float32: 32-bit floating-point (single precision).
float64: 64-bit floating-point (double precision). This is the default for floating-point literals. */

// Complex Numbers:
/* Represent numbers with real and imaginary parts.

complex64: Float32 real and imaginary parts.
complex128: Float64 real and imaginary parts. */

/* package main

import "fmt"

func main() {
	c := complex(2, 3)  // 2 + 3i

	x := 4 + 5i
	
	fmt.Println(real(c), imag(c))  // 2, 3
	fmt.Printf("%v, %T", x, x)  // (4+5i), complex128
} */

// Strings (string):
// A string is an immutable sequence of bytes, typically interpreted as UTF-8 encoded text.

/* Zero value: "" (empty string)
Interpreted string literal: Enclosed in double quotes "" (supports escape characters like \n, \t).
Raw string literal: Enclosed in backticks `...` (preserves line breaks and does not process escape sequences). */

/* package main

import "fmt"

func main() {
	s1 := "Hello, \nWorld!"
	s2:= `Line 1
	Line 2 without escaping \n`

	fmt.Println(s1)
	fmt.Println(s2)
} */
