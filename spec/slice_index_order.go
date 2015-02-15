package main

import "fmt"

// Show that slice indexes use the customary order, and behave as expected
// for a slice of slices.
func main() {
	x := make([][]int, 1)
	x[0] = make([]int, 2)
	x[0][0] = 1
	x[0][1] = 2
	fmt.Println(x)

	// A different way to initialize
	y := make([][]int, 0)
	y = append(y, make([]int, 1))
	y = append(y, make([]int, 1))
	y[0][0] = 10
	y[1][0] = 11
	fmt.Println(y)

	// Or even this
	y[0] = append(y[0], 42)
	fmt.Println(y)
	fmt.Println(y[0][1])
}
