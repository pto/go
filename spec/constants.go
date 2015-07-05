package main

import (
	"fmt"
)

// Show characteristics of constants, plus a compiler limitation.
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
		f64             = 3.0 // this is the default type for 3.0
		c64  complex64  = 3.0
		c128 complex128 = 3.0
	)

	fmt.Println(ui8, i8, ui16, i16, ui32, i32, ui64, i64, f32, f64, c64, c128)

	ui32 = 2147483648.0
	f32 = 2147483648.0
	f64 = 2147483648.0
	// i32 = 2147483648.0 // error

	fmt.Println("1<<31 is", 1<<31, ui32, f32, f64 /* , i32 */)

	fmt.Println("Largest integer constant is bigger than", float64(1<<499))

	const (
		startInt = 1 << 499
		nextInt  = startInt + 1
		diffInt  = nextInt - startInt
	)
	fmt.Println("The difference between two huge integer consts is", diffInt)

	const (
		oneFloat   = 1e9994
		twoFloat   = 6.2793694463797e9994 // old max floating constant
		ratioFloat = twoFloat / oneFloat
	)
	fmt.Println("The ratio between two huge floating consts is", ratioFloat)

	const (
		bigFloat     = 1e9999999
		biggerFloat  = 1e10000000
		anotherRatio = biggerFloat / bigFloat
	)
	fmt.Println("The ratio between two truly enormous consts is", anotherRatio)

	const (
		smallFloat   = 0.111111111122222222223333333333444444444455555555556666666666777777777788888888889999999999123456789
		smallerFloat = 0.111111111122222222223333333333444444444455555555556666666666777777777788888888889999999999
		diffFloat    = smallFloat - smallerFloat
	)
	fmt.Println("The difference between two small floats is", diffFloat)
}
