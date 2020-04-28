// Package arrays contains tests for https://progmeister.ch/blog/problems/arrays/.
package arrays

import (
	"errors"
	"log"
	"reflect"

	pr "github.com/yarcat/progmeister/reflect"
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
		in   *pr.IntArrayBuilder
		want uint
	}{
		// Usefull tests that should be executed dynamically:
		{name: "straight up", want: 1, in: (&pr.IntArrayBuilder{Count: arrLen}).FillRange(1, 1)},
		{name: "straight down", want: 1, in: (&pr.IntArrayBuilder{Count: arrLen}).FillRange(arrLen, -1)},
		{name: "random permutation", want: 1, in: (&pr.IntArrayBuilder{Count: arrLen}).FillRange(1, 1).Shuffle()},
		{name: "zero value", want: 0, in: &pr.IntArrayBuilder{Count: arrLen}},
		{name: "same value", want: 0, in: (&pr.IntArrayBuilder{Count: arrLen}).Fill(1)},
		{name: "contains 0", want: 0, in: (&pr.IntArrayBuilder{Count: arrLen}).FillRange(0, 1)},
		{name: "contains N+1", want: 0, in: (&pr.IntArrayBuilder{Count: arrLen}).FillRange(arrLen+1, -1)},
	} {
		actual := callF(test.in.Build())
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
