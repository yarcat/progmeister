package arrays

import (
	"errors"
	"log"
	"reflect"

	pr "github.com/yarcat/progmeister/reflect"
)

// f should be func([N]int)(max, min int)
func testMaxMin(f interface{}) (passed, failed int, err error) {
	intArr := &pr.ArrayMatcher{M: &pr.KindMatcher{Kind: reflect.Int}}
	wantFn := &pr.FuncMatcher{
		In: []pr.TypeMatcher{intArr},
		Out: []pr.TypeMatcher{
			&pr.KindMatcher{Kind: reflect.Int},
			&pr.KindMatcher{Kind: reflect.Int},
		},
	}
	if !wantFn.MatchType(reflect.TypeOf(f)) {
		return 0, 0, errors.New("want func([N]int)(min, max int)")
	}

	valF := reflect.ValueOf(f)
	callF := func(in reflect.Value) (max, min int) {
		args := []reflect.Value{in.Elem()}
		ret := valF.Call(args)
		return int(ret[0].Int()), int(ret[1].Int())
	}

	al := intArr.Len
	builder := func() *pr.IntArrayBuilder {
		return &pr.IntArrayBuilder{Count: al}
	}

	for _, test := range []struct {
		name             string
		in               *pr.IntArrayBuilder
		wantMin, wantMax int
	}{
		{name: "default", wantMin: 0, wantMax: 0, in: builder()},
		{name: "straight up", wantMin: 1, wantMax: al, in: builder().FillRange(1, 1)},
		{name: "straight down", wantMin: -al, wantMax: -1, in: builder().FillRange(-1, -1)},
		{name: "shuffle", wantMin: -al / 2, wantMax: -al/2 + al - 1, in: builder().FillRange(-al/2, 1).Shuffle()},
	} {
		actualMax, actualMin := callF(test.in.Build())
		if actualMin == test.wantMin && actualMax == test.wantMax {
			log.Println("PASS:", test.name)
			passed++
		} else {
			log.Println("FAIL:", test.name)
			failed++
		}
	}
	return
}
