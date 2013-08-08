package judy

import (
	"math"
	"math/rand"
	"testing"
)

func TestEmptyJudyArray(t *testing.T) {

	j := Judy1{}
	r := j.Free()

	if r != 0 {
		t.Errorf("Free should return 0, returned %v", r)
	}
}

func TestCount(t *testing.T) {

	j := Judy1{}
	defer j.Free()

	if ct := j.CountAll(); ct != 0 {
		t.Errorf("Count should be zero, was %v", ct)
	}

	var i uint64
	for i = 0; i < 100; i++ {
		j.Set(i)
	}

	var ct uint64
	if ct = j.CountAll(); ct != 100 {
		t.Errorf("Count should be 100, was %v", ct)
	}

	if ct = j.CountFrom(0, 1000); ct != 100 {
		t.Errorf("Count should be 100, was %v", ct)
	}
	if ct = j.CountFrom(200, 1000); ct != 0 {
		t.Errorf("Count should be 0, was %v", ct)
	}
	if ct = j.CountFrom(5, 5); ct != 1 {
		t.Errorf("Count should be 1, was %v", ct)
	}
	if ct = j.CountFrom(20, 29); ct != 10 {
		t.Errorf("Count should be 10, was %v", ct)
	}

}

func TestSet(t *testing.T) {

	j := Judy1{}
	defer j.Free()

	var i uint64
	for i = 0; i < 100; i++ {
		j.Set(i * 100000)
	}

	for i = 0; i < 100; i++ {
		if !j.Test(i * 100000) {
			t.Errorf("Index %v not set", i*100000)
		}
	}

	for i = 1; i < 100; i++ {
		if j.Test(i * 99999) {
			t.Errorf("Index %v incorrectly set", i*99999)
		}
	}

}

func TestUnset(t *testing.T) {

	j := Judy1{}
	defer j.Free()

	var i uint64
	for i = 0; i < 100; i++ {
		j.Set(i * 100000)
	}

	for i = 0; i < 100; i++ {
		j.Unset(i * 100000)
	}

	for i = 0; i < 100; i++ {
		if j.Test(i * 100000) {
			t.Errorf("Index %v incorrectly set", i*100000)
		}
	}

	if ct := j.CountAll(); ct != 0 {
		t.Errorf("Count should be zero, was %v", ct)
	}
}

func runOrderedMemUsageTest(t *testing.T, n int) {
	j := Judy1{}
	defer j.Free()

	for i := 0; i < n; i++ {
		j.Set(uint64(i * 10000))
	}

	if ct := j.CountAll(); int(ct) != n {
		t.Errorf("Count should be %v, was %v", n, ct)
	}
	t.Logf("Memory Usage with %7v ordered bits %8v", n, j.MemoryUsed())
}

func runRandomMemUsageTest(t *testing.T, n int) {
	j := Judy1{}
	defer j.Free()

	for i := 0; i < n; i++ {
		j.Set(uint64(rand.Int63()))
	}

	if ct := j.CountAll(); int(ct) != n {
		t.Errorf("Count should be %v, was %v", n, ct)
	}
	t.Logf("Memory Usage with %7v random bits  %8v", n, j.MemoryUsed())
}

func TestMemUsage(t *testing.T) {

	runOrderedMemUsageTest(t, 1000)
	runRandomMemUsageTest(t, 1000)
	runOrderedMemUsageTest(t, 10000)
	runRandomMemUsageTest(t, 10000)
	runOrderedMemUsageTest(t, 100000)
	runRandomMemUsageTest(t, 100000)
	runOrderedMemUsageTest(t, 1000000)
	runRandomMemUsageTest(t, 1000000)

	//t.Fail() // Uncomment to see the log output
}

func BenchmarkCountAllRand1000(b *testing.B) {
	j := Judy1{}
	defer j.Free()

	n := 1000
	for i := 0; i < n; i++ {
		j.Set(uint64(rand.Int63()))
	}

	for loops := 0; loops < b.N; loops++ {
		if ct := j.CountAll(); int(ct) != n {
			b.Errorf("Count should be %v, was %v", n, ct)
		}
	}
}

func BenchmarkCountAllRand1000000(b *testing.B) {
	j := Judy1{}
	defer j.Free()

	n := 1000000
	for i := 0; i < n; i++ {
		j.Set(uint64(rand.Int63()))
	}

	for loops := 0; loops < b.N; loops++ {
		if ct := j.CountAll(); int(ct) != n {
			b.Errorf("Count should be %v, was %v", n, ct)
		}
	}
}

func BenchmarkCountAllOrd1000(b *testing.B) {
	j := Judy1{}
	defer j.Free()

	n := 1000
	for i := 0; i < n; i++ {
		j.Set(uint64(i))
	}

	for loops := 0; loops < b.N; loops++ {
		if ct := j.CountAll(); int(ct) != n {
			b.Errorf("Count should be %v, was %v", n, ct)
		}
	}
}

func BenchmarkCountAllOrd1000000(b *testing.B) {
	j := Judy1{}
	defer j.Free()

	n := 1000000
	for i := 0; i < n; i++ {
		j.Set(uint64(i))
	}

	for loops := 0; loops < b.N; loops++ {
		if ct := j.CountAll(); int(ct) != n {
			b.Errorf("Count should be %v, was %v", n, ct)
		}
	}
}

func BenchmarkCountRangeRand1000(b *testing.B) {
	j := Judy1{}
	defer j.Free()

	n := 1000
	for i := 0; i < n; i++ {
		j.Set(uint64(rand.Int63()))
	}

	for loops := 0; loops < b.N; loops++ {
		if ct := j.CountFrom(math.MaxUint64/8, (math.MaxUint64/8)*7); int(ct) < n/2 {
			b.Errorf("Count should > %v, was %v", n/2, ct)
		}
	}
}

func BenchmarkCountRangeRand1000000(b *testing.B) {
	j := Judy1{}
	defer j.Free()

	n := 1000000
	for i := 0; i < n; i++ {
		j.Set(uint64(rand.Int63()))
	}

	for loops := 0; loops < b.N; loops++ {
		if ct := j.CountFrom(math.MaxUint64/8, (math.MaxUint64/8)*7); int(ct) < n/2 {
			b.Errorf("Count should > %v, was %v", n/2, ct)
		}
	}
}
