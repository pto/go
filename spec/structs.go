package main

import "fmt"

// T is a simple type for embedding
type T int64

func (t T) Do() {
	fmt.Println("Do:", t)
}

func (t *T) PtrDo() {
	fmt.Println("PtrDo:", *t)
}

// S embeds T
type S struct {
	T
}

// SP embeds *T
type SP struct {
	*T
}

type Doer interface {
	Do()
}

type PtrDoer interface {
	PtrDo()
}

// Shows that method sets of non-pointer objects only include non-pointer
// methods, but method sets of pointer objects include both pointer and
// non-pointer methods.
func main() {
	var s1, s2 S
	s1.T = 111
	s2.T = 222 // automatic dereference to use field
	var ps2 = &s2
	fmt.Println("s1.T:", s1.T)   // normal use of embedded field
	fmt.Println("ps2.T:", ps2.T) // automatic dereference to use field
	fmt.Println("calls on s1 and ps1:")
	s1.Do()     // normal call of non-pointer method
	s1.PtrDo()  // automatic address of (even though not in method set of S)
	ps2.Do()    // automatic dereference
	ps2.PtrDo() // normal call of pointer method
	fmt.Println()

	var i Doer
	var pi PtrDoer
	fmt.Println("calls on interface variables:")
	i = s1 // ok: method set of S includes promoted method with receiver T
	i.Do()
	i = ps2 // ok: method set of *S includes promoted method with receiver T
	i.Do()
	// pi = s1 // not ok: method set of S does not include promoted method with receiver *T
	// pi.PtrDo()
	pi = ps2 // ok: method set of *S includes promoted method with receiver *T
	pi.PtrDo()
	fmt.Println()

	var sp1, sp2 SP
	sp1.T = new(T)
	*(sp1.T) = 333
	sp2.T = new(T)
	*(sp2.T) = 444
	var psp2 = &sp2
	fmt.Println("*sp1.T:", *sp1.T)   // normal dereferece of embedded pointer
	fmt.Println("*psp2.T:", *psp2.T) // normal deref of auto deref
	fmt.Println("calls on s2 and ps2")
	sp1.Do()     // automatic dereference of *T
	sp1.PtrDo()  // normal use of *T
	psp2.Do()    // double dereference of psp2 and *T
	psp2.PtrDo() // automatic deference of psp2
	fmt.Println()

	// All of these are OK because the method set of embedded field *T
	// is available to both S and *S
	fmt.Println("calls on interface variables:")
	i = sp1
	i.Do()
	i = psp2
	i.Do()
	pi = sp1
	pi.PtrDo()
	pi = psp2
	pi.PtrDo()
}
