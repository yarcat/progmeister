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

func insertAt(s, subs string, i int) string {
	return "your implementation goes here"
}

func ExampleTestInsertAt() {
	strings.TestInsertAt(insertAt)
	// Output:
	// InsertAt() test results: pass=0, fail=13.
}

func index(s, subs string) int {
	return 65535
}

func ExampleTestIndex() {
	strings.TestIndex(index)
	// Output:
	// index() test results: pass=0, fail=7.
}

func rightIndex(s, subs string) int {
	return 65535
}

func ExampleTestRightIndex() {
	strings.TestRightIndex(index)
	// Output:
	// rightIndex() test results: pass=0, fail=7.
}

func indices(s, subs string) []int {
	return []int{65535}
}

func ExampleTestIndices() {
	strings.TestIndices(indices)
	// Output:
	// indices() test results: pass=0, fail=8.
}

func Example() {
	strings.TestTrimLeftSpace(trimLeftSpace)
	strings.TestTrimRightSpace(trimRightSpace)
	strings.TestTrimSpace(trimSpace)
	strings.TestInsertAt(insertAt)
	strings.TestIndex(index)
	strings.TestRightIndex(rightIndex)
	strings.TestIndices(indices)
	// Output:
	// trimLeftSpace() test results: pass=0, fail=8.
	// trimRightSpace() test results: pass=0, fail=8.
	// trimSpace() test results: pass=0, fail=8.
	// InsertAt() test results: pass=0, fail=13.
	// index() test results: pass=0, fail=7.
	// rightIndex() test results: pass=0, fail=7.
	// indices() test results: pass=0, fail=8.
}
