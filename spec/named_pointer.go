package main

import "fmt"

type Thing struct {
	val int
}

func (t Thing) PrintVal() {
	fmt.Println("val is", t.val)
}

type PT *Thing

// Show that the selector expression rule for named pointer types is
// boring, since unnamed pointer types are also automatically dereferenced.
func main() {
	var t Thing = Thing{val: 42}
	var named PT = &t
	var unnamed *Thing = &t
	fmt.Println(named.val)   // Special rule
	fmt.Println(unnamed.val) // Normal deferencing
	// ERROR: named.PrintVal() // Type PT has no method PrintVal
	(*named).PrintVal() // Special rule doesn't apply
	unnamed.PrintVal()
}
