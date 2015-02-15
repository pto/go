package main

import "fmt"

// Show that accented code point works as a rune, but a combining
// code point plus another code point does not.
func main() {
	fmt.Println("This is \"\\u00E9\":", "é")
	fmt.Println("This is '\\u00E9':", 'é')
	fmt.Println("This is \"e\\u0301\":", "é")
	fmt.Println("This is invalid: 'e\\u0301'")
	// illegal rune literal: fmt.Println('é')
}
