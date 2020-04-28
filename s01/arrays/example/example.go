package main

import (
	"fmt"
	"log"

	"github.com/yarcat/progmeister/s01/arrays"
)

func isPermutation(arr [4]int) uint {
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
