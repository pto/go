// Demonstrate implementing a function in assembly
package main

import "fmt"

// add is implemented in assembly
func add(a int64, b int64) int64

func main() {
	fmt.Println(add(1234567890000000, 123456))
}
