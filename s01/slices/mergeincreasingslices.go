package slices

import (
	"fmt"
	"log"
)

// TestMergeIncreasingSlices contains tests for mergeIncreasingSlices function.
func TestMergeIncreasingSlices(fn func(a, b []int) []int) {
	var pass, fail int
	for _, test := range []struct {
		in1, in2, want []int
	}{
		{},
		{in1: []int{}, in2: []int{}},
		{in1: slice(5, count()), in2: slice(5, count()), want: slice(10, dup(2, count()))},
		{in1: slice(4, dup(2, count())), in2: slice(2, count()), want: slice(6, dup(3, count()))},
		{in2: slice(4, dup(2, count())), in1: slice(2, count()), want: slice(6, dup(3, count()))},
	} {
		actual := fn(test.in1, test.in2)
		if equal(test.want, actual) {
			log.Printf("PASS: mergeIncreasingSlices(%v, %v)", test.in1, test.in2)
			pass++
		} else {
			log.Printf("FAIL: mergeIncreasingSlices(%v, %v) = %v, want %v",
				test.in1, test.in2, actual, test.want)
			fail++
		}
	}
	fmt.Printf("mergeIncreasingSlices() test results: pass=%d, fail=%d.\n", pass, fail)
}
