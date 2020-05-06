package strings_test

import "github.com/yarcat/progmeister/s01/strings"

func trimLeftSpace(s string) string {
	return "your implementation goes here"
}

func ExampleTestTrimLeftSpace() {
	strings.TestTrimLeftSpace(trimLeftSpace)
	// Output:
	// trimLeftSpace() test results: pass=0, fail=8.
}

func trimRightSpace(s string) string {
	return "your implementation goes here"
}

func ExampleTestTrimRightSpace() {
	strings.TestTrimRightSpace(trimRightSpace)
	// Output:
	// trimRightSpace() test results: pass=0, fail=8.
}

func trimSpace(s string) string {
	return "your implementation goes here"
}

func ExampleTestTrimSpace() {
	strings.TestTrimSpace(trimSpace)
	// Output:
	// trimSpace() test results: pass=0, fail=8.
}

func Example() {
	strings.TestTrimLeftSpace(trimLeftSpace)
	strings.TestTrimRightSpace(trimRightSpace)
	strings.TestTrimSpace(trimSpace)
	// Output:
	// trimLeftSpace() test results: pass=0, fail=8.
	// trimRightSpace() test results: pass=0, fail=8.
	// trimSpace() test results: pass=0, fail=8.
}
