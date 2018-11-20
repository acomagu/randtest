package randtest_test

import (
	"fmt"
	"math/rand"

	"github.com/acomagu/randtest"
)

type Behavior struct{}

func (b *Behavior) Intn(n int) int {
	return n / 2
}

func Example() {
	source := randtest.NewSource(&Behavior{})
	r := rand.New(source)
	fmt.Println(r.Intn(10))
	// Output: 5
}
