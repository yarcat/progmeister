package slices

import (
	"fmt"
	"log"
)

// TestRemoveEven implements tests for removeEven function.
//
//	Problem statement:
//		Implement a function `func removeEven(s []int) []int` that removes even numbers
//		from the input slice, and returns a new slice with the remaining numbers in the
//		same order.
//
//	Example:
//		removeEven([]int{5, 2, 3, 6, 2, 3}) -> []int{5, 3, 3}
func TestRemoveEven(fn func([]int) []int) {
	evenInts, oddInts := slice(10, even), slice(10, odd)
	var pass, fail int
	for _, test := range []struct {
		name     string
		in, want []int
	}{
		{name: "nil"},
		{name: "empty", in: []int{}, want: []int{}},
		{name: "same even", in: slice(10, fixed(246)), want: []int{}},
		{name: "same odd", in: slice(10, fixed(123)), want: slice(10, fixed(123))},
		{name: "mix", in: insert(evenInts, oddInts), want: oddInts},
	} {
		actual := fn(test.in)
		if equal(actual, test.want) {
			log.Printf("PASS: %s", test.name)
			pass++
		} else {
			log.Printf("FAIL: %s: removeEven(%v) = %v, want %v",
				test.name, test.in, actual, test.want)
			fail++
		}
	}
	fmt.Printf("removeEven() test results: pass=%d, fail=%d.\n", pass, fail)
}
