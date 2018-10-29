package randtest

import (
	"math"
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
	uint32Args uintptr
	uint64Args uintptr
)

type Source struct{}

type argtrace struct {
	Int    *intArgs
	Int31  *int31Args
	Int31n *int31nArgs
	Int63  *int63Args
	Int63n *int63nArgs
	Intn   *intnArgs
	Uint32 *uint32Args
	Uint64 *uint64Args
}

func NewSource() *Source {
	return &Source{}
}

func (s *Source) Int63() int64 {
	args := s.l()

	switch {
	case args.Intn != nil:
		return s.intn(args.Intn.n)
	case args.Int31n != nil:
		return s.int31n(args.Int31n.n)
	case args.Int63n != nil:
		return s.int63n(args.Int63n.n)
	case args.Int31 != nil:
		return s.int31(math.MaxInt32)
	case args.Int63 != nil:
		return s.int63(math.MaxInt64)
	}

	return 0
}

func (s *Source) intn(n int) int64 {
	if n <= 1<<31-1 {
		return s.int31n(int32(n))
	}
	return s.int63n(int64(n))
}

func (s *Source) int63(n int64) int64 {
	return n - 1
}

func (s *Source) int63n(n int64) int64 {
	return s.int63(n)
}

func (s *Source) int31(n int32) int64 {
	return int64(n-1) << 32
}

func (s *Source) int31n(n int32) int64 {
	return s.int31(n)
}

func (s *Source) Seed(int64) {}
