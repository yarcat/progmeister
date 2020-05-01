package slices

import (
	"fmt"
	"log"
)

// TestJosephus contains tests for josephus function.
func TestJosephus(fn func(soliders, skip int) int) {
	var pass, fail int
	for _, test := range []struct {
		soldiers, skip int
		want           int
	}{
		{soldiers: 0, skip: 0, want: 0},
		{soldiers: -1, skip: 0, want: 0},
		{soldiers: 1, skip: -1, want: 0},
		{soldiers: 1, skip: 0, want: 1},
		{soldiers: 13, skip: 0, want: 11},
		{soldiers: 41, skip: 1, want: 31},
		{soldiers: 5, skip: 1, want: 4},
	} {
		actual := fn(test.soldiers, test.skip)
		if actual == test.want {
			log.Printf("PASS: josephus(%v, %v)", test.soldiers, test.skip)
			pass++
		} else {
			log.Printf("FAIL: josephus(%v, %v) = %v, want %v",
				test.soldiers, test.skip, actual, test.want)
			fail++
		}
	}
	fmt.Printf("josephus() test results: pass=%d, fail=%d.\n", pass, fail)
}
