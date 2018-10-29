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

func TestIntn(t *testing.T) {
	cases := []struct{ n, ans int }{
		{1, 0},
		{1024, 1023},
		{10000, 9999},
		{MaxInt, MaxInt - 1},
	}

	r := rand.New(NewSource())
	for _, c := range cases {
		t.Run(fmt.Sprintf("n=%d", c.n), func(t *testing.T) {
			is := is.New(t)
			is.Equal(r.Intn(c.n), c.ans)
		})
	}
}

func TestInt63n(t *testing.T) {
	cases := []struct{ n, ans int64 }{
		{1, 0},
		{1024, 1023},
		{10000, 9999},
		{math.MaxInt64, math.MaxInt64 - 1},
	}

	r := rand.New(NewSource())
	for _, c := range cases {
		t.Run(fmt.Sprintf("n=%d", c.n), func(t *testing.T) {
			is := is.New(t)
			is.Equal(r.Int63n(c.n), c.ans)
		})
	}
}

func TestInt31n(t *testing.T) {
	cases := []struct{ n, ans int32 }{
		{1, 0},
		{1024, 1023},
		{10000, 9999},
		{math.MaxInt32, math.MaxInt32 - 1},
	}

	r := rand.New(NewSource())
	for _, c := range cases {
		t.Run(fmt.Sprintf("n=%d", c.n), func(t *testing.T) {
			is := is.New(t)
			is.Equal(r.Int31n(c.n), c.ans)
		})
	}
}

func TestInt(t *testing.T) {
	is := is.New(t)

	r := rand.New(NewSource())
	is.Equal(r.Int(), MaxInt-1)
}

func TestInt63(t *testing.T) {
	is := is.New(t)

	r := rand.New(NewSource())
	is.Equal(r.Int63(), int64(math.MaxInt64-1))
}

func TestInt31(t *testing.T) {
	is := is.New(t)

	r := rand.New(NewSource())
	is.Equal(r.Int31(), int32(math.MaxInt32-1))
}

func TestShuffle(t *testing.T) {
	is := is.New(t)

	s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	r := rand.New(NewSource())
	r.Shuffle(len(s), func(i, j int) {
		s[i], s[j] = s[j], s[i]
	})
	is.Equal(s, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})
}
