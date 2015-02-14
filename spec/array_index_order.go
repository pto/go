package main

import "fmt"

func main() {
	var x [1][2]int
	x[0][0] = 1
	x[0][1] = 2
	fmt.Println(x)

	var y [2][1]int
	y[0][0] = 10
	y[1][0] = 11
	fmt.Println(y)
}
