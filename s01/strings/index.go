package strings

import (
	"fmt"
	"log"
)

type indexTest struct {
	s, subs string
	want    []int
}

var indexTests = []indexTest{
	{s: "", subs: "", want: []int{0}},
	{s: "abc", subs: "", want: []int{0}},
	{s: "abc", subs: "abc", want: []int{0}},
	{s: "abc", subs: "bc", want: []int{1}},
	{s: "abcbc", subs: "bc", want: []int{1}},
	{s: "abc", subs: "cba", want: nil},
	{s: "", subs: "abc", want: nil},
}

// TestIndex tests index function.
//
// Problem statement:
//
//	Implement a function `func index(s, subs string) int` that returns the index
//	of the left-most instance of subs in s, or -1 if subs does not present in s:
//
//		index("foo", "oo") -> 1
//		index("foo", "") -> 0
//		index("foo", "bar") -> -1
func TestIndex(fn func(string, string) int) {
	testIndices("index", func(s, subs string) []int {
		i := fn(s, subs)
		if i < 0 {
			return nil
		}
		return []int{i}
	}, indexTests)
}

var rightIndexTests = []indexTest{
	{s: "", subs: "", want: []int{0}},
	{s: "abc", subs: "", want: []int{3}},
	{s: "abc", subs: "abc", want: []int{0}},
	{s: "abc", subs: "bc", want: []int{1}},
	{s: "abcbc", subs: "bc", want: []int{3}},
	{s: "abc", subs: "cba", want: nil},
	{s: "", subs: "abc", want: nil},
}

// TestRightIndex tests rightIndex function.
//
// Problem statement:
//
//	Implement a function `func rightIndex(s, subs string) int` that returns the
//	index of the right-most instance of subs in s, or -1 if subs does not present
//	in s:
//
//		rightIndex("foofoo", "oo") -> 4
//		rightIndex("foo", "") -> 3
//		rightIndex("foo", "bar") -> -1
func TestRightIndex(fn func(string, string) int) {
	testIndices("rightIndex", func(s, subs string) []int {
		i := fn(s, subs)
		if i < 0 {
			return nil
		}
		return []int{i}
	}, rightIndexTests)
}

var indicesTests = []indexTest{
	{s: "", subs: "", want: nil},
	{s: "abc", subs: "", want: nil},
	{s: "abc", subs: "abc", want: []int{0}},
	{s: "abc", subs: "bc", want: []int{1}},
	{s: "abcbc", subs: "bc", want: []int{1, 3}},
	{s: "abc", subs: "cba", want: nil},
	{s: "", subs: "abc", want: nil},
	{s: "aaaa", subs: "aa", want: []int{0, 1, 2}},
}

// TestIndices tests indices function.
//
// Problem statement:
//
//	Implement a function `func indices(s, subs string) []int` that returns all
//	indices of subs in s with an overlap (it means that "oo" in "ooo" will be
//	matched twice - at positions 0 and 1):
//
//		rightIndex("fooofooo", "oo") -> [1 2 5 6]
//		rightIndex("foo", "") -> []
//		rightIndex("foo", "bar") -> []
func TestIndices(fn func(string, string) []int) {
	testIndices("indices", fn, indicesTests)
}

func testIndices(name string, fn func(string, string) []int, tests []indexTest) {
	var pass, fail int
	for _, test := range tests {
		func() {
			defer func() {
				if err := recover(); err != nil {
					log.Printf("PANIC: %s(%q, %q) want %v\n%v",
						name, test.s, test.subs, test.want, err)
					fail++
				}
			}()
			actual := fn(test.s, test.subs)
			if equal(actual, test.want) {
				log.Printf("PASS: %s(%q, %q)", name, test.s, test.subs)
				pass++
			} else {
				log.Printf("FAIL: %s(%q, %q) = %v, want %v",
					name, test.s, test.subs, actual, test.want)
				fail++
			}
		}()
	}
	fmt.Printf("%s() test results: pass=%d, fail=%d.\n", name, pass, fail)
}

func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
