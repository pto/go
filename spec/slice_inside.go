package main

import (
	"fmt"
	"unsafe"
)

func main() {
	s := make([]int, 8)
	fmt.Println("unsafe.Sizeof(s) is", unsafe.Sizeof(s))
}
