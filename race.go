package main
import (
	"fmt"
	"time"
)
var a int
func f() {
	fmt.Println(a)
	a = 1
}
func main() {
	go f()
	fmt.Println(a)
	time.Sleep(1 * time.Second)
	fmt.Println(a)
}
// what is a race condition ?
// A race condition is an undesirable situation that occurs when a device 
// or system attempts to perform two or more operations at the same time, but
// because of the nature of the device or system, the operations must be done
// in the proper sequence to be done correctly.

// how does it occur ?
// A race condition occurs when two or more threads can access shared data and
// they try to change it at the same time. Because the thread scheduling
// algorithm can swap between threads at any time, you don't know the order
// in which the threads will attempt to access the shared data.