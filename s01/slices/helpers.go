package slices

import (
	"math/rand"
	"reflect"
)

func fixed(n int) func() int {
	return func() int { return n }
}

func count() func() int {
	x := -1
	return func() int {
		x++
		return x
	}
}

func rev(s []int) []int {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

func dup(n int, fn func() int) func() int {
	var x, cnt int
	return func() int {
		if cnt%n == 0 {
			x = fn()
		}
		cnt++
		return x
	}
}

func even() int {
	n := rand.Int() % 100
	if n&1 == 1 {
		return n - 1
	}
	return n
}

func odd() int {
	n := rand.Int() % 100
	if n&1 == 1 {
		return n
	}
	return n - 1
}

func slice(n int, fn func() int) (s []int) {
	if n <= 0 {
		return
	}
	s = make([]int, n)
	for i := 0; i < n; i++ {
		s[i] = fn()
	}
	return
}

func insert(a, b []int) (s []int) {
	if len(a) == 0 && len(b) == 0 {
		return
	}
	s = make([]int, 0, len(a)+len(b))
	var ia, ib int
	for ia < len(a) && ib < len(b) {
		if rand.Int()%2 == 0 {
			s = append(s, a[ia])
			ia++
		} else {
			s = append(s, b[ib])
			ib++
		}
	}
	if ia < len(a) {
		s = append(s, a[ia:]...)
	}
	if ib < len(b) {
		s = append(s, a[ib:]...)
	}
	return
}

func equal(a, b []int) bool {
	if len(a) == 0 && len(b) == 0 {
		return true
	}
	return reflect.DeepEqual(a, b)
}
