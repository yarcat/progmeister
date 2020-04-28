package arrays

// Test represents a test that can be executed.
type Test int

const (
  // IsPermutation is a test for an excercise #7 from https://progmeister.ch/blog/problems/arrays/.
  IsPermutation Test = iota
)

var tests = map[Test]func(f interface)(passed, failed int, err error) {
  IsPermutation: testIsPermutation,
}

func (t Test) run(f interface) (passed, failed int, err error) {
  return tests[t](f)
}
