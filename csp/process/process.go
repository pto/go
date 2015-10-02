// process is an implementation of Hoare processes
package process

import (
	"strings"
)

// A Process is a function that accepts an Event and returns a Process
type Process func(Event) Process

// An Event is a string from the alphabet of a Process
type Event string

// A Trace is an array of Events
type Trace []Event

func (t Trace) String() string {
	// Convert to []String so we can use Join
	s := make([]string, 0, len(t))
	for _, e := range t {
		s = append(s, string(e))
	}
	return strings.Join(s, ", ")
}

// Stop is the Process that accepts no Events
func Stop(e Event) Process {
	return nil
}

// Prefix returns the Process that accepts Event c and then acts like Process p
func Prefix(c Event, p Process) Process {
	return func(e Event) Process {
		if e == c {
			return p
		}
		return nil
	}
}

// Choice2 returns the Process that is a choice between Events c and d
func Choice2(c Event, p Process, d Event, q Process) Process {
	return func(e Event) Process {
		if e == c {
			return p
		} else if e == d {
			return q
		} else {
			return nil
		}
	}
}

// IsTrace returns true if Trace t is a possible trace of Process p
func IsTrace(t Trace, p Process) bool {
	if t == nil || len(t) == 0 {
		return true
	} else if p(t[0]) == nil {
		return false
	} else {
		return IsTrace(t[1:], p(t[0]))
	}
}
