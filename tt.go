package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Enter a string")
	var s string
	fmt.Scan(&s)
	length := len(s) - 1
	if (strings.Index(s, "i") == 0 || strings.Index(s, "I") == 0) && (strings.LastIndex(s, "i") == length || strings.LastIndex(s, "N") == length) && (strings.Contains(s, "a") || strings.Contains(s, "A")) {
		fmt.Print("Found!")
	} else {
		fmt.Print("NotFound!")
	}
}