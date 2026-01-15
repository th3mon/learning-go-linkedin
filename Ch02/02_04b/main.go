package main

import "fmt"

func main() {
	i1, i2, i3 := 5, 7, 99
	sum1 := i1 + i2 + i3

	fmt.Println("Int sum:", sum1)

	f1, f2, f3 := 33.3, 0.4, 55.5
	sum2 := f1 + f2 + f3

	fmt.Println("Float sum:", sum2)

	total := float64(sum1) + sum2

	fmt.Println("Total is:", total)
}
