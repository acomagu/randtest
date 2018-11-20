package randtest

import (
	"math"
	"math/rand"
)

// Behavior defines the expected behavior of rand.Rand struct given Source as
// the source. To define the behavior of a method, implement a method has same
// signature as rand.Rand. Each method can be implemented are also defined
// as interface in this package, Intn, Uint32, etc. All methods are optional.
// If a method is omitted, the behavior uses the top level methods of rand
// package.
type Behavior interface{}

type (
	// Intn is the interface implemented by a Behavior to define the behavior of
	// rand.Rand.Intn.
	Intn interface {
		Intn(int) int
	}
	// Int31n is the interface implemented by a Behavior to define the behavior of
	// rand.Rand.Int31n.
	Int31n interface {
		Int31n(int32) int32
	}
	// Int63n is the interface implemented by a Behavior to define the behavior of
	// rand.Rand.Int63n.
	Int63n interface {
		Int63n(int64) int64
	}
	// Int31 is the interface implemented by a Behavior to define the behavior of
	// rand.Rand.Int31.
	Int31 interface {
		Int31() int32
	}
	// Int63 is the interface implemented by a Behavior to define the behavior of
	// rand.Rand.Int63.
	Int63 interface {
		Int63() int64
	}
	// Uint32 is the interface implemented by a Behavior to define the behavior of
	// rand.Rand.Uint32.
	Uint32 interface {
		Uint32() uint32
	}
	// Uint64 is the interface implemented by a Behavior to define the behavior of
	// rand.Rand.Uint64.
	Uint64 interface {
		Uint64() uint64
	}
	// Float32 is the interface implemented by a Behavior to define the behavior
	// of rand.Rand.Float32.
	Float32 interface {
		Float32() float32
	}
	// Float64 is the interface implemented by a Behavior to define the behavior
	// of rand.Rand.Float64.
	Float64 interface {
		Float64() float64
	}
	// Perm is the interface implemented by a Behavior to define the behavior of
	// rand.Rand.Perm.
	Perm interface {
		Perm(int) []int
	}
	// Shuffle is the interface implemented by a Behavior to define the behavior
	// of rand.Rand.Shuffle.
	Shuffle interface {
		Shuffle(int, func(int, int))
	}
)

type behavior struct {
	b Behavior
}

func (b *behavior) Intn(n int) int {
	if w, ok := b.b.(Intn); ok {
		return w.Intn(n)
	}
	return rand.Intn(n)
}

func (b *behavior) Int31n(n int32) int32 {
	if w, ok := b.b.(Int31n); ok {
		return w.Int31n(n)
	}
	return rand.Int31n(n)
}

func (b *behavior) Int63n(n int64) int64 {
	if w, ok := b.b.(Int63n); ok {
		return w.Int63n(n)
	}
	return rand.Int63n(n)
}

func (b *behavior) Int31() int32 {
	if w, ok := b.b.(Int31); ok {
		return w.Int31()
	}
	return rand.Int31()
}

func (b *behavior) Int63() int64 {
	if w, ok := b.b.(Int63); ok {
		return w.Int63()
	}
	return rand.Int63()
}

func (b *behavior) Uint32() uint32 {
	if w, ok := b.b.(Uint32); ok {
		return w.Uint32()
	}
	return rand.Uint32()
}

func (b *behavior) Uint64() uint64 {
	if w, ok := b.b.(Uint64); ok {
		return w.Uint64()
	}
	return rand.Uint64()
}

func (b *behavior) Float32() float32 {
	if w, ok := b.b.(Float32); ok {
		return w.Float32()
	}
	return rand.Float32()
}

func (b *behavior) Float64() float64 {
	if w, ok := b.b.(Float64); ok {
		return w.Float64()
	}
	return rand.Float64()
}

func (b *behavior) Perm(n int) []int {
	if w, ok := b.b.(Perm); ok {
		return w.Perm(n)
	}
	return rand.Perm(n)
}

func (b *behavior) Shuffle(n int, swap func(int, int)) {
	if w, ok := b.b.(Shuffle); ok {
		w.Shuffle(n, swap)
		return
	}
	rand.Shuffle(n, swap)
}

// MaxBehavior is a Behavior always returns the largest possible value. The
// zero value is ready to work.
type MaxBehavior struct{}

// Intn returns n-1.
func (*MaxBehavior) Intn(n int) int {
	return n - 1
}

// Int31n returns n-1.
func (*MaxBehavior) Int31n(n int32) int32 {
	return n - 1
}

// Int63n returns n-1.
func (*MaxBehavior) Int63n(n int64) int64 {
	return n - 1
}

// Int31 returns MaxInt32-1.
func (*MaxBehavior) Int31() int32 {
	return math.MaxInt32 - 1
}

// Int63 returns MaxInt64-1.
func (*MaxBehavior) Int63() int64 {
	return math.MaxInt64 - 1
}

// Uint64 returns MaxUint64-1.
func (*MaxBehavior) Uint64() uint64 {
	return math.MaxUint64
}

// Uint32 returns MaxUint32-1.
func (*MaxBehavior) Uint32() uint32 {
	return math.MaxUint32
}

// Float64 returns 0.9999999999999999.
func (*MaxBehavior) Float64() float64 {
	// The bits expression is 0x3fefffffffffffff.
	return 0.9999999999999999
}

// Float32 returns 0.99999994.
func (*MaxBehavior) Float32() float32 {
	// The bits expression is 0x3f7fffff.
	return 0.99999994
}

// MinBehavior is a Behavior always returns the smallest possible value. The
// zero value is ready to work.
type MinBehavior struct{}

// Intn returns 0.
func (*MinBehavior) Intn(n int) int {
	return 0
}

// Int31n returns 0.
func (*MinBehavior) Int31n(n int32) int32 {
	return 0
}

// Int63n returns 0.
func (*MinBehavior) Int63n(n int64) int64 {
	return 0
}

// Int31 returns 0.
func (*MinBehavior) Int31() int32 {
	return 0
}

// Int63 returns 0.
func (*MinBehavior) Int63() int64 {
	return 0
}

// Uint64 returns 0.
func (*MinBehavior) Uint64() uint64 {
	return 0
}

// Uint32 returns 0.
func (*MinBehavior) Uint32() uint32 {
	return 0
}

// Float32 returns 0.
func (*MinBehavior) Float32() float32 {
	return 0
}

// Float64 returns 0.
func (*MinBehavior) Float64() float64 {
	return 0
}
