package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "hello world"
	strToChar := strings.Split(str, "")
	seen := make(map[string]bool)
	fmt.Println("seen is initialized: ", seen)

	for _, value := range strToChar {
		seen[value] = true
	}
	fmt.Println("seen has been filled with values: ", seen)
	fmt.Println(seen["s"])
}
