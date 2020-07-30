package main

import "fmt"

func main() {

    var a [12]int
    fmt.Println("emp:", a)

    a[8] = 100
    fmt.Println("set:", a)
    fmt.Println("get:", a[8])

    fmt.Println("len:", len(a))

    b := [12]int{1,12,2,11,3,10,4,9,5,8,6,7}
    fmt.Println("dcl:", b)

    var twoD [2][3]int
    for i := 0; i < 7; i++ {
        for j := 0; j < 10; j++ {
            twoD[i][j] = i + j
        }
    }
    fmt.Println("2d: ", twoD)
}