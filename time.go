package main 
import ("fmt"
       "time"
)
func main() { 
	today := time.Now() 
	tomorrow := today.Add(24 * time.Hour) 
	g1 := today.Before(tomorrow) 
	fmt.Println("today before tomorrow:", g1) 
	g2 := tomorrow.After(today) 
	fmt.Println("tomorrow after today:", g2) 
} 
