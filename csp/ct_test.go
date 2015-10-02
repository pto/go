// ct is an implementation of a process defined by an infinite set of
// recursive equations
package csp

import (
	. "github.com/pto/go/csp/process"
	"testing"
)

func TestCtSuccess(t *testing.T) {
	if !IsTrace(Trace{"around", "up", "down", "around"}, CT(0)) {
		t.Error("CT should go around, up, down and arround")
	}
	if !IsTrace(Trace{"up", "up", "up", "up", "down", "down", "down", "down", "around", "around", "around"}, CT(0)) {
		t.Error("CT should go way up, back down and around")
	}
	if !IsTrace(Trace{"down", "around", "around", "up"}, CT(1)) {
		t.Error("CT should go down from level 1")
	}
	if !IsTrace(Trace{}, CT(123)) {
		t.Error("CT(123) should satisfy an empty Trace")
	}
}

func TestCtFailure(t *testing.T) {
	if IsTrace(Trace{"down"}, CT(0)) {
		t.Error("CT(0) should not go down")
	}
	if IsTrace(Trace{"around"}, CT(1)) {
		t.Error("CT(1) should not go around")
	}
}
