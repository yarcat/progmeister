package arrays

// RunTests executes tests for the given excercise.
//
// Usage example:
//
//  package main
//
//  import (
//    "fmt"
//    "log"
//
//    "github.com/yarcat/progmeister/s01/arrays"
//  )
//
//  func isPermutation(arr [10]int) uint {
//    // Implementation goes here.
//    return 0
//  }
//
//  func main() {
//    passed, failed, err := arrays.RunTests(arrays.IsPermutation, isPermutation)
//    if err != nil {
//      log.Fatalf("RunTests failed: %v", err
//    }
//    fmt.Printf("Passed=%v, failed=%v\n", passed, failed")
//  }
//
func RunTests(test Test, f interface{}) (passed, failed int, err error) {
	return test.run(f)
}
