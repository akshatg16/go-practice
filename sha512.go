package main
import (
    "crypto/sha512"
    "fmt"
)
func main() {
    s := "sha512 this string"
    h := sha512.New()
    h.Write([]byte(s))
    bs := h.Sum(nil)
    fmt.Println(s)
    fmt.Printf("%x\n", bs)
}