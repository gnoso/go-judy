package judy

import (
	"math"
	"math/rand"
	"testing"
)

func TestEmptyJudyLArray(t *testing.T) {

	j := JudyL{}
	r := j.Free()

	if r != 0 {
		t.Errorf("Free should return 0, returned %v", r)
	}
}

func TestJudyLCount(t *testing.T) {

	j := JudyL{}
	defer j.Free()

	if ct := j.CountAll(); ct != 0 {
		t.Errorf("Count should be zero, was %v", ct)
	}

	var i uint64
	for i = 0; i < 100; i++ {
		j.Insert(i, i)
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

func TestJudyLInsertGet(t *testing.T) {

	j := JudyL{}
	defer j.Free()

	var i uint64
	for i = 0; i < 100; i++ {
		j.Insert(i*100000, i)
	}

	for i = 0; i < 100; i++ {
		if v, ok := j.Get(i * 100000); ok && v != i {
			t.Errorf("Index %v should be %v, but was %v", i*100000, i, v)
		}
	}

	for i = 1; i < 100; i++ {
		if _, ok := j.Get(i * 99999); ok {
			t.Errorf("Index %v incorrectly set", i*99999)
		}
	}

}

func TestJudyLByCount(t *testing.T) {

	j := JudyL{}
	defer j.Free()

	j.Insert(12345, 234)
	j.Insert(11235, 11235)
	j.Insert(54321, 4321)

	if idx, val, ok := j.ByCount(1); ok && (idx != 11235 || val != 11235) {
		t.Errorf("ByCount should return 11235,11235,true but was %v, %v, %v", idx, val, ok)
	}
	if idx, val, ok := j.ByCount(2); ok && (idx != 12345 || val != 234) {
		t.Errorf("ByCount should return 12345,234,true but was %v, %v, %v", idx, val, ok)
	}
	if idx, val, ok := j.ByCount(3); ok && (idx != 54321 || val != 4321) {
		t.Errorf("ByCount should return 54321,4321,true but was %v, %v, %v", idx, val, ok)
	}
	if _, _, ok := j.ByCount(0); ok {
		t.Error("There should be no return value for 0")
	}
	if _, _, ok := j.ByCount(4); ok {
		t.Error("There should be no return value for 4")
	}
}

func TestJudyLDelete(t *testing.T) {

	j := JudyL{}
	defer j.Free()

	j.Insert(12345, 234)
	j.Insert(11235, 11235)
	j.Insert(54321, 4321)

	if ct := j.CountAll(); ct != 3 {
		t.Errorf("Count should be 3")
	}
	if ok := j.Delete(11235); !ok {
		t.Errorf("Delete should return ok")
	}
	if ct := j.CountAll(); ct != 2 {
		t.Errorf("Count should be 2")
	}
	if _, ok := j.Get(11235); ok {
		t.Errorf("Value should be removed")
	}
	if ok := j.Delete(11235); ok {
		t.Errorf("Delete not should return ok")
	}

}

func TestJudyLFirst(t *testing.T) {

	j := JudyL{}
	defer j.Free()

	var i uint64
	for i = 0; i < 100; i++ {
		j.Insert(i*2, i)
	}

	if next, val, ok := j.First(20); ok && (next != 20 || val != 10) {
		t.Errorf("First(20) should be 20, 10 was %v,%v", next, val)
	}
	if next, val, ok := j.First(21); ok && (next != 22 || val != 11) {
		t.Errorf("First(21) should be 22, 11 was %v,%v", next, val)
	}
	if _, _, ok := j.First(201); ok {
		t.Errorf("First(201) should not be found")
	}

}

func TestJudyLLast(t *testing.T) {

	j := JudyL{}
	defer j.Free()

	var i uint64
	for i = 1; i < 100; i++ {
		j.Insert(i*2, i)
	}

	if next, val, ok := j.Last(20); ok && (next != 20 || val != 10) {
		t.Errorf("Last(20) should be 20,10 was %v,%v", next, val)
	}
	if next, val, ok := j.Last(21); ok && (next != 20 || val != 10) {
		t.Errorf("Last(21) should be 20, 10 was %v,%v", next, val)
	}
	if _, _, ok := j.Last(1); ok {
		t.Errorf("Last(1) should not be found")
	}
}

func TestJudyLNext(t *testing.T) {

	j := JudyL{}
	defer j.Free()

	var i uint64
	for i = 0; i < 100; i++ {
		j.Insert(i*2, i)
	}

	if next, val, ok := j.Next(20); ok && (next != 22 || val != 11) {
		t.Errorf("Next(20) should be 22,11 was %v,%v", next, val)
	}
	if next, val, ok := j.Next(21); ok && (next != 22 || val != 11) {
		t.Errorf("Next(21) should be 22,11 was %v,%v", next, val)
	}
	if _, _, ok := j.Next(200); ok {
		t.Errorf("Next(200) should not be found")
	}

}

func TestJudyLPrev(t *testing.T) {

	j := JudyL{}
	defer j.Free()

	var i uint64
	for i = 1; i < 100; i++ {
		j.Insert(i*2, i)
	}

	if next, val, ok := j.Prev(20); ok && (next != 18 || val != 9) {
		t.Errorf("Prev(20) should be 18,9 was %v,%v", next, val)
	}
	if next, val, ok := j.Prev(21); ok && (next != 20 || val != 10) {
		t.Errorf("Prev(21) should be 20,10 was %v", next, val)
	}
	if _, _, ok := j.Prev(2); ok {
		t.Errorf("Prev(2) should not be found")
	}

}

func runOrderedJudyLMemUsageTest(t *testing.T, n int) {
	j := JudyL{}
	defer j.Free()

	for i := 0; i < n; i++ {
		j.Insert(uint64(i*10000), uint64(i))
	}

	if ct := j.CountAll(); int(ct) != n {
		t.Errorf("Count should be %v, was %v", n, ct)
	}
	t.Logf("Memory Usage with %7v ordered bits %8v", n, j.MemoryUsed())
}

func runRandomJudyLMemUsageTest(t *testing.T, n int) {
	j := JudyL{}
	defer j.Free()

	for i := 0; i < n; i++ {
		j.Insert(uint64(rand.Int63()), uint64(rand.Int63()))
	}

	if ct := j.CountAll(); int(ct) != n {
		t.Errorf("Count should be %v, was %v", n, ct)
	}
	t.Logf("Memory Usage with %7v random bits  %8v", n, j.MemoryUsed())
}

func TestMemUsage(t *testing.T) {

	runOrderedJudyLMemUsageTest(t, 1000)
	runRandomJudyLMemUsageTest(t, 1000)
	runOrderedJudyLMemUsageTest(t, 10000)
	runRandomJudyLMemUsageTest(t, 10000)
	runOrderedJudyLMemUsageTest(t, 100000)
	runRandomJudyLMemUsageTest(t, 100000)
	runOrderedJudyLMemUsageTest(t, 1000000)
	runRandomJudyLMemUsageTest(t, 1000000)

	//t.Fail() // Uncomment to see the log output
}

func BenchmarkJudyLCountAllRand1000(b *testing.B) {
	j := JudyL{}
	defer j.Free()

	n := 1000
	for i := 0; i < n; i++ {
		j.Insert(uint64(rand.Int63()), uint64(rand.Int63()))
	}

	for loops := 0; loops < b.N; loops++ {
		if ct := j.CountAll(); int(ct) != n {
			b.Errorf("Count should be %v, was %v", n, ct)
		}
	}
}

func BenchmarkJudyLCountAllRand1000000(b *testing.B) {
	j := JudyL{}
	defer j.Free()

	n := 1000000
	for i := 0; i < n; i++ {
		j.Insert(uint64(rand.Int63()), uint64(rand.Int63()))
	}

	for loops := 0; loops < b.N; loops++ {
		if ct := j.CountAll(); int(ct) != n {
			b.Errorf("Count should be %v, was %v", n, ct)
		}
	}
}

func BenchmarkJudyLCountAllOrd1000(b *testing.B) {
	j := JudyL{}
	defer j.Free()

	n := 1000
	for i := 0; i < n; i++ {
		j.Insert(uint64(i), uint64(i))
	}

	for loops := 0; loops < b.N; loops++ {
		if ct := j.CountAll(); int(ct) != n {
			b.Errorf("Count should be %v, was %v", n, ct)
		}
	}
}

func BenchmarkJudyLCountAllOrd1000000(b *testing.B) {
	j := JudyL{}
	defer j.Free()

	n := 1000000
	for i := 0; i < n; i++ {
		j.Insert(uint64(i), uint64(i))
	}

	for loops := 0; loops < b.N; loops++ {
		if ct := j.CountAll(); int(ct) != n {
			b.Errorf("Count should be %v, was %v", n, ct)
		}
	}
}

func BenchmarkJudyLCountRangeRand1000(b *testing.B) {
	j := JudyL{}
	defer j.Free()

	n := 1000
	for i := 0; i < n; i++ {
		j.Insert(uint64(rand.Int63()), uint64(rand.Int63()))
	}

	for loops := 0; loops < b.N; loops++ {
		if ct := j.CountFrom(math.MaxUint64/8, (math.MaxUint64/8)*7); int(ct) < n/2 {
			b.Errorf("Count should > %v, was %v", n/2, ct)
		}
	}
}

func BenchmarkJudyLCountRangeRand1000000(b *testing.B) {
	j := JudyL{}
	defer j.Free()

	n := 1000000
	for i := 0; i < n; i++ {
		j.Insert(uint64(rand.Int63()), uint64(rand.Int63()))
	}

	for loops := 0; loops < b.N; loops++ {
		if ct := j.CountFrom(math.MaxUint64/8, (math.MaxUint64/8)*7); int(ct) < n/2 {
			b.Errorf("Count should > %v, was %v", n/2, ct)
		}
	}
}
