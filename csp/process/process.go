// process is an implementation of Hoare processes
package process

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
