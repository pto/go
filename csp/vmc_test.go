package csp

import (
	. "github.com/pto/go/csp/process"
	"testing"
)

func TestVMCSuccess(t *testing.T) {
	if !IsTrace(Trace{"in2p", "large"}, VMC) {
		t.Error("VMC should vend a large item")
	}
	if !IsTrace(Trace{"in2p", "small", "out1p", "in1p", "in1p", "large", "in1p", "small"}, VMC) {
		t.Error("VMC should vend a small with change, a large and a small")
	}
	if !IsTrace(Trace{"in1p", "in1p", "in1p"}, VMC) {
		t.Error("VMC should accept 3 pennies")
	}
	if !IsTrace(Trace{"in2p", "large", "in2p", "large", "in2p", "large"}, VMC) {
		t.Error("VMC should vend a 3 large items")
	}
	if !IsTrace(Trace{}, VMC) {
		t.Error("An empty trace should satisfy any process")
	}
}

func TestVMCFailure(t *testing.T) {
	if IsTrace(Trace{"in1p", "in1p", "in1p", "in1p"}, VMC) {
		t.Error("VMC should STOP after 3 large items")
	}
	if IsTrace(Trace{"large"}, VMC) {
		t.Error("VMC should not vend without a coin")
	}
}
