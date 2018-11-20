package randtest_test

import (
	"fmt"
	"github.com/acomagu/randtest"
	"math/rand"
)

func ExampleMaxBehavior() {
	source := randtest.NewSource(&randtest.MaxBehavior{})
	r := rand.New(source)

	fmt.Println(r.Intn(10))
	fmt.Println(r.Int31n(10))
	fmt.Println(r.Int63n(10))
	fmt.Println(r.Int31())
	fmt.Println(r.Int63())
	fmt.Println(r.Uint32())
	fmt.Println(r.Uint64())
	fmt.Println(r.Float32())
	fmt.Println(r.Float64())
	// Output:
	// 9
	// 9
	// 9
	// 2147483646
	// 9223372036854775806
	// 4294967295
	// 18446744073709551615
	// 0.99999994
	// 0.9999999999999999
}

func ExampleMinBehavior() {
	source := randtest.NewSource(&randtest.MinBehavior{})
	r := rand.New(source)

	fmt.Println(r.Intn(10))
	fmt.Println(r.Int31n(10))
	fmt.Println(r.Int63n(10))
	fmt.Println(r.Int31())
	fmt.Println(r.Int63())
	fmt.Println(r.Uint32())
	fmt.Println(r.Uint64())
	fmt.Println(r.Float32())
	fmt.Println(r.Float64())
	// Output:
	// 0
	// 0
	// 0
	// 0
	// 0
	// 0
	// 0
	// 0
	// 0
}
