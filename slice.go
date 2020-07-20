package main
import "fmt" 
for input.Scan() 
    if input.text() == "X" {break}
    fmt.Println(input.text())
	
	fmt.Println("Val i %d\n", i)
	if i > 3 {
	   fmt.Printf("Extend slice")
	   sli = sli[:1]
	}
	j := i
	intVal, err := strconv.Atoi(input.text())
	if err != nil {
       fmt.Println(err)
   }
  
   sli[j] = intVal
	
  i += 1
}