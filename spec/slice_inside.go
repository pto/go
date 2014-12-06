package main

import (
	"fmt"
	"unsafe"
)

func main() {
	s := make([]int, 1, 2)
	fmt.Println("len(s) is", len(s), "and cap(s) is", cap(s))
	fmt.Println("unsafe.Sizeof(s) is", unsafe.Sizeof(s))
	p := unsafe.Pointer(&s)
	b := (*[unsafe.Sizeof(s)]byte)(p)
	fmt.Printf("contents of s is % #x\n", b)
}
