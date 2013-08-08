package judy

import (
	"math/rand"
	"testing"
)

func TestEmptyJudyArray(t *testing.T) {

	j := JudyL{}
	r := j.Free()

	if r != 0 {
		t.Errorf("Free should return 0, returned %v", r)
	}
}

func TestInsertCount(t *testing.T) {

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

func TestMemUsage(t *testing.T) {

	j := JudyL{}

	var i uint64
	var ct uint64

	for i = 0; i < 1000; i++ {
		j.Insert(i, i)
	}

	if ct = j.CountAll(); ct != 1000 {
		t.Errorf("Count should be 1000, was %v", ct)
	}
	t.Logf("Memory Usage with 1000 uint64s %v", j.MemoryUsage())
	j.Free()

	j = JudyL{}
	for i = 0; i < 10000; i++ {
		j.Insert(i, i)
	}

	if ct = j.CountAll(); ct != 10000 {
		t.Errorf("Count should be 10000, was %v", ct)
	}
	t.Logf("Memory Usage with 10000 ordered uint64s %v", j.MemoryUsage())
	j.Free()

	j = JudyL{}
	for i = 0; i < 10000; i++ {
		j.Insert(uint64(rand.Int63()), uint64(rand.Int63()))
	}

	t.Logf("Memory Usage with 10000 random uint64s %v", j.MemoryUsage())
	j.Free()

	j = JudyL{}
	for i = 0; i < 100000; i++ {
		j.Insert(i, i)
	}

	if ct = j.CountAll(); ct != 100000 {
		t.Errorf("Count should be 100000, was %v", ct)
	}
	t.Logf("Memory Usage with 100000 ordered uint64s %v", j.MemoryUsage())
	j.Free()

	j = JudyL{}
	for i = 0; i < 100000; i++ {
		j.Insert(uint64(rand.Int63()), uint64(rand.Int63()))
	}

	t.Logf("Memory Usage with 100000 random uint64s %v", j.MemoryUsage())
	j.Free()

	j = JudyL{}
	for i = 0; i < 1000000; i++ {
		j.Insert(uint64(rand.Int63()), uint64(rand.Int63()))
	}

	t.Logf("Memory Usage with 1000000 random uint64s %v", j.MemoryUsage())
	j.Free()

	t.Fail()
}
