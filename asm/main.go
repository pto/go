package main

import "fmt"

func add(a int64, b int64) int64

func main() {
	fmt.Println(add(0xFFFFFFFF, 1))
}
