package slices

import (
	"fmt"
	"log"
)

// TestMergeIncreasingSlices implements tests for mergeIncreasingSlices function.
//
// Problem statement:
//
//	Implement a function func `mergeIncreasingSlices(a, b []int) []int`, which is
//	given two slices containing integers in an increasing order as the input. The
//	function must return a slices that contains integers from the input array in
//	the increasing order.
//
// Example:
//
//	mergeIncreasingSlices([]int{1, 5, 5, 7}, []int{1, 2, 4, 6, 9}) -> []int{1, 1, 2, 4, 5, 5, 6, 7, 9}
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
		func() {
			defer func() {
				if err := recover(); err != nil {
					log.Printf("PANIC: mergeIncreasingSlices(%v, %v) want %v\n%v",
						test.in1, test.in2, test.want, err)
				}
			}()
			actual := fn(test.in1, test.in2)
			if equal(test.want, actual) {
				log.Printf("PASS: mergeIncreasingSlices(%v, %v)", test.in1, test.in2)
				pass++
			} else {
				log.Printf("FAIL: mergeIncreasingSlices(%v, %v) = %v, want %v",
					test.in1, test.in2, actual, test.want)
				fail++
			}
		}()
	}
	fmt.Printf("mergeIncreasingSlices() test results: pass=%d, fail=%d.\n", pass, fail)
}
