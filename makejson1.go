package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func main() {
	/*
		Write a program which prompts the user to first enter a name, and then enter an address. Your program should create a map and add the name and address to the map using the keys â€œnameâ€ and â€œaddressâ€, respectively. Your program should use Marshal() to create a JSON object from the map, and then your program should print the JSON object.
	*/
	nm := make(map[string]string)
	fmt.Printf("Please, type the data:\n")
	var name string
	var address string
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Name: ")
	name, _ = reader.ReadString('\n')
	fmt.Printf("Address: ")
	address, _ = reader.ReadString('\n')

	name = strings.Trim(name, "\n")
	address = strings.Trim(address, "\n")

	// Adding everything to the map created
	nm[name] = address
	marshalled, _ := json.Marshal(nm)
	fmt.Printf("Your JSON Object is the following: \n")
	fmt.Println(string(marshalled))
}