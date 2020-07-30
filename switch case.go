package main
import (
    "fmt"
    "time"
)
func main() {
    i := 3
    fmt.Print("Write ", i, " as ")
    switch i {
    case 1:
        fmt.Println("I")
    case 2:
        fmt.Println("II")
    case 3:
        fmt.Println("III")
    }
    switch time.Now().Weekday() {
    case time.Saturday, time.Sunday:
        fmt.Println("freeday")
    default:
        fmt.Println("weekday")
    }
    t := time.Now()
    switch {
    case t.Hour() < 12:
        fmt.Println("before noon")
    default:
        fmt.Println("after noon")
    }
    whatAmI := func(i interface{}) {
        switch t := i.(type) {
        case bool:
            fmt.Println("bool")
        case int:
            fmt.Println("int")
        default:
            fmt.Printf("unknown %T\n", t)
        }
    }
    whatAmI(true)
    whatAmI(1)
    whatAmI("hey")
}