package main

import (
	"fmt"
	"unsafe"
)

func main() {
	fmt.Println("unsafe.Sizeof(true) is", unsafe.Sizeof(true))
	fmt.Println("unsafe.Sizeof(uint8(0)) is", unsafe.Sizeof(uint8(0)))
	fmt.Println("unsafe.Sizeof(uint16(0)) is", unsafe.Sizeof(uint16(0)))
	fmt.Println("unsafe.Sizeof(uint64(0)) is", unsafe.Sizeof(uint64(0)))
	fmt.Println("unsafe.Sizeof(uint(0)) is", unsafe.Sizeof(uint(0)))
	fmt.Println("unsafe.Sizeof(int(0)) is", unsafe.Sizeof(int(0)))
	var x *int
	fmt.Println("unsafe.Sizeof(*int) is", unsafe.Sizeof(x))
	fmt.Println("unsafe.Sizeof(uintptr(0)) is", unsafe.Sizeof(uintptr(0)))
	var y interface{}
	fmt.Println("unsafe.Sizeof(interface{}) is", unsafe.Sizeof(y))
}
