package randtest

import "math/rand"

var _ Behavior = (*rand.Rand)(nil)

var _ interface {
	Intn
	Int31n
	Int63n
	Int31
	Int63
	Uint32
	Uint64
	Float32
	Float64
} = (*MaxBehavior)(nil)

var _ interface {
	Intn
	Int31n
	Int63n
	Int31
	Int63
	Uint32
	Uint64
	Float32
	Float64
} = (*MinBehavior)(nil)
