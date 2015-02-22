package main

import (
	"fmt"
)

type T int

func (t T) FuncOfT() {
	fmt.Println("FuncOfT()")
}

func (t *T) FuncOfPT() {
	fmt.Println("FuncOfPT()")
}

type PT *T

type IT interface {
	FuncOfT()
}

type IPT interface {
	FuncOfPT()
}

// Show that the method set of *T includes the method set of T as well
// as the method set of *T. But the method set of T does NOT include
// the method set of *T.
// Further, the method set of new type PT defined as *T is empty.
func main() {
	t1 := T(42)
	t2 := new(T)
	*t2 = 123
	fmt.Printf("t1: %[1]v %[1]T\n", t1, t1)
	fmt.Printf("t2: %[1]v %[1]T\n", t2, t2)
	fmt.Printf("*t2: %[1]v %[1]T\n", *t2, *t2)

	var it IT
	var ipt IPT
	it = t1 // Obviously OK
	fmt.Printf("it: %[1]v %[1]T\n", it, it)
	it = t2 // Method set of *T contains method set of T
	fmt.Printf("it: %[1]v %[1]T\n", it, it)
	fmt.Printf("*(it.(*T)): %[1]v %[1]T\n", *(it.(*T)), *(it.(*T)))
	// ipt = t1 // WRONG: Method set of T does not contain method set of *T
	ipt = t2 // Obviously OK
	fmt.Printf("ipt: %[1]v %[1]T\n", ipt, ipt)
	fmt.Printf("*(ipt.(*T)): %[1]v %[1]T\n", *(ipt.(*T)), *(ipt.(*T)))

	var pt PT = &t1
	(*pt).FuncOfT() // Won't automatically dereference
	// it = pt  // WRONG: Method set of PT is empty
	// ipt = pt // WRONG: Method set of PT is empty
}
