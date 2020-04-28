package arrays

import (
	"errors"
	"log"
	"reflect"

	pr "github.com/yarcat/progmeister/reflect"
)

// f should be func(arra [N]int)uint
func testIsPermutation(f interface{}) (passed, failed int, err error) {
	intArr := &pr.ArrayMatcher{M: &pr.KindMatcher{Kind: reflect.Int}}
	wantFn := &pr.FuncMatcher{
		In:  []pr.TypeMatcher{intArr},
		Out: []pr.TypeMatcher{&pr.KindMatcher{Kind: reflect.Uint}},
	}
	if !wantFn.MatchType(reflect.TypeOf(f)) {
		return 0, 0, errors.New("want func([N]int)uint")
	}

	arrLen := intArr.Len
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
		{name: "shuffle", want: 1, in: (&pr.IntArrayBuilder{Count: arrLen}).FillRange(1, 1).Shuffle()},
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
