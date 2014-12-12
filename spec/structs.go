package main

import "fmt"

// T is a simple type for embedding
type T int64

func (t T) doIt() {
	fmt.Println("doIt:", t)
}

func (t *T) ptrDoIt() {
	fmt.Println("ptrDoIt:", *t)
}

// S embeds T
type S struct {
	T
}

// S2 embeds *T
type S2 struct {
	*T
}

// I is a doIter
type I interface {
	doIt()
}

// PI is a ptrDoIter
type PI interface {
	ptrDoIt()
}

func main() {
	var s S
	var ps = &s
	s.T = 123
	fmt.Println("values of T:", s.T, ps.T)
	fmt.Println("calls on s and ps:")
	s.doIt()
	s.ptrDoIt() // automatic address of (even though not in method set of S)
	ps.doIt()   // automatic dereference
	ps.ptrDoIt()

	var x I
	var px PI
	x = s  // ok: method set of S includes promoted method with receiver T
	x = ps // ok: method set of *S includes promoted method with receiver T
	// px = s  // not ok: method set of S does not include promoted method with receiver *T
	px = ps // ok: method set of *S includes promoted method with receiver *T
	fmt.Println("calls on interface variables:")
	x.doIt()
	px.ptrDoIt()

	var s2 S2
	var ps2 = &s2
	s2.T = new(T)
	*(s2.T) = 456
	fmt.Println("values of *T:", *s2.T, *ps2.T)
	fmt.Println("calls on s2 and ps2")
	s2.doIt() // automatic dereference of *T
	s2.ptrDoIt()
	ps2.doIt()    // double dereference of ps2 and *T
	ps2.ptrDoIt() // automatic deference of ps2

	x = s2
	x = ps2
	px = s2
	px = ps2
	fmt.Println("calls on interface variables:")
	x.doIt()
	px.ptrDoIt()
}
