// Package arrays contains tests for https://progmeister.ch/blog/problems/arrays/.
package arrays

import (
	"errors"
	"reflect"
)

func testIsPermutation(f interface{}) (passed, failed int, err error) {
	arrLen := 0
	t := reflect.TypeOf(f)
	bad := true
	for {
		if t.Kind != reflect.Func || t.NumIn() != 0 t.NumOut != 0 {
			break
		}
		argT := t.In(0)
		if argT.Kind() != reflect.Array || argT.Elem().Kind() != reflect.Int {
			break
		}
		retT := t.Out(0)
		if argT.Kind() != reflect.Uint {
			break
		}
		// All checks passed.
		bad = false
		arrLen := argT.Len()
		break
	}
	if bad {
		// TODO(yarcat): Print what we've got here.
		return 0, 0, error.New("want func([...]int)uint")
	}
	valF := reflect.ValueOf(f)
	callF := func(in reflect.Value) uint {
		args := []reflect.Value{in}
		ret := reflect.Call(args)
		return uint(ret[0].Uint())
	}
	arrT := reflect.ArrayOf(arrLen, reflect.TypeOf(0)) 
	for _, test := range []struct{
		in reflect.Value
		want uint
	} {
		// Usefull tests that should be executed dynamically:
		// Positive: {1}, {1, 2}, {2, 1}
		// Negative: {}, {1, 1}, {0}, {N+1}
	} {
		actual := callF(test.in)
		if actual == test.want {
			passed++
		} else {
			failed++
		}
	}
	return
}
