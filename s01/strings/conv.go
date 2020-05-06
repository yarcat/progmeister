package strings

import (
	"fmt"
	"log"
)

// TestItoa tests itoa function.
//
// Problem statement:
//
//	Implement a function `func itoa(n int) string` that returns a string
//	representation (base 10) of a given integer:
//
//		itoa(0)    -> "0"
//		itoa(123)  -> "123"
//		itoa(-123) -> "-123"
func TestItoa(fn func(n int) string) {
	var pass, fail int
	for _, test := range []struct {
		n    int
		want string
	}{
		{n: 0, want: "0"},
		{n: 123, want: "123"},
		{n: -123, want: "-123"},
		{n: 1234567890, want: "1234567890"},
	} {
		func() {
			defer func() {
				if err := recover(); err != nil {
					log.Printf("PANIC: itoa(%v), want %q\n%v",
						test.n, test.want, err)
					fail++
				}
			}()
			actual := fn(test.n)
			if actual == test.want {
				log.Printf("PASS: itoa(%v)", test.n)
				pass++
			} else {
				log.Printf("FAIL: itoa(%v) = %q, want %q",
					test.n, actual, test.want)
				fail++
			}
		}()
	}
	fmt.Printf("itoa() test results: pass=%d, fail=%d.\n", pass, fail)
}
