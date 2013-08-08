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

// This project is a cgo wrapper for Judy arrays implemented in libJudy.
// More information on the project can be found here: http://judy.sourceforge.net/

type Judy1 struct {
	array unsafe.Pointer
}

func (j *Judy1) Set(index uint64) {
	C.Judy1Set(C.PPvoid_t(&j.array), C.Word_t(index), nil)
}

func (j *Judy1) Unset(index uint64) {
	C.Judy1Unset(C.PPvoid_t(&j.array), C.Word_t(index), nil)
}

func (j *Judy1) Test(index uint64) bool {
	return C.Judy1Test(C.Pcvoid_t(j.array), C.Word_t(index), nil) != 0
}

func (j *Judy1) Free() uint64 {
	return uint64(C.Judy1FreeArray(C.PPvoid_t(&j.array), nil))
}

func (j *Judy1) CountAll() uint64 {
	return uint64(C.Judy1Count(C.Pcvoid_t(j.array), 0, math.MaxUint64, nil))
}

func (j *Judy1) CountFrom(indexA, indexB uint64) uint64 {
	return uint64(C.Judy1Count(C.Pcvoid_t(j.array), C.Word_t(indexA), C.Word_t(indexB), nil))
}

func (j *Judy1) MemoryUsed() uint64 {
	return uint64(C.Judy1MemUsed(C.Pcvoid_t(j.array)))
}
