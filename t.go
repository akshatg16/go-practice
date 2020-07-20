package main

import "fmt"

func main() {

	var number float64
	fmt.Printf("Enter Floating Number:")
	fmt.Scan(&number)
	var convert int = int(number)
	fmt.Println("The Integral Value of input is", convert)
}
