package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)

	fmt.Printf("Go launched %v\n", t)

	n := time.Now()
	fmt.Println("Now", n)

	fmt.Println(n.Format(time.ANSIC))
}
