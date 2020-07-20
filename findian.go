package main

import (
        "bufio"
        "fmt"
        "os"
        "regexp"
)

func main() {
             fmt.Println("Hi! Type Something: ")
             reader := bufio.NewReader(os.Stdin)
             text, _ := reader.ReadString('\n')

             re := regexp.MustCompile(`(?i)^i.*a.*n\n$`)

             if re.MatchString(text){
             fmt.Println("Found!")
} else {
        fmt.Println("Not Found!")
}
}