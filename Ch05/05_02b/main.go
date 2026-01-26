package main

import (
	"fmt"
	"strconv"
)

func main() {
	dog := Dog{"Poodle", "Woof"}
	fmt.Printf("The %v says %v!\n", dog.Breed, dog.Sound)

	var n number = 8
	fmt.Println(n.ToString())
}

type Dog struct {
	Breed string
	Sound string
}

type number int64

func (n number) ToString() string {
	return strconv.Itoa(int(n))
}
