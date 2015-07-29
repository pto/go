// process is an implementation of Hoare processes
package main

import (
	"fmt"
	"os"
	"strings"
)

// A Process is a function that accepts an Event and returns a Process
type Process func(Event) Process

// An Event is a string from the alphabet of a Process
type Event string
type Events []Event

func (events Events) String() string {
	// Convert []Event to []String so we can use Join
	s := make([]string, 0, len(events))
	for _, e := range events {
		s = append(s, string(e))
	}
	return strings.Join(s, ", ")
}

// Stop is the Process that accepts no Events
func Stop(e Event) Process {
	fmt.Println("Tried to send invalid event", e, "to Stop")
	os.Exit(1)
	return nil
}

// Prefix returns the Process that accepts Event c and then acts like Process p
func Prefix(c Event, p Process) Process {
	return func(e Event) Process {
		if e == c {
			fmt.Println("Prefix accepted", e)
			return p
		}
		fmt.Println("Tried to send invalid event", e, "to Prefix")
		os.Exit(1)
		return nil
	}
}

// Choice2 returns the Process that is a choice between Events c and d
func Choice2(c Event, p Process, d Event, q Process) Process {
	return func(e Event) Process {
		if e == c {
			fmt.Println("Choice2 accepted", e)
			return p
		} else if e == d {
			fmt.Println("Choice2 accepted", e)
			return q
		} else {
			fmt.Println("Tried to send invalid event", e, "to Choice2")
			os.Exit(1)
			return nil
		}
	}
}

// Interact runs Process p with Event slice events
func Interact(events Events, p Process) {
	fmt.Println("-- Executing ‹", events, "›")
	current := p
	for _, e := range events {
		if e == "√" {
			fmt.Println("Successful termination")
			return
		}
		current = current(e)
	}
}

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
	// When the next line of code runs, it will evaluated the current
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
	//
	// We could have avoided all this by passing a pointer to VMC instead
	// of the value VMC. But where's the fun in that? And this version
	// of VMC is closer in spirit to the original LISP program by Hoare.
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
