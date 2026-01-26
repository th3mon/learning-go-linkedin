// Write your answer here, and then test your code.

package main

import (
	"fmt"
	"strconv"
)

// Change these boolean values to control whether you see
// the expected answer and/or hints.
const (
	showExpectedResult = false
	showHints          = false
)

func main() {
	// This is how your code will be called.
	// You can edit this code to try different testing cases.
	value1 := "10"
	value2 := "5.5"
	operation := "+"
	result := calculate(value1, value2, operation)

	fmt.Println(result)
}

// calculate() returns the result of the requested operation.
func calculate(input1 string, input2 string, operation string) float64 {
	// Your code goes here.
	input1AsFloat := convertInputToValue(input1)
	input2AsFloat := convertInputToValue(input2)

	switch operation {

	case "+":
		return addValues(input1AsFloat, input2AsFloat)

	case "-":
		return subtractValues(input1AsFloat, input2AsFloat)

	case "*":
		return multiplyValues(input1AsFloat, input2AsFloat)

	case "/":
		return divideValues(input1AsFloat, input2AsFloat)

	default:
		fmt.Printf("Operation %v is not avaible.\n", operation)
		return 0
	}
}

func convertInputToValue(input string) float64 {
	inputAsFloat, err := strconv.ParseFloat(input, 64)
	if err != nil {
		fmt.Println(err)
		panic("Value should be a number")
	}

	return inputAsFloat
}

func addValues(value1, value2 float64) float64 {
	return value1 + value2
}

func subtractValues(value1, value2 float64) float64 {
	return value1 - value2
}

func multiplyValues(value1, value2 float64) float64 {
	return value1 * value2
}

func divideValues(value1, value2 float64) float64 {
	return value1 / value2
}
