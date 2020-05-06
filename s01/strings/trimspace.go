package strings

import (
	"fmt"
	"log"
	"strings"
)

var trimLeftSpaceTests = []trimTest{
	{s: "", want: ""},
	{s: " ", want: ""},
	{s: strings.Repeat(" ", 10), want: ""},
	{s: "spaceless", want: "spaceless"},
	{s: "middle space", want: "middle space"},
	{s: " left space", want: "left space"},
	{s: "right space ", want: "right space "},
	{s: "  s p  ac e  s   ", want: "s p  ac e  s   "},
}

// TestTrimLeftSpace tests trimLeftSpace function.
//
// Problem statement:
//
//	Implement a function `func trimLeftSpace` that returns a slice of the string s
//	with all leading space characters [1] removed:
//
//		trimLeftSpace("gopher") -> "gopher"
//		trimLeftSpace("  go goper") -> "go gopher"
//
//	[1]: https://en.wikipedia.org/wiki/Space_(punctuation)
func TestTrimLeftSpace(fn func(string) string) {
	testTrimFunction("trimLeftSpace", fn, trimLeftSpaceTests)
}

var trimRightSpaceTests = []trimTest{
	{s: "", want: ""},
	{s: " ", want: ""},
	{s: strings.Repeat(" ", 10), want: ""},
	{s: "spaceless", want: "spaceless"},
	{s: "middle space", want: "middle space"},
	{s: " left space", want: "left space"},
	{s: "right space ", want: "right space "},
	{s: "  s p  ac e  s   ", want: "  s p  ac e  s"},
}

// TestTrimRightSpace tests trimRightSpace function.
//
// Problem statement:
//
//	Implement a function `func trimRightSpace` that returns a slice of the string s
//	with all trailing space characters [1] removed:
//
//		trimRightSpace("gopher") -> "gopher"
//		trimRightSpace("go goper  ") -> "go gopher"
//
//	[1]: https://en.wikipedia.org/wiki/Space_(punctuation)
func TestTrimRightSpace(fn func(string) string) {
	testTrimFunction("trimRightSpace", fn, trimRightSpaceTests)
}

var trimSpaceTests = []trimTest{
	{s: "", want: ""},
	{s: " ", want: ""},
	{s: strings.Repeat(" ", 10), want: ""},
	{s: "spaceless", want: "spaceless"},
	{s: "middle space", want: "middle space"},
	{s: " left space", want: "left space"},
	{s: "right space ", want: "right space"},
	{s: "  s p  ac e  s   ", want: "s p  ac e  s"},
}

// TestTrimSpace tests trimSpace function.
//
// Problem statement:
//
//	Implement a function `func trimRightSpace` that returns a slice of the string s
//	with all leading and trailing space characters [1] removed:
//
//		trimSpace("  we are surrounded  ") -> "we are surrounded"
//
//	[1]: https://en.wikipedia.org/wiki/Space_(punctuation)
func TestTrimSpace(fn func(string) string) {
	testTrimFunction("trimSpace", fn, trimSpaceTests)
}

type trimTest struct{ s, want string }

func testTrimFunction(name string, fn func(string) string, tests []trimTest) {
	var pass, fail int
	for _, test := range tests {
		func() {
			defer func() {
				if err := recover(); err != nil {
					log.Printf("PANIC: %s(%q), want %q\n%v",
						name, test.s, test.want, err)
				}
			}()
			actual := fn(test.s)
			if actual == test.want {
				log.Printf("PASS: %s(%q)", name, test.s)
				pass++
			} else {
				log.Printf("FAIL: %s(%q) = %q, want %q",
					name, test.s, actual, test.want)
				fail++
			}
		}()
	}
	fmt.Printf("%s() test results: pass=%d, fail=%d.\n", name, pass, fail)
}
