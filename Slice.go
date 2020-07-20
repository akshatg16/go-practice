package main

import (
	"fmt"
	"strconv"
	"sort"
)

func main() {
	// It was hard to understand the task
	// Let the initial slice be like this
	slice := make([]int, 0, 3)
	var current string
	fmt.Println("Enter a number to add it to the slice or X to stop:")
	fmt.Scan(&current)
	for current != "X" {
		slice = append(slice, 1)
		slice[len(slice) - 1], _ = strconv.Atoi(current)
		sort.Ints(slice)
		fmt.Println(slice)
		fmt.Scan(&current)
	}
}