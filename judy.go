package judy

/*
#cgo LDFLAGS: -lJudy
#include <Judy.h>

static void JudyLInsert(PPvoid_t PJLArray, Word_t index, Word_t value) {
  PPvoid_t val = JudyLIns(PJLArray, index, PJE0);
  *(Word_t *)val = value;
}

static Word_t JudyLFree(PPvoid_t PJLArray) {
  return JudyLFreeArray(PJLArray, PJE0);
}

static Word_t JudyLCt(Pvoid_t PJLArray, Word_t indexA, Word_t indexB) {
  return JudyLCount(PJLArray, indexA, indexB, PJE0);
}

static Word_t JudyLMemUsage(Pvoid_t PJLArray) {
  return JudyLMemUsed(PJLArray);
}

*/
import "C"

import (
	"math"
	"unsafe"
)

type JudyL struct {
	PJLArray unsafe.Pointer
}

func (jl *JudyL) Insert(index, value uint64) {
	C.JudyLInsert(&jl.PJLArray, C.Word_t(index), C.Word_t(value))
}

func (jl *JudyL) Free() uint64 {
	return uint64(C.JudyLFree(&jl.PJLArray))
}

func (jl *JudyL) CountAll() uint64 {
	return uint64(C.JudyLCt(C.Pvoid_t(jl.PJLArray), 0, math.MaxUint64))
}

func (jl *JudyL) CountFrom(indexA, indexB uint64) uint64 {
	return uint64(C.JudyLCt(C.Pvoid_t(jl.PJLArray), C.Word_t(indexA), C.Word_t(indexB)))
}

func (jl *JudyL) MemoryUsage() uint64 {
	return uint64(C.JudyLMemUsage(C.Pvoid_t(jl.PJLArray)))
}
