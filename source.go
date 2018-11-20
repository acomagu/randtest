package randtest

import (
	"math/rand"
)

// Source implements rand.Source and rand.Source64. It controls behaviors of
// each methods of rand.Rand by being used as source.
type Source struct {
	b           *behavior
	def         rand.Source64
	permrets    []int64
	shufflerets []int64
}

// NewSource returns a new Source. The b is used to define the behavior of
// rand.Rand given the Source.
func NewSource(b Behavior) *Source {
	return &Source{
		b: &behavior{
			b: b,
		},
		def: rand.New(rand.NewSource(0)),
	}
}

// Int63 implements rand.Source.
func (s *Source) Int63() int64 {
	args := s.l()

	switch {
	case args.Perm != nil:
		return s.int63ForPerm(args.Perm.n, s.b.Perm(args.Perm.n))
	case args.Shuffle != nil:
		return s.int63ForShuffle(args.Shuffle.n, args.Shuffle.swap, s.b.Shuffle)
	case args.Float32 != nil:
		return s.int63ForFloat32(s.b.Float32())
	case args.Float64 != nil:
		return s.int63ForFloat64(s.b.Float64())
	case args.Uint32 != nil:
		return s.int63ForUint32(s.b.Uint32())
	case args.Intn != nil:
		return s.int63ForIntn(args.Intn.n, s.b.Intn(args.Intn.n))
	case args.Int31n != nil:
		return s.int63ForInt31n(args.Int31n.n, s.b.Int31n(args.Int31n.n))
	case args.Int63n != nil:
		return s.int63ForInt63n(args.Int63n.n, s.b.Int63n(args.Int63n.n))
	case args.Int31 != nil:
		return s.int63ForInt31(s.b.Int31())
	case args.Int63 != nil:
		return s.int63ForInt63(s.b.Int63())
	}

	return s.def.Int63()
}

// Uint64 implements rand.Source64.
func (s *Source) Uint64() uint64 {
	args := s.l()

	switch {
	case args.Uint64 != nil:
		return s.uint64ForUint64(s.b.Uint64())
	}
	return s.def.Uint64()
}

// Seed implements rand.Source.
func (s *Source) Seed(seed int64) {
	s.def.Seed(seed)
}

func (s *Source) int63ForPerm(n int, res []int) int64 {
	if len(s.permrets) > 0 {
		return pop(&s.permrets)
	}

	s.permrets = make([]int64, n)
	m := make([]int, n)
	copy(m, res)
	for i := n - 1; i >= 0; i-- {
		j := 0
		for m[j] != i {
			j++
		}
		s.permrets[i] = s.int63ForIntn(i+1, j)
		m[i], m[j] = m[j], m[i]
	}

	return pop(&s.permrets)
}

func (s *Source) int63ForShuffle(n int, swap func(int, int), behavior func(int, func(int, int))) int64 {
	if len(s.shufflerets) > 0 {
		p := pop(&s.shufflerets)
		return p
	}

	record := make(map[int]int)
	recorder := func(i, j int) {
		if i >= n || j >= n {
			panic("invalid argument for swap")
		}

		iv, ok := record[i]
		if !ok {
			iv = i
		}
		record[j] = iv

		jv, ok := record[j]
		if !ok {
			jv = j
		}
		record[i] = jv
	}
	behavior(n, recorder)

	for i := n - 1; i > 0; i-- {
		j := i
		iv, ok := record[i]
		if ok && iv < i {
			j = iv
		}

		if i > 1<<31-1-1 {
			s.shufflerets = append(s.shufflerets, s.int63ForInt63n(int64(i+1), int64(j)))
		} else {
			s.shufflerets = append(s.shufflerets, s.int63ForInt31nInternal(int32(i+1), int32(j)))
		}
	}

	return pop(&s.shufflerets)
}

func (s *Source) int63ForFloat32(res float32) int64 {
	return s.int63ForFloat64(float64(res))
}

func (s *Source) int63ForFloat64(res float64) int64 {
	return int64(res * (1 << 63))
}

func (s *Source) int63ForIntn(n, res int) int64 {
	if n <= 1<<31-1 {
		return s.int63ForInt31n(int32(n), int32(res))
	}
	return s.int63ForInt63n(int64(n), int64(res))
}

func (s *Source) int63ForInt63(res int64) int64 {
	return res
}

func (s *Source) int63ForInt63n(n, res int64) int64 {
	return s.int63ForInt63(res)
}

func (s *Source) int63ForInt31(res int32) int64 {
	return int64(res) << 32
}

func (s *Source) int63ForInt31n(n, res int32) int64 {
	return s.int63ForInt31(res)
}

func (s *Source) int63ForInt31nInternal(n, res int32) int64 {
	prod := uint64(res)<<32 + uint64(n)
	prod += uint64(n) - prod%uint64(n)
	return s.int63ForUint32(uint32(prod / uint64(n)))
}

func (s *Source) int63ForUint32(res uint32) int64 {
	return int64(res) << 31
}

func (s *Source) uint64ForUint64(res uint64) uint64 {
	return res
}

func pop(s *[]int64) int64 {
	ret := (*s)[0]
	*s = (*s)[1:]
	return ret
}
