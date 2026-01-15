// Write your answer here, and then test your code.
package main

// Uncomment this import section to use required Go packages
import (
	"fmt"
	"strconv"
)

func main() {
	value := calculate("665", "1")

	fmt.Println(value)
}

// Change these boolean values to control whether you see
// the expected answer and/or hints.
const (
	showExpectedResult = false
	showHints          = false
)

// calculate() returns the the result of adding 2 numbers
// after parsing them from strings
func calculate(value1 string, value2 string) float64 {
	// Your code goes here.

	// Convert the first string to a float64
	value1AsFloat, err := strconv.ParseFloat(value1, 64)
	if err != nil {
		fmt.Println(err)
		panic("Value1 must be a number")
	}

	// Convert the second string to a float64
	value2AsFloat, err := strconv.ParseFloat(value2, 64)
	if err != nil {
		fmt.Println(err)
		panic("Value2 must be a number")
	}

	// Calculate and return the result

	return value1AsFloat + value2AsFloat
}
