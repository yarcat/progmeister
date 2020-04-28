// Package arrays contains tests for https://progmeister.ch/blog/problems/arrays/.
package arrays

import "reflect"

// https://play.golang.org/p/E_fLHCJYiKM
func isPermutation(arr interface{}) uint {
	t := reflect.TypeOf(arr)
	if t.Kind() != reflect.Array || t.Elem().Kind() != reflect.Int {
		panic("arr must be an array of integers")
	}
	v := reflect.ValueOf(arr)
	m := make([]int, v.Len())
	for i := 0; i < len(m); i++ {
		n := int(v.Index(i).Int())
		if n <= 0 || n > len(m) {
			return 0
		}
		if m[n-1] > 0 {
			return 0
		}
		m[n-1]++
	}
	return 1
}
