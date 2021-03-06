// Program hello demonstrates simple gookit/color usage.
//
// More examples at https://github.com/gookit/color.
package main

import (
	"math/rand"
	"time"

	"github.com/gookit/color"
)

func main() {
	color.Cyan.Println("Hello, world!")

	var colors []color.Color
	for _, c := range color.FgColors {
		colors = append(colors, c)
	}

	rand.Seed(time.Now().UnixNano())
	for _, r := range "Hello, world!" {
		i := rand.Int() % len(colors)
		colors[i].Print(string(r))
	}
	color.Println()

	color.Println("Hello, <bold>Progmeister</>!")
}
