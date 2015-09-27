// vmc is an implementation of a complex vending machine
package main

import (
	"fmt"
	. "github.com/pto/go/csp/process"
)

// VMC is a Process modeling a complicated vending machine
var VMC Process

func init() {
	// VMC is a closure of itself, so it can be defined recursively
	VMC = func(e Event) Process {
		fmt.Println("calling VMC closure for event", e)
		return VMC(e)
	}
	// At this point the variable VMC holds a value that is a function
	// containing a closure over the variable VMC. When this
	// function runs, it will evaluate the current value of the VMC
	// variable and call it.
	//
	// There is, therefore, a difference between the value of VMC at this
	// point in the code execution (the closure of VMC over itself), and
	// the value of VMC as it will be when the function value from the
	// previous line is actually executed (the more complicated VMC that
	// will be defined next).
	//
	// When the next line of code runs, it will evaluate the current
	// value of VMC (the closure of VMC over itself) and pass that value
	// to Choice2 and Prefix. That value, like all non-pointer function
	// arguments, will be passed by value, not by reference. So the value
	// that Choice2 and Prefix see is the original value of VMC (the closure
	// of VMC over itself). That value is what Choice2 and Prefix capture
	// in the closures that they create and return.
	//
	// After the next line of code runs, VMC will have a new value. That value
	// is the more complex function that defines the complicated vending
	// machine process. It is this new value of VMC that will be passed
	// to the Interact function in main.
	//
	// The trick here is that this new value of VMC is the result of
	// executing a closure over the old value of VMC, which itself is a
	// closure over the variable VMC. So when the new value of VMC
	// executes, it will ultimately reference not the old VMC, but the new
	// VMC. And so we get the desired recursive definition.
	VMC = Choice2(
		"in2p", Choice2(
			"large", VMC,
			"small", Prefix("out1p", VMC)),
		"in1p", Choice2(
			"small", VMC,
			"in1p", Choice2("large", VMC, "in1p", Stop)))
}

func main() {
	Interact(Events{"in2p", "large", "√"}, VMC)
	Interact(Events{"in2p", "small", "out1p", "in1p", "in1p", "large",
		"in1p", "small", "√"}, VMC)
	Interact(Events{"in1p", "in1p", "in1p", "√"}, VMC)
	Interact(Events{"in2p", "large", "in2p", "large", "in2p", "large", "√"},
		VMC)
	// Three 1p coins in a row lead to STOP
	Interact(Events{"in1p", "in1p", "in1p", "in1p", "√"}, VMC)
}
