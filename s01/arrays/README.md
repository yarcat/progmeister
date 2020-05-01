## Supported Test Runners

All test runners correspond to the function names proposed on the [arrays excercises page](https://progmeister.ch/blog/problems/arrays/), however to make them public we have to name them using [CamelCase] (aka CapWords) notation.

* `arrays.MaxMin`
* `arrays.IsPermutation`

[CamelCase]: https://simple.wikipedia.org/wiki/CamelCase

## Idea & Example

The framework should help executing tests created by our students. They should work with arrays of any length. Ideal invocation example:

```golang
package main

import (
	"fmt"
	"log"

	"github.com/yarcat/progmeister/s01/arrays"
)

func isPermutation(arr [10]int) uint {
	// Custom implementation goes here.
	return 0
}

func main() {
	passed, failed, err := arrays.RunTests(arrays.IsPermutation, isPermutation)
	if err != nil {
		log.Fatalf("RunTests failed: %v", err)
	}
	fmt.Printf("Passed=%v, failed=%v\n", passed, failed)
}

```

## Some useful reflect experiments

Creating an array, setting its item and passing to a function: https://play.golang.org/p/n3TovYYUllE 
