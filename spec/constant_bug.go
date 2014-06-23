package main

import "fmt"

func main() {
	const (
		a = 1e32767
		b = 2e32767
		c = a / b
	)
	fmt.Println(c)
}
