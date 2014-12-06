package main

import (
	"fmt"
	"unsafe"
)

func main() {
	s := make([]int, 1, 2)
	fmt.Println("ACTION: make([]int, 1, 2)")
	var i int
	fmt.Println("unsafe.Sizeof(s) is", unsafe.Sizeof(s),
		"which is 3 times", unsafe.Sizeof(i), "(unsafe.Sizeof(int))")

	ps := unsafe.Pointer(&s)
	bs := (*[unsafe.Sizeof(s)]byte)(ps)

	fmt.Println("s is", s, "with len of", len(s), "and cap of", cap(s))
	fmt.Printf("contents of s is % #x\n", *bs)

	s = append(s, 123)
	fmt.Println("ACTION: append(s, 123)")
	fmt.Println("s is", s, "with len of", len(s), "and cap of", cap(s))
	fmt.Printf("contents of s is % #x\n", *bs)

	s = append(s, 456)
	fmt.Println("ACTION: append(s, 456)")
	fmt.Println("s is", s, "with len of", len(s), "and cap of", cap(s))
	fmt.Printf("contents of s is % #x\n", *bs)

	var a [3]int
	sa := a[:]
	fmt.Println("ACTION: sa = a[:]")
	psa := unsafe.Pointer(&sa)
	bsa := (*[unsafe.Sizeof(sa)]byte)(psa)
	fmt.Println("sa is", sa, "with len of", len(sa), "and cap of", cap(sa))
	fmt.Printf("contents of sa is % #x\n", *bsa)
	fmt.Printf("address of a is %p\n", &a)
}
