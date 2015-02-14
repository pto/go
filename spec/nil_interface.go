package main

import (
	"fmt"
)

func main() {
	var x interface{}
	fmt.Println("nil interface equals nil:", x == nil)
	var y []int
	fmt.Println("nil slice equals nil:", y == nil)
	x = y
	fmt.Println("interface set to nil slice equals nil:", x == nil)
	x = nil
	fmt.Println("interface set to nil equals nil:", x == nil)
}
