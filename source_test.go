package randtest

import (
	"fmt"
	"math"
	"math/rand"
	"testing"

	"github.com/matryer/is"
)

const MaxUint = ^uint(0)
const MinUint = 0
const MaxInt = int(MaxUint >> 1)
const MinInt = -MaxInt - 1

// TODO: test statefull behavior

func TestMaxSource_Intn(t *testing.T) {
	cases := []struct{ n, ans int }{
		{1, 0},
		{1024, 1023},
		{10000, 9999},
		{MaxInt, MaxInt - 1},
	}

	r := rand.New(NewSource(&MaxBehavior{}))
	for _, c := range cases {
		t.Run(fmt.Sprintf("n=%d", c.n), func(t *testing.T) {
			is := is.New(t)
			is.Equal(r.Intn(c.n), c.ans)
		})
	}
}

func TestMaxSource_Int63n(t *testing.T) {
	cases := []struct{ n, ans int64 }{
		{1, 0},
		{1024, 1023},
		{10000, 9999},
		{math.MaxInt64, math.MaxInt64 - 1},
	}

	r := rand.New(NewSource(&MaxBehavior{}))
	for _, c := range cases {
		t.Run(fmt.Sprintf("n=%d", c.n), func(t *testing.T) {
			is := is.New(t)
			is.Equal(r.Int63n(c.n), c.ans)
		})
	}
}

func TestMaxSource_Int31n(t *testing.T) {
	cases := []struct{ n, ans int32 }{
		{1, 0},
		{1024, 1023},
		{10000, 9999},
		{math.MaxInt32, math.MaxInt32 - 1},
	}

	r := rand.New(NewSource(&MaxBehavior{}))
	for _, c := range cases {
		t.Run(fmt.Sprintf("n=%d", c.n), func(t *testing.T) {
			is := is.New(t)
			is.Equal(r.Int31n(c.n), c.ans)
		})
	}
}

func TestMaxSource_Int(t *testing.T) {
	is := is.New(t)

	r := rand.New(NewSource(&MaxBehavior{}))
	is.Equal(r.Int(), MaxInt-1)
}

func TestMaxSource_Int63(t *testing.T) {
	is := is.New(t)

	r := rand.New(NewSource(&MaxBehavior{}))
	is.Equal(r.Int63(), int64(math.MaxInt64-1))
}

func TestMaxSource_Int31(t *testing.T) {
	is := is.New(t)

	r := rand.New(NewSource(&MaxBehavior{}))
	is.Equal(r.Int31(), int32(math.MaxInt32-1))
}

func TestMaxSource_Uint64(t *testing.T) {
	is := is.New(t)

	r := rand.New(NewSource(&MaxBehavior{}))
	is.Equal(r.Uint64(), uint64(math.MaxUint64))
}

func TestMaxSource_Uint32(t *testing.T) {
	is := is.New(t)

	r := rand.New(NewSource(&MaxBehavior{}))
	is.Equal(r.Uint32(), uint32(math.MaxUint32))
}

func TestMaxSource_Float32(t *testing.T) {
	is := is.New(t)

	r := rand.New(NewSource(&MaxBehavior{}))
	is.Equal(r.Float32(), float32(0.99999994))
}

func TestMaxSource_Float64(t *testing.T) {
	is := is.New(t)

	r := rand.New(NewSource(&MaxBehavior{}))
	is.Equal(r.Float64(), float64(0.9999999999999999))
}

func TestMinSource_Intn(t *testing.T) {
	is := is.New(t)

	r := rand.New(NewSource(&MinBehavior{}))
	is.Equal(r.Intn(10), 0)
}

func TestMinSource_Int31n(t *testing.T) {
	is := is.New(t)

	r := rand.New(NewSource(&MinBehavior{}))
	is.Equal(r.Int31n(10), int32(0))
}

func TestMinSource_Int63n(t *testing.T) {
	is := is.New(t)

	r := rand.New(NewSource(&MinBehavior{}))
	is.Equal(r.Int63n(10), int64(0))
}

func TestMinSource_Int31(t *testing.T) {
	is := is.New(t)

	r := rand.New(NewSource(&MinBehavior{}))
	is.Equal(r.Int31(), int32(0))
}

func TestMinSource_Int63(t *testing.T) {
	is := is.New(t)

	r := rand.New(NewSource(&MinBehavior{}))
	is.Equal(r.Int63(), int64(0))
}

func TestMinSource_Uint32(t *testing.T) {
	is := is.New(t)

	r := rand.New(NewSource(&MinBehavior{}))
	is.Equal(r.Uint32(), uint32(0))
}

func TestMinSource_Uint64(t *testing.T) {
	is := is.New(t)

	r := rand.New(NewSource(&MinBehavior{}))
	is.Equal(r.Uint64(), uint64(0))
}

func TestMinSource_Float32(t *testing.T) {
	is := is.New(t)

	r := rand.New(NewSource(&MinBehavior{}))
	is.Equal(r.Float32(), float32(0))
}

func TestMinSource_Float64(t *testing.T) {
	is := is.New(t)

	r := rand.New(NewSource(&MinBehavior{}))
	is.Equal(r.Float64(), float64(0))
}

type permBehavior func(int) []int

func (b permBehavior) Perm(n int) []int {
	return b(n)
}

func TestPerm(t *testing.T) {
	invert := func(n int) []int {
		m := make([]int, 0, n)
		for i := n - 1; i >= 0; i-- {
			m = append(m, i)
		}
		return m
	}
	cases := []struct {
		b   func(int) []int
		n   int
		ans []int
	}{
		{
			b:   invert,
			n:   1,
			ans: []int{0},
		},
		{
			b:   invert,
			n:   10,
			ans: []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0},
		},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("n=%d", c.n), func(t *testing.T) {
			is := is.New(t)

			r := rand.New(NewSource(permBehavior(c.b)))
			is.Equal(r.Perm(c.n), c.ans)
		})
	}
}

type shuffleBehavior func(int, func(int, int))

func (b shuffleBehavior) Shuffle(n int, swap func(int, int)) {
	b(n, swap)
}

func TestMaxSource_Shuffle(t *testing.T) {
	invert := func(n int, swap func(int, int)) {
		for i := 0; i < n/2; i++ {
			swap(i, n-i-1)
		}
	}
	cases := []struct {
		b   func(int, func(int, int))
		n   int
		ans []int
	}{
		{
			b:   invert,
			n:   1,
			ans: []int{0},
		},
		{
			b:   invert,
			n:   10,
			ans: []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0},
		},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("n=%d", c.n), func(t *testing.T) {
			is := is.New(t)

			r := rand.New(NewSource(shuffleBehavior(invert)))
			a := make([]int, 0, c.n)
			for i := 0; i < c.n; i++ {
				a = append(a, i)
			}

			r.Shuffle(c.n, func(i, j int) { a[i], a[j] = a[j], a[i] })
			is.Equal(a, c.ans)
		})
	}
}
