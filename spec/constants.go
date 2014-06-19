package main

import (
	"fmt"
)

func main() {
	// Unusual, but valid literal
	one := 00000000001
	fmt.Println("00000000001 equals", one)

	var (
		ui8  uint8      = 3.0
		i8   int8       = 3.0
		ui16 uint16     = 3.0
		i16  int16      = 3.0
		ui32 uint32     = 3.0
		i32  int32      = 3.0
		ui64 uint64     = 3.0
		i64  int64      = 3.0
		f32  float32    = 3.0
		f64  float64    = 3.0
		c64  complex64  = 3.0
		c128 complex128 = 3.0
	)

	fmt.Println(ui8, i8, ui16, i16, ui32, i32, ui64, i64, f32, f64, c64, c128)

	ui32 = 2147483648.0
	f32 = 2147483648.0
	f64 = 2147483648.0
	// i32 = 2147483648.0 // error

	fmt.Println(1<<31, ui32, f32, f64 /* i32 */)

	fmt.Println("Largest integer constant is bigger than", float64(1<<255))
	const (
		startInt = 1 << 255
		nextInt  = startInt + 1
		diffInt  = startInt - nextInt
	)
	fmt.Println("The difference between two huge integer consts is", diffInt)
	const (
		startFloat = 1.0e1234
		nextFloat  = 1.1e1234
		diffFloat  = startFloat - nextFloat
	)
	fmt.Println("The difference between two huge floating consts is", diffFloat)
}
