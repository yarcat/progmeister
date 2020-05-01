# slices

This package contains test functions for slices at https://progmeister.ch/blog/problems/slices/

# installation

`go get github.com/yarcat/progmeister/s01/slices`

# example

```golang
package main

import "github.com/yarcat/progmeister/s01/slices"

func removeEven(s [int]) []int {
    // Implementation goes here.
    return nil
}

func generateFibonacci(n int) []int {
    // Implementation goes here.
}

func main() {
    slices.TestRemoveEven(removeEven)
    slices.TestGenerateFibonacci(generateFibonacci)
}
```

# useful links

* https://github.com/golang/go/wiki/SliceTricks
* Function templates https://play.golang.org/p/ttE9fzLzhpU
* Some solutions "eJzLKCkpKLbS1y/ISazUS8/PScxL18svStcv0Dd1NvQMycooSwrKBAD+GA0X". But
  you should know what to do with it :wink:
