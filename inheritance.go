 package main 
import ( 
	"fmt"
)  
type Movie struct{ 
	Universe string 
} 
func (movie Movie) MovieUniverse() string { 
	return movie.Universe 
} 
type Bollywood struct{ 
	Movie
} 
type Tollywood struct{ 
	Movie 
} 
func main() { 
	c1 := Bollywood{ 
		Movie{ 
		Universe: "War & Heropanti", 
		}, 
		}  
		fmt.Println("Bollywood:", c1.MovieUniverse())	 
	
	c2 := Tollywood{ 
		Movie{ 
		Universe : "Sarrainodu", 
		}, 
	        }  
	fmt.Println("Tollywood:", c2.MovieUniverse()) 
} 
