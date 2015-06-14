// process is an implementation of Hoare processes
package main

import (
	"fmt"
	"os"
)

// A Process is a function that takes an Event and returns a Process
type Process func(Event) Process

// An Event is a string from the alphabet of a Process
type Event string

// Stop is the Process that accepts no Events
func Stop(e Event) Process {
	if e == "END" {
		fmt.Println("Stopped")
	} else {
		fmt.Println("Tried to send invalid event", e)
		os.Exit(1)
	}
	return nil
}

// Prefix returns the Process the accepts Event c and then acts like Process p
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

// Choice2 returns the Process that is a choice between Events
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

// Interact runs Process p with the Event slice events
func Interact(events []Event, p Process) {
	fmt.Println("Executing", events)
	current := p
	for _, e := range events {
		current = current(e)
	}
	if current != nil {
		fmt.Println("Did not reach Stop")
	}
}

// VMC is a complicated vending machine
var VMC Process

func init() {
	// VMC is a closure of itself, so it can be defined recursively
	VMC = func(e Event) Process { return VMC(e) }
	VMC = Choice2(
		"in2p", Choice2(
			"large", VMC,
			"small", Prefix("out1p", VMC)),
		"in1p", Choice2(
			"small", VMC,
			"in1p", Choice2("large", VMC, "in1p", Stop)))
}

func main() {
	Interact([]Event{"in2p", "large"}, VMC)
	Interact([]Event{"in2p", "small", "out1p", "in1p", "in1p", "large",
		"in1p", "small"}, VMC)
	Interact([]Event{"in1p", "in1p", "in1p", "END"}, VMC)
}
