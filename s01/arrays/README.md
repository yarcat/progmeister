## Idea & Example

The framework should help executing tests created by our students. Ideal invocation example:

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
