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

// Interact runs the given Process with the given Event slice
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

func main() {
	// vend is a closure on itself, so that it can define itself recursively
	var vend Process
	vend = func(e Event) Process { return vend(e) }
	vend = Choice2(
		"in2p", Choice2(
			"large", vend,
			"small", Prefix("out1p", vend)),
		"in1p", Choice2(
			"small", vend,
			"in1p", Choice2("large", vend, "in1p", Stop)))

	Interact([]Event{"in2p", "large"}, vend)
	Interact([]Event{"in2p", "small", "out1p", "in1p", "in1p", "large",
		"in1p", "small"}, vend)
	Interact([]Event{"in1p", "in1p", "in1p", "END"}, vend)
}
