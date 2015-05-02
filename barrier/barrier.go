package barrier

import (
	"runtime"
	"sync"
)

// Dataset simulates a data set that must be run through a series of steps.
type Dataset struct {
	data            []int
	precheckErrors  int
	postcheckErrors int
	count           int
	countMutex      sync.Mutex
	entry           chan int
	exit            chan int
	waitgroup       sync.WaitGroup
	maxGoroutines   int
}

// NewDataset initializes a data set.
func NewDataset(size int) *Dataset {
	d := new(Dataset)
	d.data = make([]int, size)
	d.entry = make(chan int, size)
	d.exit = make(chan int, size)
	d.waitgroup.Add(size)
	return d
}

// Process simulates a processor.
func (d *Dataset) Process(index int, steps int) {
	for i := 0; i < steps; i++ {
		d.entryBarrier()
		d.precheck(index)
		d.calculate(index)
		d.postcheck(index)
		d.exitBarrier()
	}
	d.waitgroup.Done()
}

// calculate simulates a calculation.
func (d *Dataset) calculate(i int) {
	d.data[i]++
	g := runtime.NumGoroutine()
	if g > d.maxGoroutines {
		d.maxGoroutines = g
	}
}

// entryBarrier blocks until all processes are ready to start.
func (d *Dataset) entryBarrier() {
	d.countMutex.Lock()
	d.count++
	if d.count == len(d.data) {
		for i := 0; i < len(d.data); i++ {
			d.entry <- 1
		}
	}
	d.countMutex.Unlock()
	<-d.entry
}

// exitBarrier blocks until all processes are finished.
func (d *Dataset) exitBarrier() {
	d.countMutex.Lock()
	d.count--
	if d.count == 0 {
		for i := 0; i < len(d.data); i++ {
			d.exit <- 1
		}
	}
	d.countMutex.Unlock()
	<-d.exit
}

// precheck counts processes that are too far ahead or behind before this
// process does its calculation.
func (d *Dataset) precheck(i int) {
	for _, v := range d.data {
		if v < d.data[i] || v > d.data[i]+1 {
			d.precheckErrors++
		}
	}
}

// postcheck counts processes that are too far ahead of behind after this
// process does its calculation.
func (d *Dataset) postcheck(i int) {
	for _, v := range d.data {
		if v < d.data[i]-1 || v > d.data[i] {
			d.postcheckErrors++
		}
	}
}

// Wait blocks until all processes are complete.
func (d *Dataset) Wait() {
	d.waitgroup.Wait()
}

// VerifyErrors counts any results that don't have the expected value.
func (d *Dataset) VerifyErrors(expected int) int {
	errors := 0
	for _, v := range d.data {
		if v != expected {
			errors++
		}
	}
	return errors
}

// CheckErrors returns a count of pre- and post-check errors.
func (d *Dataset) CheckErrors() (int, int) {
	return d.precheckErrors, d.postcheckErrors
}

// MaxGoroutines returns the maximum number of goroutines.
func (d *Dataset) MaxGoroutines() int {
	return d.maxGoroutines
}
