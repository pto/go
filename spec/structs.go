package main

import "fmt"

type T int64

func (t T) doIt() {
	fmt.Println("doIt:", t)
}

func (t *T) ptrDoIt() {
	fmt.Println("ptrDoIt:", *t)
}

type S struct {
	T
}

type S2 struct {
	*T
}

type I interface {
	doIt()
}

type PI interface {
	ptrDoIt()
}

func main() {
	var s S
	var ps *S = &s
	s.T = 123
	fmt.Println("values of T:", s.T, ps.T)
	fmt.Println("calls on s and ps:")
	s.doIt()
	s.ptrDoIt() // automatic address of
	ps.doIt()   // automatic dereference
	ps.ptrDoIt()

	var x I
	var px PI
	x = s  // ok: method set of S includes receiver T
	x = ps // ok: method set of *S includes receiver T
	// px = s  // not ok: method set of S does not include receiver *T
	px = ps // ok: method set of *S includes receiver *T
	fmt.Println("calls on interface variables:")
	x.doIt()
	px.ptrDoIt()

	var s2 S2
	var ps2 *S2 = &s2
	s2.T = new(T)
	*(s2.T) = 456
	fmt.Println("values of *T:", *s2.T, *ps2.T)
	fmt.Println("calls on s2 and ps2")
	s2.doIt() // automatic dereference
	s2.ptrDoIt()
	ps2.doIt()    // double dereference?
	ps2.ptrDoIt() // automatic deference

	x = s2
	x = ps2
	px = s2
	px = ps2
	fmt.Println("calls on interface variables:")
	x.doIt()
	px.ptrDoIt()
}
