package reflect

import (
	"math/rand"
	"reflect"
)

// IntArrayBuilder helps to create int arrays dynamically.
//
// Example:
//
//	Create an array with 10 elements filled with values [1, 3, 5, ...] in
//	random order:
//
//		v := (&IntArrayBuilder{10}).
//			FillRange(1, 2).
//			Shuffle().
//			Build()
type IntArrayBuilder struct {
	// Count is amount of elements that the array should have.
	Count int
	// fns is a list of functions that will be applied on the array
	// after its creation.
	fns []func(reflect.Value)
}

// Fill fills the entire array with x.
func (b *IntArrayBuilder) Fill(x int) *IntArrayBuilder {
	b.fns = append(b.fns, func(v reflect.Value) {
		for i := 0; i < v.Elem().Len(); i++ {
			v.Elem().Index(i).SetInt(int64(x))
		}
	})
	return b
}

// FillRange fills the entire array with values starting from start with
// the step provided. Negative steps are valid.
func (b *IntArrayBuilder) FillRange(start, step int) *IntArrayBuilder {
	b.fns = append(b.fns, func(v reflect.Value) {
		for i, x := 0, int64(start); i < v.Elem().Len(); i++ {
			v.Elem().Index(i).SetInt(x)
			x += int64(step)
		}
	})
	return b
}

// Shuffle randomly shuffles elements within array. This option should be
// used after another (e.g. FillRange) to ensure elements aren't sorted.
func (b *IntArrayBuilder) Shuffle() *IntArrayBuilder {
	b.fns = append(b.fns, func(v reflect.Value) {
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
	})
	return b
}

// Build returns dynamically created array. Please note that the value
// returned represents a pointer to the created array. You will have to use
// reflect.Value.Elem function if you need to dereference it.
func (b IntArrayBuilder) Build() reflect.Value {
	arr := reflect.New(reflect.ArrayOf(b.Count, reflect.TypeOf(int(0))))
	for _, f := range b.fns {
		f(arr)
	}
	return arr
}
