// Write your answer here, and then test your code.
// Your job is to implement the getCartFromJson() method.

package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Change these boolean values to control whether you see
// the expected answer and/or hints.
const (
	showExpectedResult = false
	showHints          = false
)

type CartItem struct {
	Name     string
	Price    float64
	Quantity int
}

func main() {
	// This is how your code will be called.
	// Your answer should be a slice containing CartItem objects.
	// You can edit this code to try different testing cases.
	jsonString := `[{"name":"apple","price":2.99,"quantity":3},
  {"name":"orange","price":1.50,"quantity":8},
  {"name":"banana","price":0.49,"quantity":12}]`

	result := getCartFromJson(jsonString)

	fmt.Println(result)
}

// getCartFromJson() returns a slice containing cartItem objects.
func getCartFromJson(jsonString string) []CartItem {
	var cart []CartItem
	// Your code goes here.
	decoder := json.NewDecoder(strings.NewReader(jsonString))
	_, err := decoder.Token()
	checkError(err)

	var item CartItem
	for decoder.More() {
		err := decoder.Decode(&item)
		checkError(err)

		cart = append(cart, item)
	}

	return cart
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
