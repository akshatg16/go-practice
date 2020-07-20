package main

import (
	"fmt"
  "io/ioutil"
  "bufio"
  "strings"
  "os"
	)

  type namestruct struct {
      fname string
      lname  string
  }

func main() {

    fmt.Print("Enter file name: ")

    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan() // use `for scanner.Scan()` to keep reading
    filename := scanner.Text()

    nameslice := make([]namestruct, 2)

    dataread, _ := ioutil.ReadFile(filename)
    s := string(dataread)
  	lines := strings.Split(s, "\n")

    var entrylist namestruct

    for _, line := range lines {
        names := strings.Split(line, " ")
        if names[0]!="" {
        entrylist.fname=names[0]
        entrylist.lname=names[1]
        nameslice = append(nameslice, entrylist)
      }
    }


    fmt.Println (nameslice)

}
