# Crazy randtest

[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](https://godoc.org/github.com/acomagu/randtest) [![CircleCI](https://img.shields.io/circleci/project/github/acomagu/randtest.svg?style=flat-square)](https://circleci.com/gh/acomagu/randtest) [![Codecov](https://img.shields.io/codecov/c/github/acomagu/randtest.svg?style=flat-square)](https://codecov.io/gh/acomagu/randtest)


randtest controls the math/rand.Rand behavior from the Source for testing purpose.

```Go
type Behavior struct{}

func (b *Behavior) Intn(n int) int {
	return n / 2
}

func main() {
	source := randtest.NewSource(&Behavior{})
	r := rand.New(source)

	fmt.Println(r.Intn(10)) // Always 5.
}
```

This is experimental package. It can be broken by updating Go because the logic is highly depends on undocumented behavior of math/rand package. Only tested on Go1.11.
