package main

import (
	"fmt"
)

type MyBool bool

// Shows that bool constants are typeless.
func main() {
	var a bool
	fmt.Println("a is", a)
	var b MyBool
	fmt.Println("b is", b)
	a = false
	b = true // OK, since true is a typeless constant
	// a = b // WRONG, because type of b is not equal to type of a
	a = bool(b) // OK, conversion
	fmt.Println("now a is", a)
}
