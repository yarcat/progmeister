package reflect

import (
	"log"
	"reflect"
)

// TypeMatcher represents a type that describes itself as a matcher of reflect.Type.
type TypeMatcher interface {
	// MatchType returns true if matching was successful.
	MatchType(reflect.Type) bool
}

// FuncMatcher matches in/out sections of the functions.
type FuncMatcher struct {
	// In is a list of argument matchers.
	In []TypeMatcher
	// Out is a list of result matchers.
	Out []TypeMatcher
}

// MatchType matches functions, arguments and return values.
func (m FuncMatcher) MatchType(t reflect.Type) bool {
	if t.Kind() != reflect.Func {
		log.Printf("FuncMatcher: Kind=%v, want Func", t.Kind())
		return false
	}
	if t.NumIn() != len(m.In) {
		log.Printf("FuncMatcher: NumIn=%v, want %v", t.NumIn(), len(m.In))
		return false
	}
	if t.NumOut() != len(m.Out) {
		log.Printf("FuncMatcher: NumOut=%v, want %v", t.NumOut(), len(m.Out))
		return false
	}
	for i, in := range m.In {
		if !in.MatchType(t.In(i)) {
			log.Printf("FuncMatcher: In[%d].Match=false, want true", i)
			return false
		}
	}
	for i, out := range m.Out {
		if !out.MatchType(t.Out(i)) {
			log.Printf("FuncMatcher: Out[%d].Match=false, want true", i)
			return false
		}
	}
	return true
}

// KindMatcher matches Kind. It is expected to be used with basic types e.g. reflect.Int.
type KindMatcher struct {
	// Kind is an expected type.
	Kind reflect.Kind
}

// MatchType matches kind of the object.
func (m KindMatcher) MatchType(o reflect.Type) bool {
	if o.Kind() == m.Kind {
		return true
	}
	log.Printf("KindMatcher: Kind=%v, want %v", o.Kind(), m.Kind)
	return false
}

// ArrayMatcher matches arrays and applies provided matcher to the element type.
// If the matching is successful, the array's length is stored in Len.
type ArrayMatcher struct {
	// M is a matcher that is applied to the element type.
	M TypeMatcher
	// Len is set with the array size.
	Len int
}

// MatchType matches arrays and its type.
func (m *ArrayMatcher) MatchType(t reflect.Type) bool {
	if t.Kind() != reflect.Array {
		log.Printf("ArrayMatcher: Kind=%v, want Array", t.Kind())
		return false
	}
	if !m.M.MatchType(t.Elem()) {
		log.Printf("ArrayMatcher: M.Match=false, want true")
		return false
	}
	m.Len = t.Len()
	return true
}
