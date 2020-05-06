package strings

import (
	"fmt"
	"log"
)

// TestInsertAt tests insertAt function.
//
// Problem statement:
//
//	Implement a function `func insertAt(s, subs string, i int) string` that
//	inserts a substring into a string after an index provided. If the initial
//	string is too short, space character should be used to fill the necessary
//	space. Negative index is valid and means that the insertion index starts
//	from the end of the string in negative direction. Example:
//
//		insertAt("bar", "foo", 0) -> "foobar"
//		insertAt("bar", "foo", 1) -> "bfooar"
//		insertAt("bar", "foo", 3) -> "barfoo"
//		insertAt("bar", "foo", -1) -> "barfoo"
//		insertAt("bar", "foo", 4) -> "bar foo"
//		insertAt("", "foo", 4) -> "    foo"
func TestInsertAt(fn func(string, string, int) string) {
	var pass, fail int
	for _, test := range []struct {
		s, subs, want string
		i             int
	}{
		{s: "bar", subs: "foo", i: 0, want: "foobar"},
		{s: "bar", subs: "foo", i: 1, want: "bfooar"},
		{s: "bar", subs: "foo", i: 3, want: "barfoo"},
		{s: "bar", subs: "foo", i: -1, want: "barfoo"},
		{s: "bar", subs: "foo", i: 4, want: "bar foo"},
		{s: "", subs: "foo", i: 4, want: "    foo"},
		{s: "", subs: "foo", i: 0, want: "foo"},
		{s: "", subs: "foo", i: -1, want: "foo"},
		{s: "", subs: "foo", i: -4, want: "foo   "},
		{s: "", subs: "", want: ""},
		{s: "", subs: "", i: 5, want: "     "},
		{s: "", subs: "", i: -5, want: "    "},
		{s: "яро", subs: "слав", i: 3, want: "ярослав"},
	} {
		func() {
			defer func() {
				if err := recover(); err != nil {
					log.Printf("PANIC: InsertAt(%q, %q, %d), want %q\n%v",
						test.s, test.subs, test.i, test.want, err)
					fail++
				}
			}()
			actual := fn(test.s, test.subs, test.i)
			if actual == test.want {
				log.Printf("PASS: InsertAt(%q, %q, %d)",
					test.s, test.subs, test.i)
				pass++
			} else {
				log.Printf("FAIL: InsertAt(%q, %q, %d) = %q, want %q",
					test.s, test.subs, test.i, actual, test.want)
				fail++
			}
		}()
	}
	fmt.Printf("InsertAt() test results: pass=%d, fail=%d.\n", pass, fail)
}
