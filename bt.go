package randtest

import (
	"runtime"
	"unsafe"
)

func getbp() uint64

func (s *Source) l() *argtrace {
	bp := getbp()

	pc := make([]uintptr, 10)
	npc := runtime.Callers(1, pc)
	pc = pc[:npc]
	frameIter := runtime.CallersFrames(pc)

	var frames []runtime.Frame
	more := true
	for more {
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
	case "math/rand.(*Rand).Int31":
		a.Int31 = (*int31Args)(unsafe.Pointer(argp))
	case "math/rand.(*Rand).Int31n":
		a.Int31n = (*int31nArgs)(unsafe.Pointer(argp))
	case "math/rand.(*Rand).Int63":
		a.Int63 = (*int63Args)(unsafe.Pointer(argp))
	case "math/rand.(*Rand).Int63n":
		a.Int63n = (*int63nArgs)(unsafe.Pointer(argp))
	case "math/rand.(*Rand).Intn":
		a.Intn = (*intnArgs)(unsafe.Pointer(argp))
	case "math/rand.(*Rand).Uint32":
		a.Uint32 = (*uint32Args)(unsafe.Pointer(argp))
	case "math/rand.(*Rand).Uint64":
		a.Uint64 = (*uint64Args)(unsafe.Pointer(argp))
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
