// ct is an implementation of a process defined by an infinite set of
// recursive equations
package main

import (
	. "github.com/pto/go/csp/process"
)

// CT is a function of int returning a Process modeling something
// that moves around, up and down
var CT func(int) Process

func init() {
	// CT passes closures over itself to enable lazy evaluation
	CT = func(n int) Process {
		if n == 0 {
			return Choice2("around", func(e Event) Process { return CT(0)(e) },
				"up", func(e Event) Process { return CT(1)(e) })
		} else {
			return Choice2("up", func(e Event) Process { return CT(n + 1)(e) },
				"down", func(e Event) Process { return CT(n - 1)(e) })
		}
	}
}

func main() {
	Interact(Events{"around", "up", "down", "around", "√"}, CT(0))
	Interact(Events{"up", "up", "up", "up", "down", "down", "down", "down",
		"around", "around", "around", "√"}, CT(0))
	Interact(Events{"down", "around", "around", "up", "√"}, CT(1))
}
