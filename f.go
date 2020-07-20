package main

import (
	"fmt"
	"log"
	"strings"
)

func main() {
	var st string

	for i := 0; i < 2; i++ {
		_, err := fmt.Scanln(&st)

		if err != nil {
			log.Fatal(err)
		}

		st = strings.ToLower(st)

		if strings.HasPrefix(st, "i") && strings.HasSuffix(st, "n") && strings.Contains(st, "a") {
			fmt.Println("Found")
		} else {
			fmt.Println("Not Found")
		}
	}
}
