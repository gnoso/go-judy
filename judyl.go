package judy

/*
#cgo LDFLAGS: -lJudy
#include <Judy.h>
*/
import "C"

import (
	"math"
	"unsafe"
)

// A JudyL array is the equivalent of a dynamic array of uint64 values. A value is addressed by an index (key).
// The array may be sparse, and the index may be any uint64 number.
//
// The default value of this struct is a valid empty JudyL array.
//
//    j := JudyL{}
//    defer j.Free()
//
//    j.Insert(5142, 142)
//    fmt.Printf("Number of items: %v", j.CountAll())
//
//
// Memory to support the array is allocated as index/value pairs are inserted, and released as index/value
// pairs are deleted. A JudyL array can also be thought of as a mapper, that is "map" a word to another
// word. If the JudyL array is freed ( by calling .Free() ), all words are removed (and the JudyL array requires no memory).
// As with an ordinary array, a JudyL array contains no duplicate indexes.
//
// NOTE: The Judy array is implemented in C and allocates memory directly from the operating system. It is NOT
// garbage collected by the Go runtime. It is very important that you call Free() on a Judy array after using
// it to prevent memory leaks. The "defer" pattern is a great way to accomplish this.
type JudyL struct {
	array unsafe.Pointer
}

// Insert an Index and Value into the JudyL array. If the Index is successfully inserted, the Value is
// initialized as well. If the Index was already present, the current Value is replaced with the provided Value.
func (j *JudyL) Insert(index uint64, value uint64) {
	pval := unsafe.Pointer(C.JudyLIns(C.PPvoid_t(&j.array), C.Word_t(index), nil))
	*((*C.Word_t)(pval)) = C.Word_t(value)
}

// Delete the Index/Value pair from the JudyL array.
// Returns true if successful. Returns false if Index was not present.
func (j *JudyL) Delete(index uint64) bool {
	return C.JudyLDel(C.PPvoid_t(&j.array), C.Word_t(index), nil) != 0
}

// Get the Value associated with Index in the Judy array
//   returns (value, true) if the index was found
//   returns (_, false) if the index was not found
func (j *JudyL) Get(index uint64) (uint64, bool) {
	pval := unsafe.Pointer(C.JudyLGet(C.Pcvoid_t(j.array), C.Word_t(index), nil))
	if pval == nil {
		return 0, false
	} else {
		return uint64(*((*C.Word_t)(pval))), true
	}
}

// Free the entire JudyL array.
// Return the number of bytes freed.
//
// NOTE: The Judy array allocates memory directly from the operating system and is NOT garbage collected by the
// Go runtime. It is very important that you call Free() on a Judy array after using it to prevent memory leaks.
func (j *JudyL) Free() uint64 {
	return uint64(C.JudyLFreeArray(C.PPvoid_t(&j.array), nil))
}

// Count the number of indexes present in the JudyL array.
// Returns the count. A return value of 0 can be valid as a count, or it can indicate a special case for fully populated array (32-bit machines only). See libjudy docs for ways to resolve this.
func (j *JudyL) CountAll() uint64 {
	return uint64(C.JudyLCount(C.Pcvoid_t(j.array), 0, math.MaxUint64, nil))
}

// Count the number of indexes present in the JudyL array between indexA and indexB (inclusive).
// Returns the count. A return value of 0 can be valid as a count, or it can indicate a special case for fully populated array (32-bit machines only). See libjudy docs for ways to resolve this.
func (j *JudyL) CountFrom(indexA, indexB uint64) uint64 {
	return uint64(C.JudyLCount(C.Pcvoid_t(j.array), C.Word_t(indexA), C.Word_t(indexB), nil))
}

// Return the number of bytes of memory currently in use by JudyL array. This is a very fast routine,
// and may be used with little performance impact.
func (j *JudyL) MemoryUsed() uint64 {
	return uint64(C.JudyLMemUsed(C.Pcvoid_t(j.array)))
}

// Search (inclusive) for the first index present that is equal to or greater than the passed index.
// (Start with index = 0 to find the first index in the array.) This is typically used to begin a sorted-order scan of the indexes present in a JudyL array.
//
//   index - search index
//   returns uint64 - value of the first index that is equal to or greater than the passed index
//                    (only if bool return value is true)
//           uint64 - value pointed to by the index
//           bool   - true if the search was successful, false if an index was not found
func (j *JudyL) First(index uint64) (uint64, uint64, bool) {
	idx := C.Word_t(index)
	pval := unsafe.Pointer(C.JudyLFirst(C.Pcvoid_t(j.array), &idx, nil))

	if pval == nil {
		return 0, 0, false
	} else {
		return uint64(idx), uint64(*((*C.Word_t)(pval))), true
	}
}

// Search (exclusive) for the first index present that is greater than the passed index.
// This is typically used to continue a sorted-order scan of the indexes present in a JudyL array.
//
//   index - search index
//   returns uint64 - value of the first index that is greater than the passed index
//                    (only if bool return value is true)
//           uint64 - value pointed to by the index
//           bool   - true if the search was successful, false if an index was not found
func (j *JudyL) Next(index uint64) (uint64, uint64, bool) {
	idx := C.Word_t(index)
	pval := unsafe.Pointer(C.JudyLNext(C.Pcvoid_t(j.array), &idx, nil))

	if pval == nil {
		return 0, 0, false
	} else {
		return uint64(idx), uint64(*((*C.Word_t)(pval))), true
	}
}

// Search (inclusive) for the last index present that is equal to or less than than the passed index.
// (Start with index = math.MaxUint64 to find the last index in the array.) This is typically used to begin a reverse-sorted-order scan of the indexes present in a JudyL array.
//
//   index - search index
//   returns uint64 - value of the last index that is equal to or less than the passed index
//                    (only if bool return value is true)
//           uint64 - value pointed to by the index
//           bool   - true if the search was successful, false if an index was not found
func (j *JudyL) Last(index uint64) (uint64, uint64, bool) {
	idx := C.Word_t(index)
	pval := unsafe.Pointer(C.JudyLLast(C.Pcvoid_t(j.array), &idx, nil))

	if pval == nil {
		return 0, 0, false
	} else {
		return uint64(idx), uint64(*((*C.Word_t)(pval))), true
	}
}

// Search (exclusive) for the last index present that is less than the passed index.
// This is typically used to continue a reverse sorted-order scan of the indexes present in a JudyL array.
//
//   index - search index
//   returns uint64 - value of the last index that is less than the passed index
//                    (only if bool return value is true)
//           uint64 - value pointed to by the index
//           bool   - true if the search was successful, false if an index was not found
func (j *JudyL) Prev(index uint64) (uint64, uint64, bool) {
	idx := C.Word_t(index)
	pval := unsafe.Pointer(C.JudyLPrev(C.Pcvoid_t(j.array), &idx, nil))

	if pval == nil {
		return 0, 0, false
	} else {
		return uint64(idx), uint64(*((*C.Word_t)(pval))), true
	}
}

// Locate the Nth index that is present in the JudyL array (Nth = 1 returns the first index present).
//
//   nth - nth index to find
//   returns uint64 - nth index (unless return false)
//           uint64 - nth value (unless return false)
//           bool   - true if the search was successful, false if an index was not found
func (j *JudyL) ByCount(nth uint64) (uint64, uint64, bool) {
	var idx C.Word_t
	pval := unsafe.Pointer(C.JudyLByCount(C.Pcvoid_t(j.array), C.Word_t(nth), &idx, nil))

	if pval == nil {
		return 0, 0, false
	} else {
		return uint64(idx), uint64(*((*C.Word_t)(pval))), true
	}
}
