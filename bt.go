package randtest

import (
	"runtime"
	"unsafe"
)

type (
	intArgs    uintptr
	int31Args  uintptr
	int31nArgs struct {
		_ uintptr
		n int32
	}
	int63Args  uintptr
	int63nArgs struct {
		_ uintptr
		n int64
	}
	intnArgs struct {
		_ uintptr
		n int
	}
	uint32Args  uintptr
	uint64Args  uintptr
	float32Args uintptr
	float64Args uintptr
	shuffleArgs struct {
		_    uintptr
		n    int
		swap func(int, int)
	}
	permArgs struct {
		_ uintptr
		n int
	}
)

type argtrace struct {
	Int     *intArgs
	Int31   *int31Args
	Int31n  *int31nArgs
	Int63   *int63Args
	Int63n  *int63nArgs
	Intn    *intnArgs
	Uint32  *uint32Args
	Uint64  *uint64Args
	Float32 *float32Args
	Float64 *float64Args
	Shuffle *shuffleArgs
	Perm    *permArgs
}

func getbp() uint64

func (s *Source) l() *argtrace {
	bp := getbp()

	pc := make([]uintptr, 10)
	npc := runtime.Callers(1, pc)
	pc = pc[:npc]
	frameIter := runtime.CallersFrames(pc)

	var frames []runtime.Frame
	for more := true; more; {
		var frame runtime.Frame
		frame, more = frameIter.Next()
		frames = append(frames, frame)
	}

	a := new(argtrace)
	a.backtrace(uintptr(bp), frames)

	return a
}

func (a *argtrace) backtrace(lbp uintptr, frames []runtime.Frame) {
	if len(frames) == 0 {
		return
	}
	funcname := frames[0].Function

	argp := lbp + 16
	switch funcname {
	case "math/rand.(*Rand).Int":
		a.Int = (*intArgs)(unsafe.Pointer(argp))
	case "math/rand.(*Rand).Intn":
		a.Intn = (*intnArgs)(unsafe.Pointer(argp))
	case "math/rand.(*Rand).Int31":
		a.Int31 = (*int31Args)(unsafe.Pointer(argp))
	case "math/rand.(*Rand).Int31n":
		a.Int31n = (*int31nArgs)(unsafe.Pointer(argp))
	case "math/rand.(*Rand).Int63":
		a.Int63 = (*int63Args)(unsafe.Pointer(argp))
	case "math/rand.(*Rand).Int63n":
		a.Int63n = (*int63nArgs)(unsafe.Pointer(argp))
	case "math/rand.(*Rand).Uint32":
		a.Uint32 = (*uint32Args)(unsafe.Pointer(argp))
	case "math/rand.(*Rand).Uint64":
		a.Uint64 = (*uint64Args)(unsafe.Pointer(argp))
	case "math/rand.(*Rand).Float32":
		a.Float32 = (*float32Args)(unsafe.Pointer(argp))
	case "math/rand.(*Rand).Float64":
		a.Float64 = (*float64Args)(unsafe.Pointer(argp))
	case "math/rand.(*Rand).Perm":
		a.Perm = (*permArgs)(unsafe.Pointer(argp))
	case "math/rand.(*Rand).Shuffle":
		a.Shuffle = (*shuffleArgs)(unsafe.Pointer(argp))
	}

	bp, ok := deref(lbp)
	if !ok {
		return
	}

	a.backtrace(bp, frames[1:])
}

func deref(p uintptr) (uintptr, bool) {
	defer func() { recover() }()
	return uintptr(*(*unsafe.Pointer)(unsafe.Pointer(p))), true
}
