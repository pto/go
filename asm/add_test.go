package main

import (
	"math"
	"testing"
)

func TestAdd(t *testing.T) {
	if add(2, 2) != 4 {
		t.Error("2 + 2 is not 4")
	}
	if add(-1, 1) != 0 {
		t.Error("-1 + 1 is not 0")
	}
	if add(1234567890000000, 123456) != 1234567890123456 {
		t.Error("1234567890000000 + 123456 is not 1234567890123456")
	}
	if add(math.MaxInt64, 1) != math.MinInt64 {
		t.Error("math.MaxInt64 + 1 is not math.MinInt64")
	}
	if add(math.MinInt64, math.MinInt64) != 0 {
		t.Error("math.MinInt64 + math.MinInt64 is not 0")
	}
	if add(0, math.MinInt64) != math.MinInt64 {
		t.Error("0 + math.MinInt64 is not math.MinInt64")
	}
	if add(1, math.MaxInt64-1) != math.MaxInt64 {
		t.Error("1 + (math.MaxInt64 - 1) is not math.MaxInt64")
	}
}
