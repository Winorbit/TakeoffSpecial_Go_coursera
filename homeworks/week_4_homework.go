package main

import "fmt"
import "io"

type Animal interface { 
	Eat()
	Move()
	Speak() 
}

type Cow struct {}

func (c Cow) Eat() {
  fmt.Println("grass") 
}

func (c Cow) Move() {
  fmt.Println("walk") 
}

func (c Cow) Speak() {
  fmt.Println("moo") 
}


type Bird struct {}

func (b Bird) Eat() {
  fmt.Println("worms") 
}
func (b Bird) Move() {
  fmt.Println("fly") 
}
func (b Bird) Speak() {
  fmt.Println("peep") 
}


type Snake struct {}

func (s Snake) Eat() { 
  fmt.Println("mice") 
}
func (s Snake) Move() { 
  fmt.Println("slither") 
}
func (s Snake) Speak() { 
  fmt.Println("hsss") 
}


func main() {
	
	var animal_map = make(map[string]Animal)

	for {
		fmt.Print("> ")
	    fmt.Println("Please enter message to create new animal. \n For this enter words like:\n newanimal cow Korovka")

		var cmdType string 
		var animal string
		var action string
		var name string
		
		var err error	
		_, err = fmt.Scan(&cmdType)
		if err == io.EOF { return }
		

		if cmdType == "newanimal" {
			fmt.Scan(&name)
			fmt.Scan(&animal)

			fmt.Println(cmdType, name, animal, action)

			if animal == "cow" {
				animal_map[name] = Cow{}
			} else if animal == "bird" {
				animal_map[name] = Bird{}
			} else {
				animal_map[name] = Snake{}
			}
			fmt.Println("Created it!")

		} else {
			fmt.Scan(&name)
			fmt.Scan(&action)
			an := animal_map[name]

			//Handle action type.
			if action == "move" {
				an.Move()
			} else if action == "speak" {
				an.Speak()
			} else {
				an.Eat()
			}
		}
	}
}
