package strings

import (
	"fmt"
	"log"
)

// TestIsNumber tests isNumber function.
//
// Problems statement:
//
//	Implement a function `func isNumber(s string) bool` that returns true if s
//	consists of digits [0..9] only:
//
//		isNumber("")     -> false
//		isNumber("abc")  -> false
//		isNumber("12ab") -> false
//		isNumber("0")    -> true
//		isNumber("123")  -> true
func TestIsNumber(fn func(string) bool) {
	var pass, fail int
	for _, test := range []struct {
		s    string
		want bool
	}{
		{s: "", want: false},
		{s: "abc", want: false},
		{s: "01234five6789", want: false},
		{s: "0", want: true},
		{s: "000", want: true},
		{s: "0123456789", want: true},
		{s: "9876543210", want: true},
		{s: "4832651097", want: true},
	} {
		func() {
			defer func() {
				if err := recover(); err != nil {
					log.Printf("FAIL: isNumber(%q), want %v\n%v",
						test.s, test.want, err)
					fail++
				}
			}()
			actual := fn(test.s)
			if actual == test.want {
				log.Printf("PASS: isNumber(%q)", test.s)
				pass++
			} else {
				log.Printf("FAIL: isNumber(%q) = %v, want %v",
					test.s, actual, test.want)
				fail++
			}
		}()
	}
	fmt.Printf("isNumber() test results: pass=%d, fail=%d.\n", pass, fail)
}

// TestNormalizeNumber tests normalizeNumber function.
//
// Problem statement:
//
//	Implement a function `func normalizeNumber(n string) (norm string, num bool)`
//	that removes trailing space characters and leading '0' and space characters
//	if n is a number (as returned by isNumber() function). If n is not a number
//	then isNumber return value should be set to false:
//
//		normalizeNumber("")        -> ["" false]
//		normalizeNumber("  ")      -> ["" false]
//		normalizeNumber("  abc  ") -> ["" false]
//		normalizeNumber("  00123") -> ["" true]
func TestNormalizeNumber(fn func(string) (string, bool)) {
	var pass, fail int
	for _, test := range []struct {
		s        string
		wantnorm string
		wantnum  bool
	}{
		{s: "", wantnum: false},
		{s: "  ", wantnum: false},
		{s: " abc ", wantnum: false},
		{s: "  0001234567890  ", wantnum: true, wantnorm: "1234567890"},
		{s: "  012345678900  ", wantnum: true, wantnorm: "12345678900"},
		{s: "  000  ", wantnum: true, wantnorm: "0"},
	} {
		func() {
			defer func() {
				if err := recover(); err != nil {
					log.Printf("FAIL: normalizeNumber(%q), want [%q %v]\n%v",
						test.s, test.wantnorm, test.wantnum, err)
					fail++
				}
			}()
			actualNorm, actualNum := fn(test.s)
			if test.wantnum && actualNum && actualNum == test.wantnum {
				log.Printf("PASS: normalizeNumber(%q)", test.s)
				pass++
			} else if !test.wantnum && !actualNum {
				log.Printf("PASS: normalizeNumber(%q)", test.s)
				pass++
			} else {
				log.Printf("FAIL: normalizeNumber(%q) = [%q %v], want [%q %v]",
					test.s, actualNorm, actualNum, test.wantnorm, test.wantnum)
				fail++
			}
		}()
	}
	fmt.Printf("normalizeNumber() test results: pass=%d, fail=%d.\n", pass, fail)
}
