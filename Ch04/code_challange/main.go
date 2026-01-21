// Write your answer here, and then test your code.

package main

import "fmt"

// Change these boolean values to control whether you see
// the expected answer and/or hints.
const (
	showExpectedResult = true
	showHints          = false
)

func main() {
	// This is how your code will be called.
	// Your answer should be the largest value in the numbers array.
	// You can edit this code to try different testing cases.
	var cart []CartItem
	apples := CartItem{"apple", 2.99, 3}
	oranges := CartItem{"orange", 1.50, 8}
	bananas := CartItem{"banana", .49, 12}
	cart = append(cart, apples, oranges, bananas)
	result := calculateTotal(cart)

	fmt.Println(result)
}

type CartItem struct {
	Name     string
	Price    float64
	Quantity int
}

// calculateTotal() returns the total value of the shopping cart.
func calculateTotal(cart []CartItem) float64 {
	var sum float64 = 0

	for _, cartItem := range cart {
		sum += cartItem.Price * float64(cartItem.Quantity)
	}

	return sum
}
