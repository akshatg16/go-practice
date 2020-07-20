package main

import (
  "bufio"
  "fmt"
  "os"
  "sort"
  "strconv"
  "strings"
)

func main() {
  fmt.Println("Enter distinct integers:")
  scanner := bufio.NewScanner(os.Stdin)
  for scanner.Scan() {
    var ls []int
    t := scanner.Text()
    for _, f := range strings.Fields(t) {
      n, err := strconv.Atoi(f)
      if err == nil {
        ls = append(ls, n)
      }
    }
    sort.Ints(ls)
    fmt.Println(ls)
  }
}