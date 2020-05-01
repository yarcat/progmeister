// This application is a template for the exercises at https://progmeister.ch/blog/problems/slices/.
// Also, it can be found at https://play.golang.org/p/ttE9fzLzhpU.
package main

import (
	"github.com/yarcat/progmeister/s01/slices"
)

func removeEven(s []int) []int {
	return nil
}

func generateFibonacci(n int) (s []int) {
	return nil
}

func removeAdjacent(s []int) []int {
	return nil
}

func mergeIncreasingSlices(a, b []int) []int {
	return nil
}

func josephus(soldiers, skip int) int {
	return 0
}

func main() {
	slices.TestRemoveEven(removeEven)
	slices.TestGenerateFibonacci(generateFibonacci)
	slices.TestRemoveAdjacent(removeAdjacent)
	slices.TestMergeIncreasingSlices(mergeIncreasingSlices)
	slices.TestJosephus(josephus)
}
