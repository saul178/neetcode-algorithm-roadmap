package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "hello world"
	strToChar := strings.Split(str, "")

	fmt.Println(strToChar)
	fmt.Println(strToChar[0])
}
