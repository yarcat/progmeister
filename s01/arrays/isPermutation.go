// Package arrays contains tests for https://progmeister.ch/blog/problems/arrays/.
package arrays

import (
	"errors"
	"log"
	"math/rand"
	"reflect"
)

func testIsPermutation(f interface{}) (passed, failed int, err error) {
	arrLen := 0
	t := reflect.TypeOf(f)
	bad := true
	for {
		if t.Kind() != reflect.Func || t.NumIn() != 1 || t.NumOut() != 1 {
			break
		}
		argT := t.In(0)
		if argT.Kind() != reflect.Array || argT.Elem().Kind() != reflect.Int {
			break
		}
		if t.Out(0).Kind() != reflect.Uint {
			break
		}
		// All checks passed.
		bad = false
		arrLen = argT.Len()
		break
	}
	if bad {
		// TODO(yarcat): Print what we've got here.
		return 0, 0, errors.New("want func([...]int)uint")
	}
	valF := reflect.ValueOf(f)
	callF := func(in reflect.Value) uint {
		args := []reflect.Value{in.Elem()}
		ret := valF.Call(args)
		return uint(ret[0].Uint())
	}
	for _, test := range []struct {
		name string
		in   reflect.Value
		want uint
	}{
		// Usefull tests that should be executed dynamically:
		{name: "straight up", want: 1, in: makeIntArray(arrLen, fillRange(1, 1))},
		{name: "straight down", want: 1, in: makeIntArray(arrLen, fillRange(arrLen, -1))},
		{name: "random permutation", want: 1, in: makeIntArray(arrLen, fillRange(1, 1), randomize())},
		{name: "zero value", want: 0, in: makeIntArray(arrLen)},
		{name: "same value", want: 0, in: makeIntArray(arrLen, fillWith(1))},
		{name: "contains 0", want: 0, in: makeIntArray(arrLen, fillRange(0, 1))},
		{name: "contains N+1", want: 0, in: makeIntArray(arrLen, fillRange(2, 1))},
	} {
		actual := callF(test.in)
		if actual == test.want {
			log.Println("PASS:", test.name)
			passed++
		} else {
			log.Println("FAIL:", test.name)
			failed++
		}
	}
	return
}

type makeIntArrayOption func(v reflect.Value)

func fillWith(n int) func(reflect.Value) {
	return func(v reflect.Value) {
		for i := 0; i < v.Elem().Len(); i++ {
			v.Elem().Index(i).SetInt(int64(n))
		}
	}
}

func fillRange(start, step int) func(reflect.Value) {
	return func(v reflect.Value) {
		for i, x := 0, int64(start); i < v.Elem().Len(); i++ {
			v.Elem().Index(i).SetInt(x)
			x += int64(step)
		}
	}
}

func randomize() func(reflect.Value) {
	return func(v reflect.Value) {
		for i := 0; i < v.Elem().Len()-1; i++ {
			j := rand.Int()%(v.Elem().Len()-i) + i
			if j == i {
				continue
			}
			// Swap i-th and j-th.
			ival := v.Elem().Index(i).Int()
			jval := v.Elem().Index(j).Int()
			v.Elem().Index(i).SetInt(jval)
			v.Elem().Index(j).SetInt(ival)
		}
	}
}

func makeIntArray(count int, opts ...makeIntArrayOption) reflect.Value {
	arr := reflect.New(reflect.ArrayOf(count, reflect.TypeOf(int(0))))
	for _, f := range opts {
		f(arr)
	}
	return arr
}
