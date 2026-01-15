package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")

	messge, _ := reader.ReadString('\n')
	fmt.Println(messge)

	messge, _ = reader.ReadString('\n')
	number, err := strconv.ParseFloat(strings.TrimSpace(messge), 64)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("The number is:", number)
	}
}
