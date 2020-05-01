package slices

import (
	"fmt"
	"log"
)

// TestJosephus implements tests for josephus function.
//
// Problem statement:
//
//	Implement a function `func josephus(soldiers, skip int) int` that solves the
//	Josephus problem. Given the number of soldiers and number to be skipped, the
//	functions returns the position counted from the selected first soldier in the
//	initial circle to avoid execution. Positions start from 1.
//
// Josephus Problem:
//
//	https://en.wikipedia.org/wiki/Josephus_problem
//
//	Soldiers are standing in a circle waiting to be executed. Counting begins from
//	the first soldier in the circle and proceeds around the circle in a specified
//	direction. After a specified number of soldiers are skipped, the next soldier
//	is executed. The procedure is repeated with the remaining soldiers, starting
//	with the next soldier, going in the same direction and skipping the same number
//	of soldiers, until only one soldier remains, and is freed.
//
// Examples:
//
//	1) josephus(13, 0) -> 11
//	2) josephus(41, 1) -> 30
//	3) josephus(5, 1) -> 13
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
