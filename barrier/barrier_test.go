package barrier

import (
	"fmt"
	"runtime"
	"testing"
)

func TestBarrierSmall(t *testing.T) {
	testBarrier(t, 100, 5)
}

func TestBarrierMedium(t *testing.T) {
	testBarrier(t, 1000, 100)
}

func TestBarriersLarge(t *testing.T) {
	testBarrier(t, 10000, 10)
}

func testBarrier(t *testing.T, size int, steps int) {
	runtime.GOMAXPROCS(runtime.NumCPU())
	d := NewDataset(size)
	for i := 0; i < size; i++ {
		go d.Process(i, steps)
	}
	d.Wait()
	if errs := d.VerifyErrors(steps); errs > 0 {
		t.Errorf("%d verify errors", errs)
	}
	if pre, post := d.CheckErrors(); pre > 0 || post > 0 {
		t.Errorf("%d precheck and %d postcheck errors", pre, post)
	}
	fmt.Println("maxGoroutines", d.MaxGoroutines())
}
