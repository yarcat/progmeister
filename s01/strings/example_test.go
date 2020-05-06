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

func isNumber(s string) bool {
	return false
}

func ExampleTestIsNumber() {
	strings.TestIsNumber(isNumber)
	// Output:
	// isNumber() test results: pass=3, fail=5.
}

func normalizeNumber(s string) (norm string, num bool) {
	return "your code goes here", false
}

func ExampleTestNormalizeNumber() {
	strings.TestNormalizeNumber(normalizeNumber)
	// Output:
	// normalizeNumber() test results: pass=3, fail=3.
}

func itoa(n int) string {
	return "your code goes here"
}

func ExampleTestItoa() {
	strings.TestItoa(itoa)
	// Output:
	// itoa() test results: pass=0, fail=4.
}
func Example() {
	strings.TestTrimLeftSpace(trimLeftSpace)
	strings.TestTrimRightSpace(trimRightSpace)
	strings.TestTrimSpace(trimSpace)
	strings.TestInsertAt(insertAt)
	strings.TestIndex(index)
	strings.TestRightIndex(rightIndex)
	strings.TestIndices(indices)
	strings.TestIsNumber(isNumber)
	strings.TestNormalizeNumber(normalizeNumber)
	strings.TestItoa(itoa)
	// Output:
	// trimLeftSpace() test results: pass=0, fail=8.
	// trimRightSpace() test results: pass=0, fail=8.
	// trimSpace() test results: pass=0, fail=8.
	// InsertAt() test results: pass=0, fail=13.
	// index() test results: pass=0, fail=7.
	// rightIndex() test results: pass=0, fail=7.
	// indices() test results: pass=0, fail=8.
	// isNumber() test results: pass=3, fail=5.
	// normalizeNumber() test results: pass=3, fail=3.
	// itoa() test results: pass=0, fail=4.
}
