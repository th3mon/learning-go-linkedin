// Write your answer here, and then test your code.
// Your job is to implement the findLargest() method.

package main

import (
	"fmt"
)

// Change these boolean values to control whether you see
// the expected answer and/or hints.
const (
	showExpectedResult = false
	showHints          = false
)

func main() {
	colorNames := []string{"Red", "Green", "Blue"}
	hexValues := []int{0xFF0000, 0x00FF00, 0x0000FF}
	colors := slicesToObjects(colorNames, hexValues)

	fmt.Println(colors)
}

// slicesToObjects() returns a slice of Color objects
func slicesToObjects(colorNames []string, hexValues []int) []Color {
	// Your code goes here.
	colors := make([]Color, 0)

	for i := range colorNames {
		colors = append(colors, Color{colorNames[i], hexValues[i]})
	}
	// Create a slice of Color objects
	return colors
}

type Color struct {
	Name string
	Hex  int
}
