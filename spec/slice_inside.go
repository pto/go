package main

import (
	"fmt"
	"unsafe"
)

func main() {
	s := make([]int, 8)
	fmt.Println("unsafe.Sizeof(s) is", unsafe.Sizeof(s))
	p := unsafe.Pointer(&s)
	b := (*[unsafe.Sizeof(s)]byte)(p)
	fmt.Printf("contents of s is % #x\n", b)
}
