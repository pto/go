package main

import "fmt"

func main() {
	var s = make([][]int, 2)
	s[0] = make([]int, 2)
	s[1] = make([]int, 3)
	s[0][0] = 1
	s[0][1] = 2
	s[1][2] = 3
	fmt.Println(s)
}
