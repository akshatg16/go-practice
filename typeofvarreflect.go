package main 
import ( 
	"fmt"
	"reflect"
) 
func main() { 
var1 := "hello world"
var2 := 10 
var3 := 1.55 
var4 := true
var5 := []string{"foo", "bar", "baz"} 
var6 := map[int]string{100: "Ana", 101: "Lisa", 102: "Rob"} 
var7 := complex(9, 15) 
        fmt.Println("Using Percent T with Printf") 
	fmt.Println() 
        fmt.Printf("var1 = %T\n", var1) 
	fmt.Printf("var2 = %T\n", var2) 
	fmt.Printf("var3 = %T\n", var3) 
	fmt.Printf("var4 = %T\n", var4) 
	fmt.Printf("var5 = %T\n", var5) 
	fmt.Printf("var6 = %T\n", var6) 
	fmt.Printf("var7 = %T\n", var7) 
        fmt.Println() 
	fmt.Println() 
	fmt.Println("Using reflect.ValueOf.Kind() Function") 
	fmt.Println() 
        fmt.Println("var1 = ", reflect.ValueOf(var1).Kind()) 
	fmt.Println("var2 = ", reflect.ValueOf(var2).Kind()) 
	fmt.Println("var3 = ", reflect.ValueOf(var3).Kind()) 
	fmt.Println("var4 = ", reflect.ValueOf(var4).Kind()) 
	fmt.Println("var5 = ", reflect.ValueOf(var5).Kind()) 
	fmt.Println("var6 = ", reflect.ValueOf(var6).Kind()) 
	fmt.Println("var7 = ", reflect.ValueOf(var7).Kind()) 

} 
