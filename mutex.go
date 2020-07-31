package main 
import ( 
	"fmt"
	"sync"  
) 
var GFG = 0 
func worker(wg *sync.WaitGroup) { 
	GFG = GFG + 1 
wg.Done() 
} 
func main() { 
var w sync.WaitGroup 
for i := 100; i < 1000; i++ { 
		w.Add(1)		 
		go worker(&w) 
	} 
w.Wait() 
	fmt.Println("Value of x", GFG) 
} 
