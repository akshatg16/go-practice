package main
import(
       "fmt"
       "sort"
      )
func main(){
strs := []string{"q","w","e","r","t","y"}
sort.Strings(strs)
fmt.Println("Strings:" ,strs)

ints:= []int{8,0,3,6,7,4}
sort.Ints(ints)
fmt.Println("Ints:",ints)

s:= sort.IntsAreSorted(ints)
fmt.Println("Sorted:",s)
}