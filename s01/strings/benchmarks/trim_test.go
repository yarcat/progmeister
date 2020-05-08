/*
go test -bench=. trim_test.go
goos: linux
goarch: amd64
BenchmarkTrimLeft/1-4                        121          10819238 ns/op
BenchmarkTrimLeft/2-4                        310           4163989 ns/op
BenchmarkTrimLeft/strings-4                   31          33865116 ns/op
BenchmarkTrimRight/1-4                       288           3954300 ns/op
BenchmarkTrimRight/2-4                       289           3945484 ns/op
BenchmarkTrimRight/strings-4                  19          61983880 ns/op
BenchmarkTrim/1-4                            152           7667442 ns/op
BenchmarkTrim/2-4                            124           9648740 ns/op
BenchmarkTrim/3-4                            156           7912740 ns/op
BenchmarkTrim/strings-4                       75          14156754 ns/op
PASS
ok      command-line-arguments  17.048s
*/
package trim_test

import (
	"strings"
	"testing"
)

var (
	input = strings.Repeat(" ", 10_000_000) + "a" + strings.Repeat(" ", 10_000_000)
	// Using this as a side-effect to guarantee Go doesn't optimize benchmarks too much.
	ss string
)

func trimLeft1(s string) string {
	for len(s) > 0 && s[0] == ' ' {
		s = s[1:]
	}
	return s
}

func trimLeft2(s string) string {
	i := 0
	for i < len(s) && s[i] == ' ' {
		i++
	}
	return s[i:]
}

func trimRight1(s string) string {
	for len(s) > 0 && s[len(s)-1] == ' ' {
		s = s[:len(s)-1]
	}
	return s
}

func trimRight2(s string) string {
	i := len(s) - 1
	for i >= 0 && s[i] == ' ' {
		i--
	}
	return s[:i+1]
}

func trim1(s string) string {
	return trimRight1(trimLeft2(s))
}

func trim2(s string) string {
	i := 0
	for i < len(s) && s[i] == ' ' {
		i++
	}
	j := len(s) - 1
	for j > i && s[j] == ' ' {
		j--
	}
	return s[i : j+1]
}

func trim3(s string) string {
	for len(s) > 0 && s[len(s)-1] == ' ' {
		s = s[:len(s)-1]
	}
	i := 0
	for i < len(s) && s[i] == ' ' {
		i++
	}
	return s[i:]
}

func TestTrimLeft(t *testing.T) {
	test := func(f func(string) string) func(*testing.T) {
		return func(t *testing.T) {
			actual := f(input)
			if actual != "a"+strings.Repeat(" ", 10_000_000) {
				t.Errorf("got %v, want 0", len(actual))
			}
		}
	}
	t.Run("1", test(trimLeft1))
	t.Run("2", test(trimLeft2))
	t.Run("strings", test(func(s string) string { return strings.TrimLeft(s, " ") }))
}

func BenchmarkTrimLeft(b *testing.B) {
	test := func(f func(string) string) func(*testing.B) {
		return func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				ss = f(input)
			}
		}
	}
	b.Run("1", test(trimLeft1))
	b.Run("2", test(trimLeft2))
	b.Run("strings", test(func(s string) string { return strings.TrimLeft(s, " ") }))
}

func TestTrimRight(t *testing.T) {
	test := func(f func(string) string) func(*testing.T) {
		return func(t *testing.T) {
			actual := f(input)
			if actual != strings.Repeat(" ", 10_000_000)+"a" {
				t.Errorf("got %v, want 0", len(actual))
			}
		}
	}
	t.Run("1", test(trimRight1))
	t.Run("2", test(trimRight2))
	t.Run("strings", test(func(s string) string { return strings.TrimRight(s, " ") }))
}

func BenchmarkTrimRight(b *testing.B) {
	test := func(f func(string) string) func(*testing.B) {
		return func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				ss = f(input)
			}
		}
	}
	b.Run("1", test(trimRight1))
	b.Run("2", test(trimRight2))
	b.Run("strings", test(func(s string) string { return strings.TrimRight(s, " ") }))
}

func TestTrim(t *testing.T) {
	test := func(f func(string) string) func(*testing.T) {
		return func(t *testing.T) {
			actual := f(input)
			if actual != "a" {
				t.Errorf("got %v, want 0", len(actual))
			}
		}
	}
	t.Run("1", test(trim1))
	t.Run("2", test(trim2))
	t.Run("3", test(trim3))
	t.Run("strings", test(strings.TrimSpace))
}

func BenchmarkTrim(b *testing.B) {
	test := func(f func(string) string) func(*testing.B) {
		return func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				ss = f(input)
			}
		}
	}
	b.Run("1", test(trim1))
	b.Run("2", test(trim2))
	b.Run("3", test(trim3))
	b.Run("strings", test(strings.TrimSpace))
}
