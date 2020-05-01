package slices

import (
	"fmt"
	"log"
)

// TestRemoveAdjacent contains tests for removeAdjacent function.
func TestRemoveAdjacent(fn func([]int) []int) {
	var pass, fail int
	for _, test := range []struct {
		name     string
		in, want []int
	}{
		{name: "nil"},
		{name: "empty", in: []int{}, want: []int{}},
		{name: "fixed", in: slice(10, fixed(123)), want: []int{123}},
		{name: "count", in: slice(10, count()), want: slice(10, count())},
		{name: "count with dups", in: slice(10, dup(3, count())), want: slice(4, count())},
		{name: "count with dups reversed", in: rev(slice(10, dup(3, count()))), want: rev(slice(4, count()))},
	} {
		actual := fn(test.in)
		if equal(actual, test.want) {
			log.Printf("PASS: %s", test.name)
			pass++
		} else {
			log.Printf("FAIL: %s: removeAdjacent(%v) = %v, want %v",
				test.name, test.in, actual, test.want)
			fail++
		}
	}
	fmt.Printf("removeAdjacent() test results: pass=%d, fail=%d.\n", pass, fail)
}
