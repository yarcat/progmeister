package arrays

import "fmt"

// Test represents a test that can be executed.
type Test int

const (
	// MaxMin is a test for an excercise #1 from https://progmeister.ch/blog/problems/arrays/.
	MaxMin Test = iota
	// IsPermutation is a test for an excercise #7 from https://progmeister.ch/blog/problems/arrays/.
	IsPermutation
)

var tests = map[Test]func(f interface{}) (passed, failed int, err error){
	MaxMin:        testMaxMin,
	IsPermutation: testIsPermutation,
}

func (t Test) run(f interface{}) (passed, failed int, err error) {
	testFn, ok := tests[t]
	if !ok {
		return 0, 0, fmt.Errorf("function not found for Test(%v). it must be added to tests map", t)
	}
	return testFn(f)
}
