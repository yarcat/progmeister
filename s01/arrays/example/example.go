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

func maxMin(arr [4]int) (max, min int) {
	min, max = arr[0], arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] < min {
			min = arr[i]
		}
		if arr[i] > max {
			max = arr[i]
		}
	}
	return
}

func main() {
	passed, failed, err := arrays.RunTests(arrays.MaxMin, maxMin)
	if err != nil {
		log.Fatalf("RunTests(MaxMin) failed: %v", err)
	}
	fmt.Printf("minMax: passed=%v, failed=%v\n", passed, failed)

	passed, failed, err = arrays.RunTests(arrays.IsPermutation, isPermutation)
	if err != nil {
		log.Fatalf("RunTests(IsPermutation) failed: %v", err)
	}
	fmt.Printf("isPermutation: passed=%v, failed=%v\n", passed, failed)
}
