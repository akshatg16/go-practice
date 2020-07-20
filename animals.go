package main
import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)
func main() {
	inputReader := bufio.NewReader(os.Stdin)
	sliceOfAnimals := []animal{} 
	for {
		fmt.Println("Welcome to Animal Park where you can create and query animals")
		fmt.Println("To begin, you can do either of the following types of queries:")
		fmt.Println("To create an animal for the database, you will need to use the following Syntax:")
		fmt.Println("newanimal (name of animal) (type of animal)")
		fmt.Println("Upon creation of an animal, you can then query the animal(s) by doing the following:")
		fmt.Println("query (name of animal) (action of animal)")
		fmt.Print(">")
		userQuery, err := inputReader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		userQuery = userQuery[:len(userQuery)-2] 
		sliceOfQuery := strings.Split(userQuery, " ")
		if len(sliceOfQuery) > 3 || len(sliceOfQuery) < 3 {
			fmt.Println("Invalid query 1!")
		}
		determinant := sliceOfQuery[0]
switch determinant {
		case "newanimal":
			fmt.Printf("5. '%v'\n\n", sliceOfQuery[2])
			if sliceOfQuery[2] == "cow" {
				sliceOfAnimals = append(sliceOfAnimals, cow{name: sliceOfQuery[1]})
				fmt.Println("Created it!")
			} else if sliceOfQuery[2] == "snake" {
				sliceOfAnimals = append(sliceOfAnimals, snake{name: sliceOfQuery[1]})
				fmt.Println("Created it!")
			} else if sliceOfQuery[2] == "bird" {
				sliceOfAnimals = append(sliceOfAnimals, bird{name: sliceOfQuery[1]})
				fmt.Println("Created it!")
			} else {
				fmt.Println("Invalid Query 2!")
			}
		case "query":
			for _, animal := range sliceOfAnimals {
				if animal.getName() == sliceOfQuery[1] { 
					if sliceOfQuery[2] == "move" {
						animal.move()
					} else if sliceOfQuery[2] == "eat" {
						animal.eat()
					} else if sliceOfQuery[2] == "speak" {
						animal.speak()
					}
				} else {
					fmt.Println("Animal name not found!")
				}
			}
default:
			fmt.Println("Invalid user query 3!")
            }
	}
}
type animal interface {
	eat()
	move()
	speak()
	getName() string 
}
type cow struct{ name string }
type snake struct{ name string }
type bird struct{ name string }
func (c cow) getName() string {
	return c.name
}
func (s snake) getName() string {
	return s.name
}
func (b bird) getName() string {
	return b.name
}
func (c cow) eat() {
	fmt.Printf("%s eats grass\n", c.name)
}
func (c cow) move() {
	fmt.Printf("%s walks\n", c.name)
}
func (c cow) speak() {
	fmt.Printf("%s moos\n", c.name)
}
func (s snake) eat() {
	fmt.Printf("%s eats mice\n", s.name)
}
func (s snake) move() {
	fmt.Printf("%s slithers\n", s.name)
}
func (s snake) speak() {
	fmt.Printf("%s hisses\n", s.name)
}
func (b bird) eat() {
	fmt.Printf("%s eats worms\n", b.name)
}
func (b bird) move() {
	fmt.Printf("%s flys\n", b.name)
}
func (b bird) speak() {
	fmt.Printf("%s peeps and chirps\n", b.name)
}
