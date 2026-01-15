package main

import (
	"fmt"
)

func main() {
	m := make(map[string]int)
	m["jeden"] = 1
	fmt.Println(m)

	fmt.Println("Pointers")

	number := 42
	var p *int = &number
	var p2 *int = &number

	*p += *p2

	if p == nil {
		fmt.Println("p is nil")
	} else {
		fmt.Println(*p)
	}
}
