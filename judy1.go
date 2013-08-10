// Go language wrapper for Judy arrays (as found at http://judy.sourceforge.net)
//
// Judy arrays are a fast and memory efficient dynamic array structure. Judy arrays were invented by Doug Baskins
// and implemented by Hewlett-Packard.
//
// Judy is designed to avoid cache-line fills wherever possible. There are several different variants of Judy
// arrays. This package implements the Judy1 bitvector and the JudyL integer map currently. Adding other
// variants should be relatively simple, however.
//
// Counting and range counting operations are particularly fast, and do not require a scan of the array.
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

// A Judy1 array is the equivalent of a bit array or bit map. A bit is addressed by an index (key). The array may be sparse, and the index is a uint64 value. If an index is present, it represents a set bit (a bit set represents an index present). If an index is absent, it represents an unset bit (a bit unset represents an absent index).
// The default value of this struct is a valid empty Judy1 array.
//
//    j := Judy1{}
//    defer j.Free()
//
//    j.Set(5142)
//    fmt.Printf("Number of items: %v", j.CountAll())
//
//
// Memory to support the array is allocated as bits are set, and released as bits are unset. If the Judy1 array is freed ( by calling .Free() ), all bits are unset (and the Judy1 array requires no memory).
// As with an ordinary array, a Judy1 array contains no duplicate indexes.
//
// NOTE: The Judy array is implemented in C and allocates memory directly from the operating system. It is NOT
// garbage collected by the Go runtime. It is very important that you call Free() on a Judy array after using
// it to prevent memory leaks. The "defer" pattern is a great way to accomplish this.
type Judy1 struct {
	array unsafe.Pointer
}

// Set index's bit in the Judy1 array.
// Return true if index's bit was previously unset (successful), otherwise false if the bit was already set (unsuccessful).
func (j *Judy1) Set(index uint64) bool {
	return C.Judy1Set(C.PPvoid_t(&j.array), C.Word_t(index), nil) != 0
}

// Unset index's bit in the Judy1 array.
// Return true if index's bit was previously set (successful), otherwise false if the bit was already unset (unsuccessful).
func (j *Judy1) Unset(index uint64) bool {
	return C.Judy1Unset(C.PPvoid_t(&j.array), C.Word_t(index), nil) != 0
}

// Test if index's bit is set in the Judy1 array.
// Return true if index's bit is set (index is present), false if it is unset (index is absent).
func (j *Judy1) Test(index uint64) bool {
	return C.Judy1Test(C.Pcvoid_t(j.array), C.Word_t(index), nil) != 0
}

// Free the entire Judy1 array.
// Return the number of bytes freed.
//
// NOTE: The Judy array allocates memory directly from the operating system and is NOT garbage collected by the
// Go runtime. It is very important that you call Free() on a Judy array after using it to prevent memory leaks.
func (j *Judy1) Free() uint64 {
	return uint64(C.Judy1FreeArray(C.PPvoid_t(&j.array), nil))
}

// Count the number of indexes present in the Judy1 array.
// A return value of 0 can be valid as a count, or it can indicate a special case for fully populated array (32-bit machines only). See libjudy docs for ways to resolve this.
func (j *Judy1) CountAll() uint64 {
	return uint64(C.Judy1Count(C.Pcvoid_t(j.array), 0, math.MaxUint64, nil))
}

// Count the number of indexes present in the Judy1 array between indexA and indexB (inclusive).
// A return value of 0 can be valid as a count, or it can indicate a special case for fully populated array (32-bit machines only). See libjudy docs for ways to resolve this.
func (j *Judy1) CountFrom(indexA, indexB uint64) uint64 {
	return uint64(C.Judy1Count(C.Pcvoid_t(j.array), C.Word_t(indexA), C.Word_t(indexB), nil))
}

// Return the number of bytes of memory currently in use by Judy1 array. This is a very fast routine, and may be used with little performance impact.
func (j *Judy1) MemoryUsed() uint64 {
	return uint64(C.Judy1MemUsed(C.Pcvoid_t(j.array)))
}

// Search (inclusive) for the first index present that is equal to or greater than the passed index.
// (Start with index = 0 to find the first index in the array.) This is typically used to begin a sorted-order scan of the indexes present in a Judy1 array.
//
//   index - search index
//   returns uint64 - value of the first index that is equal to or greater than the passed index (only if bool return value is true)
//           bool   - true if the search was successful, false if an index was not found
func (j *Judy1) First(index uint64) (uint64, bool) {
	var idx C.Word_t = C.Word_t(index)

	if C.Judy1First(C.Pcvoid_t(j.array), &idx, nil) != 0 {
		return uint64(idx), true
	} else {
		return 0, false
	}
}

// Search (exclusive) for the first index present that is greater than the passed index.
// This is typically used to continue a sorted-order scan of the indexes present in a Judy1 array.
//
//   index - search index
//   returns uint64 - value of the first index that is greater than the passed index (only if bool return value is true)
//           bool   - true if the search was successful, false if an index was not found
func (j *Judy1) Next(index uint64) (uint64, bool) {
	var idx C.Word_t = C.Word_t(index)

	if C.Judy1Next(C.Pcvoid_t(j.array), &idx, nil) != 0 {
		return uint64(idx), true
	} else {
		return 0, false
	}
}

// Search (inclusive) for the last index present that is equal to or less than than the passed index.
// (Start with index = math.MaxUint64 to find the last index in the array.) This is typically used to begin a reverse-sorted-order scan of the indexes present in a Judy1 array.
//
//   index - search index
//   returns uint64 - value of the last index that is equal to or less than the passed index (only if bool return value is true)
//           bool   - true if the search was successful, false if an index was not found
func (j *Judy1) Last(index uint64) (uint64, bool) {
	var idx C.Word_t = C.Word_t(index)

	if C.Judy1Last(C.Pcvoid_t(j.array), &idx, nil) != 0 {
		return uint64(idx), true
	} else {
		return 0, false
	}
}

// Search (exclusive) for the last index present that is less than the passed index.
// This is typically used to continue a reverse sorted-order scan of the indexes present in a Judy1 array.
//
//   index - search index
//   returns uint64 - value of the last index that is less than the passed index (only if bool return value is true)
//           bool   - true if the search was successful, false if an index was not found
func (j *Judy1) Prev(index uint64) (uint64, bool) {
	var idx C.Word_t = C.Word_t(index)

	if C.Judy1Prev(C.Pcvoid_t(j.array), &idx, nil) != 0 {
		return uint64(idx), true
	} else {
		return 0, false
	}
}
