package main

import (
	"fmt"
)

func main() {
	arr := []string{"jeden", "dwa"}
	arr2 := make([]string, 0, 1)
	arr2 = append(arr2, arr[1:]...)
	fmt.Println(arr[0:])
	fmt.Println(arr2)
}
