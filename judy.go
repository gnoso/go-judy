// This project is a cgo wrapper for Judy arrays found in libjudy.
// More information on Judy arrays can be found here: http://judy.sourceforge.net/
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

// Wrapper for the Judy1 array pointer. The default value of this struct is a valid empty Judy1 array.
type Judy1 struct {
	array unsafe.Pointer
}

// Set index's bit in the Judy1 array
// Return true if index's bit was previously unset (successful), otherwise false if the bit was already set (unsuccessful).
func (j *Judy1) Set(index uint64) bool {
	return C.Judy1Set(C.PPvoid_t(&j.array), C.Word_t(index), nil) != 0
}

// Unset index's bit in the Judy1 array
// Return true if index's bit was previously set (successful), otherwise false if the bit was already unset (unsuccessful).
func (j *Judy1) Unset(index uint64) bool {
	return C.Judy1Unset(C.PPvoid_t(&j.array), C.Word_t(index), nil) != 0
}

// Test if index's bit is set in the Judy1 array.
// Return true if index's bit is set (index is present), false if it is unset (index is absent).
func (j *Judy1) Test(index uint64) bool {
	return C.Judy1Test(C.Pcvoid_t(j.array), C.Word_t(index), nil) != 0
}

// Free the entire Judy1 array
// Return the number of bytes freed.
func (j *Judy1) Free() uint64 {
	return uint64(C.Judy1FreeArray(C.PPvoid_t(&j.array), nil))
}

// Count the number of indexes present in the Judy1 array.
// Returns the count. A return value of 0 can be valid as a count, or it can indicate a special case for fully populated array (32-bit machines only). See libjudy docs for ways to resolve this.
func (j *Judy1) CountAll() uint64 {
	return uint64(C.Judy1Count(C.Pcvoid_t(j.array), 0, math.MaxUint64, nil))
}

// Count the number of indexes present in the Judy1 array between indexA and indexB (inclusive).
// Returns the count. A return value of 0 can be valid as a count, or it can indicate a special case for fully populated array (32-bit machines only). See libjudy docs for ways to resolve this.
func (j *Judy1) CountFrom(indexA, indexB uint64) uint64 {
	return uint64(C.Judy1Count(C.Pcvoid_t(j.array), C.Word_t(indexA), C.Word_t(indexB), nil))
}

// Return the number of bytes of memory currently in use by Judy1 array. This is a very fast routine, and may be used with little performance impact.
func (j *Judy1) MemoryUsed() uint64 {
	return uint64(C.Judy1MemUsed(C.Pcvoid_t(j.array)))
}
