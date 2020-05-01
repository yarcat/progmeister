package slices

import (
	"fmt"
	"log"
)

// TestGenerateFibonacci implements tests for `generateFibonacci` function.
//
//	Problem statement:
//		Implement a function `func generateFibonacci(n int) []int` that returns a slice
//		containing the first `n` elements of the Fibonacci sequence.
//
//	Fibonacci sequence:
//
//		https://en.wikipedia.org/wiki/Fibonacci_number
//
//		The Fibonacci sequence is a sequence of integers, where the next number is
//		calculated as the sum of two previous numbers, i.e. Fn=Fn−2+Fn−1. Assume
//		that F1=0 and F2=1.
//
//	Example:
//
//		generateFibonacci(10) -> []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34}
func TestGenerateFibonacci(fn func(int) []int) {
	var pass, fail int
	for _, test := range []struct {
		in   int
		want []int
	}{
		{in: -100},
		{in: 0},
		{in: 1, want: []int{0}},
		{in: 2, want: []int{0, 1}},
		{in: 3, want: []int{0, 1, 1}},
		{in: 4, want: []int{0, 1, 1, 2}},
		{in: 5, want: []int{0, 1, 1, 2, 3}},
		{in: 6, want: []int{0, 1, 1, 2, 3, 5}},
		{in: 7, want: []int{0, 1, 1, 2, 3, 5, 8}},
		{in: 8, want: []int{0, 1, 1, 2, 3, 5, 8, 13}},
		{in: 9, want: []int{0, 1, 1, 2, 3, 5, 8, 13, 21}},
		{in: 10, want: []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34}},
		{in: 11, want: []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55}},
		{in: 12, want: []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89}},
		{in: 13, want: []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144}},
	} {
		actual := fn(test.in)
		if equal(actual, test.want) {
			log.Printf("PASS: generateFibonacci(%v)", test.in)
			pass++
		} else {
			log.Printf("FAIL: generateFibonacci(%v) = %v, want %v",
				test.in, actual, test.want)
			fail++
		}
	}
	fmt.Printf("generateFibonacci() test results: pass=%d, fail=%d.\n", pass, fail)
}
